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
	"errors"
	"fmt"
	"math/rand"
	"net"
	"runtime"
	"strings"
	"time"

	"github.com/XeLabs/go-mysqlstack/common"
	"github.com/XeLabs/go-mysqlstack/consts"
	"github.com/XeLabs/go-mysqlstack/packet"
	"github.com/XeLabs/go-mysqlstack/proto"
)

type MockResult struct {
	AffectedRows uint64
	LastInsertID uint64
	Fields       []*proto.Column
	Rows         [][][]byte
}

type Session struct {
	errType  int
	r        *MockResult
	greeting *proto.Greeting
	auth     *proto.Auth
	packets  *packet.Packets
}

func NewSession(errType int, c net.Conn) *Session {
	s := &Session{
		errType:  errType,
		greeting: proto.NewGreeting(0),
		auth:     proto.NewAuth(),
		packets:  packet.NewPackets(c),
	}

	return s
}

func (s *Session) Kickoff() error {
	if err := s.handleGreeting(); err != nil {
		return err
	}

	if err := s.handleAuth(); err != nil {
		return err
	}

	if err := s.handleCommand(); err != nil {
		return err
	}

	return nil
}

func (s *Session) handleGreeting() error {
	if s.errType == ERR_GREETING {
		msg := "mock.mysql.greeting.error"
		s.sendERR(msg)
		return errors.New(msg)
	} else {
		body := s.greeting.Pack()
		if err := s.packets.Write(body); err != nil {
			return err
		}
	}

	return nil
}

func (s *Session) handleAuth() error {
	body, err := s.packets.Next()
	if err != nil {
		return err
	}

	// unpack the auth packet
	if err := s.auth.UnPack(body); err != nil {
		return err
	}

	if s.errType == ERR_AUTH {
		msg := "mock.mysql.auth.error"
		s.sendERR(msg)
		return errors.New(msg)
	} else {
		ok := proto.NewOK()
		if err := s.sendOK(ok); err != nil {
			return err
		}
	}

	return nil
}

func (s *Session) sendResult() error {
	if s.r == nil {
		return nil
	}

	if err := s.packets.WriteColumns(s.r.Fields); err != nil {
		return err
	}

	batch := common.NewBuffer(64)
	for _, row := range s.r.Rows {
		buf := common.NewBuffer(64)
		for _, value := range row {
			buf.WriteLenEncodeBytes([]byte(value))
		}
		s.packets.Append(batch, buf.Datas())
	}
	// EOF
	s.packets.Append(batch, []byte{0xfe})

	// flush all
	s.packets.Flush(batch.Datas())
	s.packets.ResetSeq()

	return nil
}

func (s *Session) handleCommand() error {
	for {
		// reset packet sequence ID
		s.packets.ResetSeq()

		body, err := s.packets.Next()
		if err != nil {
			return err
		}

		sql := strings.TrimRight(common.BytesToString(body[1:]), ";")
		if strings.HasPrefix(sql, "select") {
			if err := s.sendResult(); err != nil {
				return err
			}
		} else {
			ok := proto.NewOK()
			if s.r != nil {
				ok.AffectedRows = s.r.AffectedRows
				ok.LastInsertID = s.r.LastInsertID
			}
			s.sendOK(ok)
		}

		// reset packet sequence ID
		s.packets.ResetSeq()
	}
}

// https://dev.mysql.com/doc/internals/en/packet-OK_Packet.html
func (s *Session) sendOK(ok *proto.OK) error {
	cflags := s.auth.ClientFlags()
	status := s.greeting.Status()
	buf := common.NewBuffer(64)

	// OK
	buf.WriteU8(proto.OK_PACKET)

	// affected rows
	buf.WriteLenEncode(ok.AffectedRows)

	// last insert id
	buf.WriteLenEncode(ok.LastInsertID)
	if (cflags & consts.CLIENT_PROTOCOL_41) > 0 {
		// status
		buf.WriteU16(uint16(status))
		// warnings
		buf.WriteU16(ok.Warnings)
	}

	if err := s.packets.Write(buf.Datas()); err != nil {
		return err
	}

	return nil
}

func (s *Session) sendERR(msg string) error {
	cflags := s.auth.ClientFlags()

	buf := common.NewBuffer(64)
	buf.WriteU8(proto.ERR_PACKET)
	buf.WriteU16(101)
	if (cflags & consts.CLIENT_PROTOCOL_41) > 0 {
		// sql-state marker #
		buf.WriteU8('#')
		// sql-state (?) 5 ascii bytes
		buf.WriteString("10001")
	}
	buf.WriteString(msg)

	if err := s.packets.Write(buf.Datas()); err != nil {
		return err
	}

	return nil
}

func (s *Session) sendResults(rows Rows) error {
	batch := common.NewBuffer(64)
	// 1. write columns
	if err := s.packets.WriteColumns(rows.Fields()); err != nil {
		return err
	}

	// 2. append rows
	for rows.Next() {
		if err := s.packets.Append(batch, rows.Datas()); err != nil {
			return err
		}
	}
	if err := rows.LastError(); err != nil {
		return err
	}
	// rows EOF
	if err := s.packets.AppendEOF(batch); err != nil {
		return err
	}

	// 3. write to stream
	if err := s.packets.Flush(batch.Datas()); err != nil {
		return err
	}

	return nil
}

func randomPort(min int, max int) int {
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	d, delta := min, (max - min)
	if delta > 0 {
		d += rand.Intn(int(delta))
	}
	return d
}

// for test only
func errOK(err error) {
	if err != nil {
		panic(err)
	}
}

const (
	ERR_NONE int = iota
	ERR_GREETING
	ERR_AUTH
)

// MySQL server mocker
type MockServer struct {
	port     int
	errType  int
	ID       uint32
	cleanup  bool
	err      string
	l        net.Listener
	sessions map[uint32]*Session
}

func NewMockServer(port int) (*MockServer, func()) {
	m := &MockServer{
		port:     port,
		sessions: make(map[uint32]*Session),
	}

	return m, func() {
		m.cleanup = true
		m.l.Close()
	}
}

func (m *MockServer) SetError(errType int) {
	m.errType = errType
}

func (m *MockServer) SetResult(r *MockResult) {
	for _, session := range m.sessions {
		session.r = r
	}
}

func (m *MockServer) Start() {
	var err error
	runtime.GOMAXPROCS(runtime.NumCPU())

	ep := fmt.Sprintf(":%d", m.port)
	m.l, err = net.Listen("tcp", ep)
	errOK(err)
	time.Sleep(time.Second)

	go func() {
		for {
			c, _ := m.l.Accept()
			if m.cleanup {
				return
			}
			go m.handleConnection(c)
		}
	}()

	time.Sleep(time.Second)
}

func (m *MockServer) handleConnection(c net.Conn) {
	defer c.Close()
	session := NewSession(m.errType, c)
	m.sessions[m.ID] = session
	m.ID = m.ID + 1

	if err := session.Kickoff(); err != nil {
		if err.Error() != "EOF" {
			fmt.Printf("mockserver.kickoff.error[%+v]\n", err)
		}
	}
}
