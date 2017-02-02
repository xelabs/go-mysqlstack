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
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	querypb "github.com/XeLabs/go-mysqlstack/sqlparser/depends/query"
	"github.com/XeLabs/go-mysqlstack/sqlparser/depends/sqltypes"
)

type testHandler struct {
	authDeny bool
	qr       *sqltypes.Result
}

func newTestHandler() *testHandler {
	return &testHandler{}
}

func (th *testHandler) AuthCheck(session *Session) bool {
	if session.Auth.User() == "mock" {
		return true
	}

	return false
}

func (th *testHandler) ComQuery(session *Session, query string) (*sqltypes.Result, error) {
	switch {
	case strings.HasPrefix(query, "insert") == true:
		return &sqltypes.Result{
			RowsAffected: 123,
			InsertID:     123456789,
		}, nil

	case strings.HasPrefix(query, "panic") == true:
		panic("test.panic....")

	case strings.HasPrefix(query, "error") == true:
		return nil, NewSQLError(ERUnknownComError, SSUnknownComError, "query.error: %v", query)
	}

	return th.qr, nil
}

func (th *testHandler) setResult(r *sqltypes.Result) {
	th.qr = r
}

func TestServer(t *testing.T) {
	th := newTestHandler()
	result1 := &sqltypes.Result{
		Fields: []*querypb.Field{
			{
				Name: "id",
				Type: querypb.Type_INT32,
			},
			{
				Name: "name",
				Type: querypb.Type_VARCHAR,
			},
		},
		Rows: [][]sqltypes.Value{
			{
				sqltypes.MakeTrusted(querypb.Type_INT32, []byte("10")),
				sqltypes.MakeTrusted(querypb.Type_VARCHAR, []byte("nice name")),
			},
			{
				sqltypes.MakeTrusted(querypb.Type_INT32, []byte("20")),
				sqltypes.NULL,
			},
		},
	}
	result2 := &sqltypes.Result{}

	th.setResult(result1)

	address := fmt.Sprintf(":%d", randomPort(8000, 9000))
	server, err := NewListener(address, th)
	assert.Nil(t, err)

	go func() {
		server.Accept()
	}()

	// query
	{
		client, err := NewConn("mock", "mock", address, "test")
		assert.Nil(t, err)
		_, err = client.Query("select")
		assert.Nil(t, err)
		client.Close()
	}

	// query1
	{
		th.setResult(result2)
		client, err := NewConn("mock", "mock", address, "test")
		assert.Nil(t, err)
		_, err = client.Query("select")
		assert.Nil(t, err)
		client.Close()
		th.setResult(result1)
	}

	// exec
	{
		client, err := NewConn("mock", "mock", address, "test")
		assert.Nil(t, err)
		err = client.Exec("exec")
		assert.Nil(t, err)
		client.Close()
	}

	// fetch all
	{
		client, err := NewConn("mock", "mock", address, "test")
		assert.Nil(t, err)
		results, err := client.FetchAll("select", -1)
		assert.Nil(t, err)
		assert.Equal(t, result1, results)
		client.Close()
	}

	// fetch one
	{
		client, err := NewConn("mock", "mock", address, "test")
		assert.Nil(t, err)
		results, err := client.FetchAll("select", 1)
		assert.Nil(t, err)
		want := 1
		got := len(results.Rows)
		assert.Equal(t, want, got)
		client.Close()
	}

	// error
	{
		client, err := NewConn("mock", "mock", address, "test")
		assert.Nil(t, err)
		err = client.Exec("error")
		want := "query.error: error"
		got := err.Error()
		assert.Equal(t, want, got)
	}

	// panic
	{
		client, err := NewConn("mock", "mock", address, "test")
		assert.Nil(t, err)
		client.Exec("panic")
		want := true
		got := client.Closed()
		assert.Equal(t, want, got)
	}

	// ping
	{
		client, err := NewConn("mock", "mock", address, "test")
		assert.Nil(t, err)
		err = client.Ping()
		assert.Nil(t, err)
	}

	// init db
	{
		client, err := NewConn("mock", "mock", address, "test")
		assert.Nil(t, err)
		err = client.InitDB("fxk")
		assert.Nil(t, err)
	}

	// auth denied
	{
		_, err := NewConn("mockx", "mock", address, "test")
		want := "Access denied for user 'mockx'"
		got := err.Error()
		assert.Equal(t, want, got)
	}
}
