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
	"testing"

	"github.com/stretchr/testify/assert"

	querypb "github.com/XeLabs/go-mysqlstack/sqlparser/depends/query"
	"github.com/XeLabs/go-mysqlstack/sqlparser/depends/sqltypes"
	"github.com/XeLabs/go-mysqlstack/xlog"
)

func TestClient(t *testing.T) {
	result2 := &sqltypes.Result{
		RowsAffected: 123,
		InsertID:     123456789,
	}

	log := xlog.NewStdLog(xlog.Level(xlog.ERROR))
	th := NewTestHandler(log)
	svr, err := MockMysqlServer(log, th)
	assert.Nil(t, err)
	defer svr.Close()
	address := svr.Addr()

	// query
	{

		client, err := NewConn("mock", "mock", address, "test", "")
		assert.Nil(t, err)
		defer client.Close()

		// connection ID
		assert.Equal(t, uint32(1), client.ConnectionID())

		th.AddQuery("SELECT2", result2)
		rows, err := client.Query("SELECT2")
		assert.Nil(t, err)

		assert.Equal(t, uint64(123), rows.RowsAffected())
		assert.Equal(t, uint64(123456789), rows.LastInsertID())
	}
}

func TestClientClosed(t *testing.T) {
	result2 := &sqltypes.Result{}

	log := xlog.NewStdLog(xlog.Level(xlog.ERROR))
	th := NewTestHandler(log)
	svr, err := MockMysqlServer(log, th)
	assert.Nil(t, err)
	defer svr.Close()
	address := svr.Addr()

	{
		// create session 1
		client1, err := NewConn("mock", "mock", address, "test", "")
		assert.Nil(t, err)

		th.AddQuery("SELECT2", result2)
		r, err := client1.FetchAll("SELECT2", -1)
		assert.Nil(t, err)
		assert.Equal(t, result2, r)

		// kill session 1
		client2, err := NewConn("mock", "mock", address, "test", "")
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

func TestClientFetchAllWithFunc(t *testing.T) {
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

	log := xlog.NewStdLog(xlog.Level(xlog.ERROR))
	th := NewTestHandler(log)
	svr, err := MockMysqlServer(log, th)
	assert.Nil(t, err)
	defer svr.Close()
	address := svr.Addr()

	// query
	{

		client, err := NewConn("mock", "mock", address, "test", "")
		assert.Nil(t, err)
		defer client.Close()

		th.AddQuery("SELECT2", result1)
		checkFunc := func(rows Rows) error {
			if rows.Bytes() > 2 {
				return errors.New("client.checkFunc.error")
			}
			return nil
		}
		_, err = client.FetchAllWithFunc("SELECT2", -1, checkFunc)
		want := "client.checkFunc.error"
		got := err.Error()
		assert.Equal(t, want, got)
	}
}
