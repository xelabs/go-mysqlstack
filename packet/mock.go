/*
 * go-mysqlstack
 * xelabs.org
 *
 * Copyright (c) XeLabs
 * Copyright 2016 The Go-MySQL-Driver Authors. All rights reserved.
 * GPL License
 *
 */

package packet

import (
	"io"
	"net"
	"time"
)

var _ net.Conn = &MockConn{}

// struct to mock a net.Conn for testing purposes
type MockConn struct {
	laddr  net.Addr
	raddr  net.Addr
	data   []byte
	closed bool
	read   int
}

func NewMockConn() *MockConn {
	return &MockConn{}
}

func (m *MockConn) Read(b []byte) (n int, err error) {
	// handle the EOF
	if len(m.data) == 0 {
		err = io.EOF
		return
	}

	n = copy(b, m.data)
	m.read += n
	m.data = m.data[n:]

	return
}

func (m *MockConn) Write(b []byte) (n int, err error) {
	m.data = append(m.data, b...)

	return len(b), nil
}

func (m *MockConn) Datas() []byte {
	return m.data
}

func (m *MockConn) Close() error {
	m.closed = true
	return nil
}

func (m *MockConn) LocalAddr() net.Addr {
	return m.laddr
}

func (m *MockConn) RemoteAddr() net.Addr {
	return m.raddr
}

func (m *MockConn) SetDeadline(t time.Time) error {
	return nil
}

func (m *MockConn) SetReadDeadline(t time.Time) error {
	return nil
}

func (m *MockConn) SetWriteDeadline(t time.Time) error {
	return nil
}
