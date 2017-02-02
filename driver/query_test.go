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
	//	"fmt"
	"testing"
	//	"github.com/XeLabs/go-mysqlstack/proto"
	//	"github.com/stretchr/testify/assert"
)

func TestQuery(t *testing.T) {
	/*
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

			// mysql cient
			conn, err := NewConn("mock", "mock", address, "test")
			assert.Nil(t, err)
			defer conn.Close()

			// query
			{
				var got int

				mysql.SetResult(result1)
				rows, _ := conn.Query("select * from mock")
				want := 2
				for rows.Next() {
					got++
				}
				err := rows.LastError()
				assert.Nil(t, err)
				assert.Equal(t, want, got)
			}

			// exec
			{
				var got int

				mysql.SetResult(result1)
				rows, _ := conn.Query("use xx")
				want := 0
				for rows.Next() {
					got++
				}
				err := rows.LastError()
				assert.Nil(t, err)
				assert.Equal(t, want, got)
			}
		}

		func TestQueryExec(t *testing.T) {
			port := randomPort(8000, 9000)
			address := fmt.Sprintf(":%d", port)

			result1 := &MockResult{
				Fields: []*proto.Column{
					{Name: "id"},
					{Name: "col"},
				},
				AffectedRows: 2015,
				LastInsertID: 2015,
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

			// mysql cient
			conn, err := NewConn("mock", "mock", address, "test")
			assert.Nil(t, err)
			defer conn.Close()

			// exec
			{
				mysql.SetResult(result1)
				rows, _ := conn.Exec("use db")
				want := 2015
				got := int(rows.RowsAffected())
				err := rows.LastError()
				assert.Nil(t, err)
				assert.Equal(t, want, got)
			}

			// exec
			{
				mysql.SetResult(result1)
				rows, _ := conn.Exec("select * from xx")
				err := rows.LastError()
				assert.Nil(t, err)
			}
	*/
}
