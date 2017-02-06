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
	"strconv"
	"strings"
	"sync"
	"time"

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

type Cond struct {
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
	lock  sync.RWMutex
	ss    map[uint32]*Session
	conds map[string]*Cond
}

func NewTestHandler(log *xlog.Log) *TestHandler {
	return &TestHandler{
		log:   log,
		ss:    make(map[uint32]*Session),
		conds: make(map[string]*Cond),
	}
}

func (th *TestHandler) SetCond(cond *Cond) {
	th.conds[strings.ToUpper(cond.Query)] = cond
}

func (th *TestHandler) RemoveCond(q string) {
	delete(th.conds, strings.ToUpper(q))
}

// IPCheck impl.
func (th *TestHandler) IPCheck(addr string) bool {
	return true
}

// AuthCheck impl.
func (th *TestHandler) AuthCheck(s *Session) bool {
	if s.Auth.User() == "mock" {
		return true
	}

	return false
}

// Register impl.
func (th *TestHandler) Register(s *Session) {
	th.lock.Lock()
	defer th.lock.Unlock()
	th.ss[s.ID] = s
}

// UnRegister impl.
func (th *TestHandler) UnRegister(s *Session) {
	th.lock.Lock()
	defer th.lock.Unlock()
	delete(th.ss, s.ID)
}

// ComQuery impl.
func (th *TestHandler) ComQuery(s *Session, q string) (*sqltypes.Result, error) {
	q = strings.ToUpper(q)
	th.log.Debug("testHandler.ComQuery:%v", q)

	cond := th.conds[q]
	if cond != nil {
		if cond.Delay > 0 {
			time.Sleep(time.Millisecond * time.Duration(cond.Delay))
		}

		switch {
		case cond.Error != nil:
			th.log.Debug("mock.handler.error:%+v....", cond)
			return nil, cond.Error

		case cond.Panic:
			th.log.Panic("test.panic....")

		default:
			return cond.Result, nil
		}
	}

	// kill filter.
	if strings.HasPrefix(q, "KILL") {
		if id, err := strconv.ParseUint(strings.Split(q, " ")[1], 10, 32); err == nil {
			th.log.Info("mock.to.kill.%v.session", id)
			th.lock.RLock()
			session := th.ss[uint32(id)]
			if session != nil {
				session.Close()
			}
			th.lock.RUnlock()
		}
		return &sqltypes.Result{}, nil
	}

	return nil, errors.New("testhandler.comquery.error")
}

func MockMysqlServer(log *xlog.Log, h Handler) (port int, svr *Listener, err error) {
	port = randomPort(4000, 6000)
	addr := fmt.Sprintf(":%d", port)
	if svr, err = NewListener(log, addr, h); err != nil {
		return
	}

	go func() {
		svr.Accept()
	}()
	log.Debug("mock.server[%v].start...", addr)

	return
}
