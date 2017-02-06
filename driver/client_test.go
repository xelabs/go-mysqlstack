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
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/XeLabs/go-mysqlstack/sqlparser/depends/sqltypes"
	"github.com/XeLabs/go-mysqlstack/xlog"
)

func TestClient(t *testing.T) {
	result2 := &sqltypes.Result{
		RowsAffected: 123,
		InsertID:     123456789,
	}

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

		th.SetCond(&Cond{Query: "SELECT2", Result: result2})
		rows, err := client.Query("SELECT2")
		assert.Nil(t, err)

		assert.Equal(t, uint64(123), rows.RowsAffected())
		assert.Equal(t, uint64(123456789), rows.LastInsertID())
	}
}

func TestClientClosed(t *testing.T) {
	result2 := &sqltypes.Result{}

	log := xlog.NewStdLog(xlog.Level(xlog.DEBUG))
	th := NewTestHandler(log)
	port, svr, err := MockMysqlServer(log, th)
	assert.Nil(t, err)
	defer svr.Close()
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

		// check client1 connection
		err = client1.Ping()
		assert.NotNil(t, err)
		want := true
		got := client1.Closed()
		assert.Equal(t, want, got)
	}
}
