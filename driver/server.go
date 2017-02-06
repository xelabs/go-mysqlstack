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
	"runtime"
	"strings"

	"github.com/XeLabs/go-mysqlstack/common"
	"github.com/XeLabs/go-mysqlstack/consts"
	"github.com/XeLabs/go-mysqlstack/xlog"

	"github.com/XeLabs/go-mysqlstack/sqlparser/depends/sqltypes"
)

type Handler interface {
	// Register session.
	Register(session *Session)

	// UnRegister session, used when session exit.
	UnRegister(session *Session)

	// Check the client ip.
	IPCheck(ip string) bool

	// Check the Auth request.
	AuthCheck(session *Session) bool

	// Handle the queries.
	ComQuery(session *Session, query string) (*sqltypes.Result, error)
}

type Listener struct {
	// Logger.
	log *xlog.Log

	// Query handler.
	handler Handler

	// This is the main listener socket.
	listener net.Listener

	// Incrementing ID for connection id.
	connectionID uint32

	// sessions maps all sessions.
	sessions map[uint64]*Session

	// PasswordMap maps users to passwords.
	PasswordMap map[string]string
}

// NewListener creates a new Listener.
func NewListener(log *xlog.Log, address string, handler Handler) (*Listener, error) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return nil, err
	}

	return &Listener{
		log:          log,
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
func (l *Listener) handle(conn net.Conn, ID uint32) {
	var err error
	var data []byte
	var authPkt []byte
	var greetingPkt []byte

	addr := conn.RemoteAddr().String()
	if !l.handler.IPCheck(addr) {
		l.log.Warning("ip[%v].check.failed", ID, addr)
		return
	}

	session := NewSession(ID, conn)
	l.handler.Register(session)

	// Catch panics, and close the connection in any case.
	defer func() {
		l.log.Warning("session[%v].client[%v].exit.error[%+v]...", ID, addr, err)
		conn.Close()
		l.handler.UnRegister(session)
		if x := recover(); x != nil {
			buf := make([]byte, 1024)
			buf = buf[:runtime.Stack(buf, false)]
			l.log.Error("server.handle.panic:%v\n%s", x, buf)
		}
	}()

	// greeting packet.
	greetingPkt = session.Greeting.Pack()
	if err = session.Packets.Write(greetingPkt); err != nil {
		return
	}

	// auth packet.
	if authPkt, err = session.Packets.Next(); err != nil {
		return
	}
	if err = session.Auth.UnPack(authPkt); err != nil {
		return
	}

	//  auth check.
	if !l.handler.AuthCheck(session) {
		l.log.Warning("session.user[%+v].auth.check.failed", session.Auth.User())
		session.Packets.WriteERR(ERAccessDeniedError, SSAccessDeniedError, "Access denied for user '%v'", session.Auth.User())
		return
	} else {
		if err = session.Packets.WriteOK(0, 0, session.Greeting.Status(), 0); err != nil {
			return
		}
	}

	for {
		// reset packet sequence ID.
		session.Packets.ResetSeq()
		if data, err = session.Packets.Next(); err != nil {
			l.log.Error("session[%v].read.next:%+v", ID, err)
			return
		}

		switch data[0] {
		case consts.COM_QUIT:
			l.log.Debug("session[%v].com.quit", ID)
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
				l.log.Error("session[%v].query.error:%+v", ID, err)
				if werr := session.writeErrFromError(err); werr != nil {
					err = werr
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
			l.log.Error("session.command:%v.not.implemented:%+v", data[0])
			if err = session.Packets.WriteERR(ERUnknownComError, SSUnknownComError, "command handling not implemented yet: %v", data[0]); err != nil {
				return
			}
		}
		// reset packet sequence ID.
		session.Packets.ResetSeq()
	}
}

// Close close the listener and all connections.
func (l *Listener) Close() {
	l.listener.Close()
}
