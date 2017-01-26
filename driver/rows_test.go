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

	"github.com/XeLabs/go-mysqlstack/proto"
	"github.com/stretchr/testify/assert"
)

func TestRowsNext(t *testing.T) {
	port := randomPort(8000, 9000)
	address := fmt.Sprintf(":%d", port)

	result1 := &MockResult{
		Fields: []*proto.Column{
			{Name: "id"},
			{Name: "col"},
		},
		AffectedRows: 0,
		LastInsertID: 0,
		Rows: [][][]byte{
			{
				[]byte("11"),
				[]byte("12"),
			},
			{
				[]byte("21"),
				[]byte("22"),
			},
		},
	}

	// mysql server
	mysql, cleanup := NewMockServer(port)
	mysql.Start()
	defer cleanup()

	// mysql client
	conn, err := NewConn("mock", "mock", address, "test")
	assert.Nil(t, err)
	defer conn.Close()

	// query
	{
		mysql.SetResult(result1)
		rows, _ := conn.Query("select * from mock")
		for rows.Next() {
			assert.NotNil(t, rows.Datas())
		}
		err := rows.LastError()
		assert.Nil(t, err)

		fields := rows.Fields()
		assert.Equal(t, result1.Fields, fields)

		rows.RowsAffected()
		rows.InsertID()
	}
}
