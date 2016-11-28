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
	"github.com/pkg/errors"
)

// readColumns read packets as Field Packets until EOF-Packet or an Error appears
func (c *Conn) readColumns() (columns []proto.Column, err error) {
	var count uint64
	var payload []byte
	var pkt *proto.ERR

	if payload, err = c.packets.Next(); err != nil {
		return
	}

	switch payload[0] {
	case proto.OK_PACKET:
		// maybe we are exec OK response
		if _, err = c.packets.ParseOK(payload, c.greeting.Capability); err != nil {
			return
		}
		return

	case proto.ERR_PACKET:
		if pkt, err = c.packets.ParseERR(payload, c.greeting.Capability); err != nil {
			return
		}
		err = errors.New(pkt.ErrorMessage)
		return
	}

	// column info unpack
	if count, err = proto.ColumnCount(payload); err != nil {
		return
	}

	columns = make([]proto.Column, count)
	for i := 0; i < int(count); i++ {
		if payload, err = c.packets.Next(); err != nil {
			return
		}

		if err = columns[i].UnPack(payload); err != nil {
			return
		}
	}

	// EOF packet
	if payload, err = c.packets.Next(); err != nil {
		return
	} else {
		if payload[0] != proto.EOF_PACKET {
			err = errors.Errorf("read.columns.EOF.error.collen[%v]", columns)
			return
		}
	}

	return
}

func (c *Conn) Query(sql string) (*Rows, error) {
	// query request
	if err := c.packets.WriteCommand(consts.COM_QUERY,
		common.StringToBytes(sql)); err != nil {
		return nil, err
	}

	// query response
	return c.QueryResponse()
}

// http://dev.mysql.com/doc/internals/en/com-query-response.html#packet-ProtocolText::Resultset
func (c *Conn) QueryResponse() (rows *Rows, err error) {
	rows = NewRows(c.packets)
	if rows.columns, err = c.readColumns(); err != nil {
		return
	}

	return
}

func (c *Conn) Exec(sql string) (*proto.OK, error) {
	// query request
	if err := c.packets.WriteCommand(consts.COM_QUERY,
		common.StringToBytes(sql)); err != nil {
		return nil, err
	}

	// exec response
	return c.ExecResponse()
}

// http://dev.mysql.com/doc/internals/en/com-query-response.html#packet-ProtocolText::Resultset
func (c *Conn) ExecResponse() (ok *proto.OK, err error) {
	rows := NewRows(c.packets)
	if rows.columns, err = c.readColumns(); err != nil {
		return
	}

	if len(rows.columns) > 0 {
		// read until EOF
		for rows.Next() {
		}
		if err = rows.LastError(); err != nil {
			return
		}
	}

	ok = &proto.OK{}
	return
}
