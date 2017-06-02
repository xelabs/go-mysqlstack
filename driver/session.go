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
	"sync"

	"github.com/XeLabs/go-mysqlstack/common"
	"github.com/XeLabs/go-mysqlstack/packet"
	"github.com/XeLabs/go-mysqlstack/proto"
	"github.com/XeLabs/go-mysqlstack/sqldb"
	"github.com/XeLabs/go-mysqlstack/xlog"

	"github.com/XeLabs/go-mysqlstack/sqlparser/depends/sqltypes"
)

type Session struct {
	id       uint32
	mu       sync.RWMutex
	log      *xlog.Log
	conn     net.Conn
	schema   string
	auth     *proto.Auth
	packets  *packet.Packets
	greeting *proto.Greeting
}

func newSession(log *xlog.Log, ID uint32, conn net.Conn) *Session {
	return &Session{
		id:       ID,
		log:      log,
		conn:     conn,
		auth:     proto.NewAuth(),
		greeting: proto.NewGreeting(ID),
		packets:  packet.NewPackets(conn),
	}
}

func (s *Session) writeErrFromError(err error) error {
	if se, ok := err.(*sqldb.SQLError); ok {
		return s.packets.WriteERR(se.Num, se.State, "%v", se.Message)
	}
	unknow := sqldb.NewSQLError(sqldb.ER_UNKNOWN_ERROR, "%v", err)
	return s.packets.WriteERR(unknow.Num, unknow.State, unknow.Message)
}

func (s *Session) writeResult(result *sqltypes.Result) error {
	if len(result.Fields) == 0 {
		// This is just an INSERT result, send an OK packet.
		return s.packets.WriteOK(result.RowsAffected, result.InsertID, s.greeting.Status(), 0)
	}

	// 1. Write columns.
	if err := s.packets.AppendColumns(result.Fields); err != nil {
		return err
	}

	if (s.auth.ClientFlags() & sqldb.CLIENT_DEPRECATE_EOF) == 0 {
		if err := s.packets.AppendEOF(); err != nil {
			return err
		}
	}

	// 2. Append rows.
	for _, row := range result.Rows {
		rowBuf := common.NewBuffer(16)
		for _, val := range row {
			if val.IsNull() {
				rowBuf.WriteLenEncodeNUL()
			} else {
				rowBuf.WriteLenEncodeBytes(val.Raw())
			}
		}
		if err := s.packets.Append(rowBuf.Datas()); err != nil {
			return err
		}
	}

	// 3. Write EOF.
	if (s.auth.ClientFlags() & sqldb.CLIENT_DEPRECATE_EOF) == 0 {
		if err := s.packets.AppendEOF(); err != nil {
			return err
		}
	} else {
		if err := s.packets.AppendOKWithEOFHeader(result.RowsAffected, result.InsertID, s.greeting.Status(), 0); err != nil {
			return err
		}
	}

	// 4. Write to stream.
	if err := s.packets.Flush(); err != nil {
		return err
	}
	return nil
}

func (s *Session) Close() {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.conn != nil {
		s.conn.Close()
		s.conn = nil
	}
}

func (s *Session) ID() uint32 {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.id
}

func (s *Session) Addr() string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.conn != nil {
		return s.conn.RemoteAddr().String()
	} else {
		return "unknow"
	}
}

func (s *Session) SetSchema(schema string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.schema = schema
}

func (s *Session) Schema() string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.schema
}

func (s *Session) User() string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.auth.User()
}

func (s *Session) Salt() []byte {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.greeting.Salt
}

func (s *Session) Scramble() []byte {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.auth.AuthResponse()
}

func (s *Session) Charset() uint8 {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.auth.Charset()
}
