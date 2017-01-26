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
)

func TestConn(t *testing.T) {
	port := randomPort(8000, 9000)
	address := fmt.Sprintf(":%d", port)

	mysql, cleanup := NewMockServer(port)
	defer cleanup()
	mysql.Start()

	// greeting error
	{
		mysql.SetError(ERR_GREETING)
		conn, err := NewConn("mock", "mock", address, "test")
		assert.NotNil(t, err)
		conn.Close()
		mysql.SetError(ERR_NONE)
	}

	// auth error
	{
		mysql.SetError(ERR_AUTH)
		conn, err := NewConn("mock", "mock", address, "test")
		assert.NotNil(t, err)
		conn.Close()
		mysql.SetError(ERR_NONE)
	}

	// open ok
	{
		conn, err := NewConn("mock", "mock", address, "test")
		assert.Nil(t, err)
		conn.Close()

		want := true
		got := conn.Closed()
		assert.Equal(t, want, got)
	}
}
