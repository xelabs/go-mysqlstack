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
	"testing"

	"github.com/XeLabs/go-mysqlstack/common"
	"github.com/XeLabs/go-mysqlstack/proto"
	"github.com/stretchr/testify/assert"
)

func TestPacketsNext(t *testing.T) {
	conn := NewMockConn()
	packets := NewPackets(conn)
	data := []byte{0x01, 0x02, 0x03}

	{
		// header
		buff := common.NewBuffer(64)
		buff.WriteU24(3)
		buff.WriteU8(0)
		buff.WriteBytes(data)

		conn.Write(buff.Datas())
		body, err := packets.Next()
		assert.Nil(t, err)
		assert.Equal(t, body, data)
	}

	{
		// header
		buff := common.NewBuffer(64)
		buff.WriteU24(3)
		buff.WriteU8(1)
		buff.WriteBytes(data)

		conn.Write(buff.Datas())
		body, err := packets.Next()
		assert.Nil(t, err)
		assert.Equal(t, body, data)
	}

	// seq error test
	{
		// header
		buff := common.NewBuffer(64)
		buff.WriteU24(3)
		buff.WriteU8(1)
		buff.WriteBytes(data)

		conn.Write(buff.Datas())
		_, err := packets.Next()
		want := "pkt.read.seq[1]!=pkt.actual.seq[2]"
		got := err.Error()
		assert.Equal(t, want, got)
	}

	// reset seq
	{
		assert.Equal(t, packets.seq, uint8(2))
		packets.ResetSeq()
		assert.Equal(t, packets.seq, uint8(0))
	}
}

func TestPacketsNextFail(t *testing.T) {
	conn := NewMockConn()
	packets := NewPackets(conn)
	data1 := []byte{0x00, 0x00, 0x00}
	data2 := []byte{0x00, 0x00, 0x00, 0x00}
	data3 := []byte{0x01, 0x10, 0x00, 0x00}

	{
		conn.Write(data1)
		_, err := packets.Next()
		assert.NotNil(t, err)
	}

	{
		conn.Write(data2)
		_, err := packets.Next()
		assert.Nil(t, err)
	}

	{
		conn.Write(data3)
		_, err := packets.Next()
		assert.NotNil(t, err)
	}
}

func TestPacketsWrite(t *testing.T) {
	conn := NewMockConn()
	buff := common.NewBuffer(64)
	packets := NewPackets(conn)
	data := []byte{0x01, 0x02, 0x03}

	{
		buff.WriteU24(3)
		buff.WriteU8(0)
		buff.WriteBytes(data)
		want := buff.Datas()

		err := packets.Write(data)
		assert.Nil(t, err)
		got := conn.Datas()
		assert.Equal(t, want, got)
	}

	{
		buff.WriteU24(3)
		buff.WriteU8(1)
		buff.WriteBytes(data)
		want := buff.Datas()

		err := packets.Write(data)
		assert.Nil(t, err)
		got := conn.Datas()
		assert.Equal(t, want, got)
	}
}

func TestPacketsBatchWrite(t *testing.T) {
	conn := NewMockConn()
	buff := common.NewBuffer(64)
	batch := common.NewBuffer(64)
	packets := NewPackets(conn)
	bodys := [][]byte{
		[]byte{0x01, 0x11, 0x12},
		[]byte{0x02, 0x21, 0x22},
		[]byte{0x03, 0x31, 0x32},
		[]byte{0x04, 0x41, 0x42},
	}

	// batch write
	{
		for _, body := range bodys {
			buff.WriteBytes(body)
			err := packets.Append(batch, body)
			assert.Nil(t, err)
		}

		// commit
		{
			err := packets.Flush(batch.Datas())
			assert.Nil(t, err)
			want := batch.Datas()
			got := conn.Datas()
			assert.Equal(t, want, got)
		}
	}

	// batch read
	{
		read := NewPackets(conn)

		for _, want := range bodys {
			got, err := read.Next()
			assert.Nil(t, err)
			assert.Equal(t, want, got)
		}
	}
}

func TestPacketsWriteCommand(t *testing.T) {
	buff := common.NewBuffer(64)
	conn := NewMockConn()
	packets := NewPackets(conn)
	cmd := 0x03
	data := []byte{0x01, 0x02, 0x03}

	{
		buff.WriteU24(3 + 1)
		buff.WriteU8(0)
		buff.WriteU8(uint8(cmd))
		buff.WriteBytes(data)
		want := buff.Datas()

		err := packets.WriteCommand(byte(cmd), data)
		assert.Nil(t, err)
		got := conn.Datas()
		assert.Equal(t, want, got)
	}
}

func TestPacketsParseERR(t *testing.T) {
	buff := common.NewBuffer(32)
	conn := NewMockConn()
	packets := NewPackets(conn)

	// header
	buff.WriteU8(0xff)
	// error_code
	buff.WriteU16(0x01)
	// sql_state_marker
	buff.WriteString("a")
	// sql_state
	buff.WriteString("ABCDE")
	buff.WriteString("ERROR")

	want := proto.NewERR()
	want.Header = 0xff
	want.ErrorCode = 0x1
	want.SQLStateMarker = "a"
	want.SQLState = "ABCDE"
	want.ErrorMessage = "ERROR"

	got, err := packets.ParseERR(buff.Datas(), proto.DefaultCapability)
	assert.Nil(t, err)
	assert.Equal(t, want, got)
}

func TestPacketsParseOK(t *testing.T) {
	buff := common.NewBuffer(32)
	conn := NewMockConn()
	packets := NewPackets(conn)

	// header
	buff.WriteU8(0x00)
	// affected_rows
	buff.WriteLenEncode(uint64(3))
	// last_insert_id
	buff.WriteLenEncode(uint64(40000000000))

	// status_flags
	buff.WriteU16(0x01)
	// warnings
	buff.WriteU16(0x02)

	want := proto.NewOK()
	want.AffectedRows = 3
	want.LastInsertID = 40000000000
	want.StatusFlags = 1
	want.Warnings = 2

	got, err := packets.ParseOK(buff.Datas(), proto.DefaultCapability)
	assert.Nil(t, err)
	assert.Equal(t, want, got)
}

func TestPacketsColumns(t *testing.T) {
	conn := NewMockConn()
	wPackets := NewPackets(conn)
	rPackets := NewPackets(conn)
	columns := []*proto.Column{
		&proto.Column{
			Catalog:    "def",
			Schema:     "test",
			Table:      "t1",
			Org_Table:  "t1",
			Name:       "a",
			Org_Name:   "a",
			Charset:    11,
			ColumnLen:  11,
			FieldType:  11,
			FieldFlags: 11,
		},
		&proto.Column{
			Catalog:    "def",
			Schema:     "test",
			Table:      "t1",
			Org_Table:  "t1",
			Name:       "b",
			Org_Name:   "b",
			Charset:    12,
			ColumnLen:  12,
			FieldType:  12,
			FieldFlags: 12,
		},
	}

	{
		err := wPackets.WriteColumns(columns)
		assert.Nil(t, err)
	}

	{
		cols, _, err := rPackets.ReadColumns(proto.DefaultCapability)
		assert.Nil(t, err)
		assert.Equal(t, columns, cols)
	}
}

func TestPacketsColumnsOK(t *testing.T) {
	conn := NewMockConn()
	wPackets := NewPackets(conn)
	rPackets := NewPackets(conn)
	{
		buff := common.NewBuffer(32)

		// header
		buff.WriteU8(0x00)
		// affected_rows
		buff.WriteLenEncode(uint64(3))
		// last_insert_id
		buff.WriteLenEncode(uint64(40000000000))

		// status_flags
		buff.WriteU16(0x01)
		// warnings
		buff.WriteU16(0x02)
		wPackets.Write(buff.Datas())
	}

	{
		want := proto.NewOK()
		want.AffectedRows = 3
		want.LastInsertID = 40000000000
		want.StatusFlags = 1
		want.Warnings = 2

		_, got, err := rPackets.ReadColumns(proto.DefaultCapability)
		assert.Nil(t, err)
		assert.Equal(t, want, got)
	}
}

func TestPacketsColumnsERR(t *testing.T) {
	conn := NewMockConn()
	wPackets := NewPackets(conn)
	rPackets := NewPackets(conn)
	{
		buff := common.NewBuffer(32)

		// header
		buff.WriteU8(0xff)
		// error_code
		buff.WriteU16(0x01)
		// sql_state_marker
		buff.WriteString("a")
		// sql_state
		buff.WriteString("ABCDE")
		buff.WriteString("ERROR")
		wPackets.Write(buff.Datas())
	}

	{
		want := "ERROR"
		_, _, err := rPackets.ReadColumns(proto.DefaultCapability)
		got := err.Error()
		assert.Equal(t, want, got)
	}
}

func TestPacketsColumnsError(t *testing.T) {
	conn := NewMockConn()
	wPackets := NewPackets(conn)
	rPackets := NewPackets(conn)
	{
		buff := common.NewBuffer(32)

		// random datas
		buff.WriteU8(0xf0)
		buff.WriteU16(0x11)
		wPackets.Write(buff.Datas())
	}

	{
		want := "EOF"
		_, _, err := rPackets.ReadColumns(proto.DefaultCapability)
		got := err.Error()
		assert.Equal(t, want, got)
	}
}
