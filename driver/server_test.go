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
	"sync"
	"testing"

	"github.com/XeLabs/go-mysqlstack/xlog"
	"github.com/stretchr/testify/assert"

	querypb "github.com/XeLabs/go-mysqlstack/sqlparser/depends/query"
	"github.com/XeLabs/go-mysqlstack/sqlparser/depends/sqltypes"
)

func TestServer(t *testing.T) {
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

	log := xlog.NewStdLog(xlog.Level(xlog.DEBUG))
	th := NewTestHandler(log)
	port, svr, err := MockMysqlServer(log, th)
	assert.Nil(t, err)
	defer svr.Close()
	address := fmt.Sprintf(":%v", port)

	// query
	{
		client, err := NewConn("mock", "mock", address, "test")
		assert.Nil(t, err)
		defer client.Close()

		th.SetCond(&Cond{Query: "SELECT1", Result: result1})
		_, err = client.Query("SELECT1")
		assert.Nil(t, err)
	}

	// query1
	{
		client, err := NewConn("mock", "mock", address, "test")
		assert.Nil(t, err)

		th.SetCond(&Cond{Query: "SELECT2", Result: result2})
		_, err = client.Query("SELECT2")
		assert.Nil(t, err)
		client.Close()
	}

	// exec
	{
		client, err := NewConn("mock", "mock", address, "test")
		assert.Nil(t, err)
		defer client.Close()

		th.SetCond(&Cond{Query: "SELECT1", Result: result1})
		err = client.Exec("SELECT1")
		assert.Nil(t, err)
	}

	// fetch all
	{
		client, err := NewConn("mock", "mock", address, "test")
		assert.Nil(t, err)
		defer client.Close()

		th.SetCond(&Cond{Query: "SELECT1", Result: result1})
		r, err := client.FetchAll("SELECT1", -1)
		assert.Nil(t, err)
		assert.Equal(t, result1, r)
	}

	// fetch one
	{
		client, err := NewConn("mock", "mock", address, "test")
		assert.Nil(t, err)

		th.SetCond(&Cond{Query: "SELECT1", Result: result1})
		r, err := client.FetchAll("SELECT1", 1)
		assert.Nil(t, err)
		defer client.Close()

		want := 1
		got := len(r.Rows)
		assert.Equal(t, want, got)
	}

	// error
	{
		client, err := NewConn("mock", "mock", address, "test")
		assert.Nil(t, err)
		defer client.Close()

		th.SetCond(&Cond{Query: "ERROR1", Error: NewSQLError(ERUnknownComError, SSUnknownComError, "query.error")})
		err = client.Exec("ERROR1")
		assert.NotNil(t, err)
		want := "query.error"
		got := err.Error()
		assert.Equal(t, want, got)
	}

	// panic
	{
		client, err := NewConn("mock", "mock", address, "test")
		assert.Nil(t, err)
		defer client.Close()

		th.SetCond(&Cond{Query: "PANIC", Panic: true})
		client.Exec("PANIC")

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

func TestServerSessionClose(t *testing.T) {
	result2 := &sqltypes.Result{}

	log := xlog.NewStdLog(xlog.Level(xlog.DEBUG))
	th := NewTestHandler(log)
	port, _, err := MockMysqlServer(log, th)
	assert.Nil(t, err)
	address := fmt.Sprintf(":%v", port)

	{
		// create session 1
		client1, err := NewConn("mock", "mock", address, "test")
		assert.Nil(t, err)

		th.SetCond(&Cond{Query: "SELECT2", Result: result2})
		r, err := client1.FetchAll("SELECT2", -1)
		assert.Nil(t, err)
		assert.Equal(t, result2, r)

		// kill session 1
		client2, err := NewConn("mock", "mock", address, "test")
		assert.Nil(t, err)
		_, err = client2.Query("KILL 1")
		assert.Nil(t, err)
	}
}

func TestServerSessionKilled(t *testing.T) {
	result2 := &sqltypes.Result{}

	log := xlog.NewStdLog(xlog.Level(xlog.DEBUG))
	th := NewTestHandler(log)
	port, _, err := MockMysqlServer(log, th)
	assert.Nil(t, err)
	address := fmt.Sprintf(":%v", port)

	{
		// create session 1
		client1, err := NewConn("mock", "mock", address, "test")
		assert.Nil(t, err)

		var wg sync.WaitGroup
		th.SetCond(&Cond{Query: "SELECT2", Result: result2, Delay: 10})
		wg.Add(1)
		go func() {
			_, err := client1.FetchAll("SELECT2", -1)
			assert.NotNil(t, err)
			wg.Done()
		}()

		// kill session 1
		client2, err := NewConn("mock", "mock", address, "test")
		assert.Nil(t, err)
		_, err = client2.Query("KILL 1")
		assert.Nil(t, err)
		wg.Wait()
	}
}
