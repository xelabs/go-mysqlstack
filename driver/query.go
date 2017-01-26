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
	"github.com/XeLabs/go-mysqlstack/common"
	"github.com/XeLabs/go-mysqlstack/consts"
	"github.com/XeLabs/go-mysqlstack/proto"
)

func (c *conn) Query(sql string) (Rows, error) {
	var err error
	var ok *proto.OK
	var columns []*proto.Column

	// query request
	if err = c.packets.WriteCommand(consts.COM_QUERY,
		common.StringToBytes(sql)); err != nil {
		c.Cleanup()
		return nil, err
	}

	// columns
	if columns, ok, err = c.packets.ReadColumns(c.greeting.Capability); err != nil {
		c.Cleanup()
		return nil, err
	}

	rows := NewTextRows(c)
	rows.rowsAffected = ok.AffectedRows
	rows.insertID = ok.LastInsertID
	rows.fields = columns

	return rows, nil
}

func (c *conn) Exec(sql string) (Rows, error) {
	rows, err := c.Query(sql)
	if err != nil {
		return nil, err
	}

	if err := rows.Close(); err != nil {
		return nil, err
	}

	return rows, nil
}
