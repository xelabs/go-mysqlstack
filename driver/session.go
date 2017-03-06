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
	"net"

	"github.com/XeLabs/go-mysqlstack/common"
	"github.com/XeLabs/go-mysqlstack/packet"
	"github.com/XeLabs/go-mysqlstack/proto"
	"github.com/XeLabs/go-mysqlstack/sqldb"

	"github.com/XeLabs/go-mysqlstack/sqlparser/depends/sqltypes"
)

type Session struct {
	ID       uint32
	conn     net.Conn
	Schema   string
	Auth     *proto.Auth
	Packets  *packet.Packets
	Greeting *proto.Greeting
}

func newSession(ID uint32, conn net.Conn) *Session {
	return &Session{
		ID:       ID,
		conn:     conn,
		Auth:     proto.NewAuth(),
		Greeting: proto.NewGreeting(ID),
		Packets:  packet.NewPackets(conn),
	}
}

func (s *Session) writeErrFromError(err error) error {
	if se, ok := err.(*sqldb.SQLError); ok {
		return s.Packets.WriteERR(se.Num, se.State, "%v", se.Message)
	}
	unknow := sqldb.NewSQLError(sqldb.ER_UNKNOWN_ERROR, "%v", err)
	return s.Packets.WriteERR(unknow.Num, unknow.State, unknow.Message)
}

func (s *Session) writeResult(result *sqltypes.Result) error {
	if len(result.Fields) == 0 {
		// This is just an INSERT result, send an OK packet.
		return s.Packets.WriteOK(result.RowsAffected, result.InsertID, s.Greeting.Status(), 0)
	}

	// 1. Write columns.
	if err := s.Packets.WriteColumns(result.Fields); err != nil {
		return err
	}

	// 2. Append rows.
	batch := common.NewBuffer(64)
	for _, row := range result.Rows {
		rowBuf := common.NewBuffer(16)
		for _, val := range row {
			if val.IsNull() {
				rowBuf.WriteLenEncodeNUL()
			} else {
				rowBuf.WriteLenEncodeBytes(val.Raw())
			}
		}
		if err := s.Packets.Append(batch, rowBuf.Datas()); err != nil {
			return err
		}
	}

	// 3. Write EOF.
	if err := s.Packets.AppendEOF(batch); err != nil {
		return err
	}

	// 4. Write to stream.
	if err := s.Packets.Flush(batch.Datas()); err != nil {
		return err
	}

	return nil
}

func (s *Session) Addr() string {
	return s.conn.RemoteAddr().String()
}

func (s *Session) Close() {
	if s.conn != nil {
		s.conn.Close()
		s.conn = nil
	}
}
