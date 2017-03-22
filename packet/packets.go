/*
 * go-mysqlstack
 * xelabs.org
 *
 * Copyright (c) XeLabs
 * GPL License
 *
 */

package packet

import (
	"fmt"
	"net"

	"github.com/XeLabs/go-mysqlstack/common"
	"github.com/XeLabs/go-mysqlstack/proto"
	"github.com/XeLabs/go-mysqlstack/sqldb"

	querypb "github.com/XeLabs/go-mysqlstack/sqlparser/depends/query"
)

const (
	PACKET_MAX_SIZE = (1<<24 - 1) // (16MB - 1）
)

type Packet struct {
	SequenceID byte
	Datas      []byte
}

type Packets struct {
	seq    uint8
	stream *Stream
}

func NewPackets(c net.Conn) *Packets {
	return &Packets{
		stream: NewStream(c, PACKET_MAX_SIZE),
	}
}

// Read reads packet from stream
func (p *Packets) Next() (v []byte, e error) {
	pkt, err := p.stream.Read()
	if err != nil {
		return nil, err
	}

	if pkt.SequenceID != p.seq {
		return nil, sqldb.NewSQLError(sqldb.ER_MALFORMED_PACKET, "pkt.read.seq[%v]!=pkt.actual.seq[%v]", pkt.SequenceID, p.seq)
	}
	p.seq++

	return pkt.Datas, nil
}

// Write writes the packet to stream
// It packs as:
// [header]
// [payload]
func (p *Packets) Write(payload []byte) error {
	payLen := len(payload)
	pkt := common.NewBuffer(128)

	// body length(24bits)
	pkt.WriteU24(uint32(payLen))

	// SequenceID
	pkt.WriteU8(p.seq)

	// body
	pkt.WriteBytes(payload)
	if err := p.stream.Write(pkt.Datas()); err != nil {
		return err
	}
	p.seq++

	return nil
}

// WriteCommand writes a command packet to stream
func (p *Packets) WriteCommand(command byte, payload []byte) error {
	// reset packet sequence
	p.seq = 0
	pkt := common.NewBuffer(128)

	// body length(24bits):
	// command length + payload length
	payLen := len(payload)
	pkt.WriteU24(uint32(1 + payLen))

	// SequenceID
	pkt.WriteU8(p.seq)

	// command
	pkt.WriteU8(command)

	// body
	pkt.WriteBytes(payload)
	if err := p.stream.Write(pkt.Datas()); err != nil {
		return err
	}
	p.seq++

	return nil
}

// Append appends packets to buffer but not write to stream
// NOTICE: SequenceID++
func (p *Packets) Append(buff *common.Buffer, rawdata []byte) error {
	pkt := common.NewBuffer(128)

	// body length(24bits):
	// payload length
	pkt.WriteU24(uint32(len(rawdata)))

	// SequenceID
	pkt.WriteU8(p.seq)

	// body
	pkt.WriteBytes(rawdata)
	if err := p.stream.Append(buff, pkt.Datas()); err != nil {
		return err
	}
	p.seq++

	return nil
}

// AppendEOF appends EOF packet to buff
func (p *Packets) AppendEOF(buff *common.Buffer) error {
	return p.Append(buff, []byte{proto.EOF_PACKET})
}

// Flush writes all append-packets to stream
func (p *Packets) Flush(packets []byte) error {
	return p.stream.Flush(packets)
}

// ResetSeq reset sequence to zero
func (p *Packets) ResetSeq() {
	p.seq = 0
}

func (p *Packets) ParseERR(data []byte) error {
	return proto.UnPackERR(data)
}

func (p *Packets) ParseOK(data []byte) (*proto.OK, error) {
	return proto.UnPackOK(data)
}

// ReadColumns parses columns info
// http://dev.mysql.com/doc/internals/en/com-query-response.html#packet-ProtocolText::Resultset
// If myerr is not nil, we dont close the connection.
// if err is not nil, we(the client) will close the connection.
func (p *Packets) ReadColumns() ([]*querypb.Field, *proto.OK, error, error) {
	var err error
	var count uint64
	var payload []byte

	if payload, err = p.Next(); err != nil {
		return nil, nil, nil, err
	}

	ok := &proto.OK{}
	switch payload[0] {
	case proto.OK_PACKET:
		// maybe we are OK response, such as Exec response
		if ok, err = p.ParseOK(payload); err != nil {
			return nil, nil, nil, err
		}
		return nil, ok, nil, nil

	case proto.ERR_PACKET:
		return nil, nil, p.ParseERR(payload), nil
	}

	// column count
	if count, err = proto.ColumnCount(payload); err != nil {
		return nil, nil, nil, err
	}

	// column info
	columns := make([]*querypb.Field, 0, count)
	for i := 0; i < int(count); i++ {
		if payload, err = p.Next(); err != nil {
			return nil, nil, nil, err
		}

		column, err := proto.UnpackColumn(payload)
		if err != nil {
			return nil, nil, nil, err
		}
		columns = append(columns, column)
	}

	// EOF packet
	if payload, err = p.Next(); err != nil {
		return nil, nil, nil, err
	} else {
		if payload[0] != proto.EOF_PACKET {
			return nil, nil, nil, sqldb.NewSQLError(sqldb.ER_MALFORMED_PACKET, "read.columns.EOF.error.columns[%+v]", columns)
		}
	}
	return columns, ok, nil, nil
}

// WriteColumns writes columns packet to stream
func (p *Packets) WriteColumns(columns []*querypb.Field) error {
	batch := common.NewBuffer(128)

	// column count
	count := len(columns)
	buf := common.NewBuffer(64)
	buf.WriteLenEncode(uint64(count))
	if err := p.Append(batch, buf.Datas()); err != nil {
		return err
	}

	// columns info
	for i := 0; i < count; i++ {
		buf := common.NewBuffer(64)
		buf.WriteBytes(proto.PackColumn(columns[i]))
		if err := p.Append(batch, buf.Datas()); err != nil {
			return err
		}
	}
	// EOF
	if err := p.AppendEOF(batch); err != nil {
		return err
	}

	// write to stream
	if err := p.Flush(batch.Datas()); err != nil {
		return err
	}

	return nil
}

func (p *Packets) WriteOK(affectedRows, lastInsertID uint64, flags uint16, warnings uint16) error {
	ok := &proto.OK{
		AffectedRows: affectedRows,
		LastInsertID: lastInsertID,
		StatusFlags:  flags,
		Warnings:     warnings,
	}

	if err := p.Write(proto.PackOK(ok)); err != nil {
		return err
	}

	return nil
}

func (p *Packets) WriteERR(errorCode uint16, sqlState string, format string, args ...interface{}) error {
	e := &proto.ERR{
		ErrorCode:    errorCode,
		SQLState:     sqlState,
		ErrorMessage: fmt.Sprintf(format, args...),
	}
	if err := p.Write(proto.PackERR(e)); err != nil {
		return err
	}

	return nil
}
