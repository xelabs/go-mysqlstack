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
	"net"
	"runtime"
	"strings"

	"github.com/XeLabs/go-mysqlstack/common"
	"github.com/XeLabs/go-mysqlstack/consts"
	"github.com/XeLabs/go-mysqlstack/packet"
	"github.com/XeLabs/go-mysqlstack/proto"

	"github.com/XeLabs/go-mysqlstack/sqlparser/depends/sqltypes"
)

type Handler interface {
	AuthCheck(session *Session) bool
	ComQuery(session *Session, query string) (*sqltypes.Result, error)
}

type Listener struct {
	// Query handler.
	handler Handler

	// This is the main listener socket.
	listener net.Listener

	// Incrementing ID for connection id.
	connectionID uint64

	// sessions maps all sessions.
	sessions map[uint64]*Session

	// PasswordMap maps users to passwords.
	PasswordMap map[string]string
}

// NewListener creates a new Listener.
func NewListener(address string, handler Handler) (*Listener, error) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return nil, err
	}

	return &Listener{
		PasswordMap:  make(map[string]string),
		sessions:     make(map[uint64]*Session),
		handler:      handler,
		listener:     listener,
		connectionID: 1,
	}, nil
}

// Accept runs an accept loop until the listener is closed.
func (l *Listener) Accept() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	for {
		conn, err := l.listener.Accept()
		if err != nil {
			// Close() was probably called.
			return
		}

		ID := l.connectionID
		l.connectionID++
		go l.handle(conn, ID)
	}
}

// handle is called in a go routine for each client connection.
func (l *Listener) handle(conn net.Conn, ID uint64) {
	var err error
	var data []byte
	var authPkt []byte
	var greetingPkt []byte

	// session
	session := NewSession(conn)

	// Catch panics, and close the connection in any case.
	defer func() {
		if x := recover(); x != nil {
			fmt.Sprintf("mysql_server caught panic:\n%v\n", x)
		}
		conn.Close()
	}()

	// greeting packet
	greetingPkt = session.Greeting.Pack()
	if err = session.Packets.Write(greetingPkt); err != nil {
		return
	}

	// auth packet
	if authPkt, err = session.Packets.Next(); err != nil {
		return
	}
	if err = session.Auth.UnPack(authPkt); err != nil {
		return
	}

	//  auth check
	if !l.handler.AuthCheck(session) {
		session.Packets.WriteERR(ERAccessDeniedError, SSAccessDeniedError, "Access denied for user '%v'", session.Auth.User())
		return
	} else {
		if err = session.Packets.WriteOK(0, 0, session.Greeting.Status(), 0); err != nil {
			return
		}
	}

	for {
		// reset packet sequence ID
		session.Packets.ResetSeq()
		if data, err = session.Packets.Next(); err != nil {
			return
		}

		switch data[0] {
		case consts.COM_QUIT:
			return
		case consts.COM_INIT_DB:
			session.Schema = string(data[1:])
			if err = session.Packets.WriteOK(0, 0, session.Greeting.Status(), 0); err != nil {
				return
			}
		case consts.COM_QUERY:
			var result *sqltypes.Result
			query := strings.TrimRight(common.BytesToString(data[1:]), ";")
			if result, err = l.handler.ComQuery(session, query); err != nil {
				if werr := session.writeErrFromError(err); werr != nil {
					return
				}
				continue
			}
			if err = session.writeResult(result); err != nil {
				return
			}
		case consts.COM_PING:
			if err = session.Packets.WriteOK(0, 0, session.Greeting.Status(), 0); err != nil {
				return
			}
		default:
			if err = session.Packets.WriteERR(ERUnknownComError, SSUnknownComError, "command handling not implemented yet: %v", data[0]); err != nil {
				return
			}
		}
		// reset packet sequence ID
		session.Packets.ResetSeq()
	}
}

type Session struct {
	Schema   string
	Auth     *proto.Auth
	Packets  *packet.Packets
	Greeting *proto.Greeting
}

func NewSession(conn net.Conn) *Session {
	return &Session{
		Auth:     proto.NewAuth(),
		Greeting: proto.NewGreeting(0),
		Packets:  packet.NewPackets(conn),
	}
}

func (s *Session) writeErrFromError(err error) error {
	if se, ok := err.(*SQLError); ok {
		return s.Packets.WriteERR(uint16(se.Num), se.State, "%v", se.Message)
	}

	return s.Packets.WriteERR(ERUnknownError, SSSignalException, "unknown error: %v", err)
}

func (s *Session) writeResult(result *sqltypes.Result) error {
	if len(result.Fields) == 0 {
		// This is just an INSERT result, send an OK packet.
		return s.Packets.WriteOK(result.RowsAffected, result.InsertID, s.Greeting.Status(), 0)
	}

	// 1. Write columns
	if err := s.Packets.WriteColumns(result.Fields); err != nil {
		return err
	}

	// 2. Append rows
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

	// 3. Write EOF
	if err := s.Packets.AppendEOF(batch); err != nil {
		return err
	}

	// 4. Write to stream
	if err := s.Packets.Flush(batch.Datas()); err != nil {
		return err
	}

	return nil
}
