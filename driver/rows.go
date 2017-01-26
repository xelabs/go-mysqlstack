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
	"github.com/XeLabs/go-mysqlstack/proto"
	"github.com/pkg/errors"
)

type Rows interface {
	Next() bool
	Close() error
	Datas() []byte
	RowsAffected() uint64
	InsertID() uint64
	LastError() error
	Fields() []*proto.Column
}

type TextRows struct {
	c            Conn
	err          error
	payload      []byte
	fields       []*proto.Column
	rowsAffected uint64
	insertID     uint64
	buffer       *common.Buffer
}

func NewTextRows(c Conn) *TextRows {
	return &TextRows{
		c:      c,
		fields: make([]*proto.Column, 0, 32),
		buffer: common.NewBuffer(8),
	}
}

// http://dev.mysql.com/doc/internals/en/com-query-response.html#packet-ProtocolText::ResultsetRow
func (r *TextRows) Next() bool {
	// if fields count is 0
	// the packet is OK-Packet without Resultset
	if len(r.fields) == 0 {
		return false
	}

	if r.payload, r.err = r.c.NextPacket(); r.err != nil {
		r.c.Cleanup()
		return false
	}

	switch r.payload[0] {
	case proto.EOF_PACKET:
		return false

	case proto.ERR_PACKET:
		r.err = errors.New("next.packet.error")
		return false
	}
	r.buffer.Reset(r.payload)

	return true
}

func (r *TextRows) Close() error {
	for r.Next() {
	}
	if err := r.LastError(); err != nil {
		return err
	}

	return nil
}

// since packet datas is only guaranteed to be valid
// until the next read, we return cloned bytes
func (r *TextRows) Datas() []byte {
	l := r.buffer.Length()
	data := make([]byte, l)
	copy(data, r.buffer.Datas())

	return data
}

func (r *TextRows) Fields() []*proto.Column {
	return r.fields
}

func (r *TextRows) RowsAffected() uint64 {
	return r.rowsAffected
}

func (r *TextRows) InsertID() uint64 {
	return r.insertID
}

func (r *TextRows) LastError() error {
	return r.err
}
