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
	"github.com/XeLabs/go-mysqlstack/common"
	"github.com/XeLabs/go-mysqlstack/proto"
	"github.com/stretchr/testify/assert"
	"testing"
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
		buff.WriteBytes(data, len(data))

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
		buff.WriteBytes(data, len(data))

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
		buff.WriteBytes(data, len(data))

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
		buff.WriteBytes(data, len(data))
		want := buff.Datas()

		err := packets.Write(data)
		assert.Nil(t, err)
		got := conn.Datas()
		assert.Equal(t, want, got)
	}

	{
		buff.WriteU24(3)
		buff.WriteU8(1)
		buff.WriteBytes(data, len(data))
		want := buff.Datas()

		err := packets.Write(data)
		assert.Nil(t, err)
		got := conn.Datas()
		assert.Equal(t, want, got)
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
		buff.WriteBytes(data, len(data))
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
	buff.WriteString("a", 1)
	// sql_state
	buff.WriteString("ABCDE", 5)
	buff.WriteString("ERROR", 5)

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
