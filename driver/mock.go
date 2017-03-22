/*
 * go-mysqlstack
 * xelabs.org
 *
 * Copyright (c) XeLabs
 * GPL License
 *
 */

package driver

import (
	"errors"
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/XeLabs/go-mysqlstack/sqldb"
	"github.com/XeLabs/go-mysqlstack/sqlparser/depends/sqltypes"
	"github.com/XeLabs/go-mysqlstack/xlog"
)

func randomPort(min int, max int) int {
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	d, delta := min, (max - min)
	if delta > 0 {
		d += rand.Intn(int(delta))
	}
	return d
}

type exprResult struct {
	expr   *regexp.Regexp
	result *sqltypes.Result
	err    error
}

type CondType int

const (
	COND_NORMAL CondType = iota
	COND_DELAY
	COND_ERROR
	COND_PANIC
)

type Cond struct {
	// Cond type.
	Type CondType

	// Query string
	Query string

	// Query results
	Result *sqltypes.Result

	// Panic or Not
	Panic bool

	// Return Error if Error is not nil
	Error error

	// Delay(ms) for results return
	Delay int
}

// Test Handler
type TestHandler struct {
	deny  bool
	log   *xlog.Log
	mu    sync.RWMutex
	ss    map[uint32]*Session
	conds map[string]*Cond

	// patterns is a list of regexp to results.
	patterns      []exprResult
	patternErrors []exprResult

	// How many times a query was called.
	queryCalled map[string]int
}

func NewTestHandler(log *xlog.Log) *TestHandler {
	return &TestHandler{
		log:         log,
		ss:          make(map[uint32]*Session),
		conds:       make(map[string]*Cond),
		queryCalled: make(map[string]int),
	}
}

func (th *TestHandler) setCond(cond *Cond) {
	th.mu.Lock()
	defer th.mu.Unlock()
	th.conds[strings.ToLower(cond.Query)] = cond
	th.queryCalled[strings.ToLower(cond.Query)] = 0
}

func (th *TestHandler) removeCond(query string) {
	th.mu.Lock()
	defer th.mu.Unlock()
	delete(th.conds, strings.ToLower(query))
	delete(th.queryCalled, strings.ToLower(query))
}

// ResetAll resets all querys.
func (th *TestHandler) ResetAll() {
	th.mu.Lock()
	defer th.mu.Unlock()
	for k, _ := range th.conds {
		delete(th.conds, k)
	}
	th.patterns = make([]exprResult, 0, 4)
	th.patternErrors = make([]exprResult, 0, 4)
}

// ConnectionCheck impl.
func (th *TestHandler) SessionCheck(s *Session) error {
	return nil
}

// AuthCheck impl.
func (th *TestHandler) AuthCheck(s *Session) error {
	user := s.Auth.User()
	if user != "mock" {
		return sqldb.NewSQLError(sqldb.ER_ACCESS_DENIED_ERROR, "Access denied for user '%v'", user)
	}
	return nil
}

// Register impl.
func (th *TestHandler) NewSession(s *Session) {
	th.mu.Lock()
	defer th.mu.Unlock()
	th.ss[s.ID] = s
}

// UnRegister impl.
func (th *TestHandler) SessionClosed(s *Session) {
	th.mu.Lock()
	defer th.mu.Unlock()
	delete(th.ss, s.ID)
}

// ComInitDB impl.
func (th *TestHandler) ComInitDB(s *Session, db string) error {
	return nil
}

// ComQuery impl.
func (th *TestHandler) ComQuery(s *Session, query string) (*sqltypes.Result, error) {
	th.log.Debug("test.handler.ComQuery:%v", query)
	query = strings.ToLower(query)

	th.mu.Lock()
	th.queryCalled[query]++
	th.mu.Unlock()

	cond := th.conds[query]
	if cond != nil {
		switch cond.Type {
		case COND_DELAY:
			time.Sleep(time.Millisecond * time.Duration(cond.Delay))
			return cond.Result, nil
		case COND_ERROR:
			return nil, cond.Error
		case COND_PANIC:
			th.log.Panic("test.handler.panic....")
		case COND_NORMAL:
			return cond.Result, nil
		}
	}

	// kill filter.
	if strings.HasPrefix(query, "kill") {
		if id, err := strconv.ParseUint(strings.Split(query, " ")[1], 10, 32); err == nil {
			th.log.Info("mock.to.kill.%v.session", id)
			session := th.ss[uint32(id)]
			if session != nil {
				session.Close()
			}
		}
		return &sqltypes.Result{}, nil
	}

	th.mu.Lock()
	defer th.mu.Unlock()
	// Check query patterns from AddQueryPattern().
	for _, pat := range th.patternErrors {
		if pat.expr.MatchString(query) {
			return nil, pat.err
		}
	}
	for _, pat := range th.patterns {
		if pat.expr.MatchString(query) {
			return pat.result, nil
		}
	}

	return nil, errors.New(fmt.Sprintf("testhandler.query[%v].error[can.not.found.the.cond.please.set.first]", query))
}

// AddQuery used to add a query and its expected result.
func (th *TestHandler) AddQuery(query string, result *sqltypes.Result) {
	th.setCond(&Cond{Type: COND_NORMAL, Query: query, Result: result})
}

// AddQueryDelay used to add a query and returns the expected result after delay_ms.
func (th *TestHandler) AddQueryDelay(query string, result *sqltypes.Result, delay_ms int) {
	th.setCond(&Cond{Type: COND_DELAY, Query: query, Result: result, Delay: delay_ms})
}

// AddQueryError used to add a query which will be rejected by a error.
func (th *TestHandler) AddQueryError(query string, err error) {
	th.setCond(&Cond{Type: COND_ERROR, Query: query, Error: err})
}

// AddQueryPanic used to add query but underflying blackhearted.
func (th *TestHandler) AddQueryPanic(query string) {
	th.setCond(&Cond{Type: COND_PANIC, Query: query})
}

// This code was derived from https://github.com/youtube/vitess.
// AddQueryPattern adds an expected result for a set of queries.
// These patterns are checked if no exact matches from AddQuery() are found.
// This function forces the addition of begin/end anchors (^$) and turns on
// case-insensitive matching mode.
func (th *TestHandler) AddQueryPattern(queryPattern string, expectedResult *sqltypes.Result) {
	if len(expectedResult.Rows) > 0 && len(expectedResult.Fields) == 0 {
		panic(fmt.Errorf("Please add Fields to this Result so it's valid: %v", queryPattern))
	}
	expr := regexp.MustCompile("(?is)^" + queryPattern + "$")
	result := *expectedResult
	th.mu.Lock()
	defer th.mu.Unlock()
	th.patterns = append(th.patterns, exprResult{expr, &result, nil})
}

func (th *TestHandler) AddQueryErrorPattern(queryPattern string, err error) {
	expr := regexp.MustCompile("(?is)^" + queryPattern + "$")
	th.mu.Lock()
	defer th.mu.Unlock()
	th.patternErrors = append(th.patternErrors, exprResult{expr, nil, err})
}

// This code was derived from https://github.com/youtube/vitess.
// GetQueryCalledNum returns how many times db executes a certain query.
func (th *TestHandler) GetQueryCalledNum(query string) int {
	th.mu.Lock()
	defer th.mu.Unlock()
	num, ok := th.queryCalled[strings.ToLower(query)]
	if !ok {
		return 0
	}
	return num
}

func MockMysqlServer(log *xlog.Log, h Handler) (svr *Listener, err error) {
	port := randomPort(7000, 20000)
	addr := fmt.Sprintf(":%d", port)
	if svr, err = NewListener(log, addr, h); err != nil {
		return
	}
	go func() {
		svr.Accept()
	}()
	time.Sleep(100 * time.Millisecond)
	log.Debug("mock.server[%v].start...", addr)
	return
}
