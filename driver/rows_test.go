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

	querypb "github.com/XeLabs/go-mysqlstack/sqlparser/depends/query"
	"github.com/XeLabs/go-mysqlstack/sqlparser/depends/sqltypes"
)

func TestRows(t *testing.T) {
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
	result2 := &sqltypes.Result{
		RowsAffected: 123,
		InsertID:     123456789,
	}

	address := fmt.Sprintf(":%d", randomPort(8000, 9000))
	server, err := NewListener(address, th)
	assert.Nil(t, err)

	go func() {
		server.Accept()
	}()

	// query
	{
		th.setResult(result2)
		client, err := NewConn("mock", "mock", address, "test")
		assert.Nil(t, err)
		defer client.Close()

		rows, err := client.Query("select")
		assert.Nil(t, err)

		assert.Equal(t, uint64(123), rows.RowsAffected())
		assert.Equal(t, uint64(123456789), rows.LastInsertID())
	}

	// query
	{
		th.setResult(result1)
		client, err := NewConn("mock", "mock", address, "test")
		assert.Nil(t, err)
		defer client.Close()

		rows, err := client.Query("select")
		assert.Nil(t, err)
		assert.Equal(t, result1.Fields, rows.Fields())
		for rows.Next() {
			_ = rows.Datas()
			_, _ = rows.RowValues()
		}
	}
}
