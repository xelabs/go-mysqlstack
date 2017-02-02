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
)

func TestClient(t *testing.T) {
	th := newTestHandler()
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
}
