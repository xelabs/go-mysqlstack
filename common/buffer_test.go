/*
 * go-mysqlstack
 * xelabs.org
 *
 * Copyright (c) XeLabs
 * GPL License
 *
 */

package common

import (
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
)

func TestBuffer(t *testing.T) {
	buf := NewBuffer(6)
	buf.WriteU32(22222232)
	buf.WriteU32(31)
	buf.WriteU32(30)
	buf.WriteU8(208)
	buf.WriteU16(65535)
	buf.WriteBytes([]byte{1, 2, 3, 4, 5}, 5)
	buf.WriteZero(3)
	buf.WriteString("abc", 3)
	buf.WriteEOF(1)
	buf.WriteString("xyz", 3)
	buf.WriteEOF(2)
	buf.WriteU24(1024)

	{
		want := (4 + 4 + 4 + 1 + 2 + 5 + 3 + 3 + 1 + 3 + 2 + 3)
		got := buf.Length()
		assert.Equal(t, want, got)
	}

	{
		want := uint32(22222232)
		got, _ := buf.ReadU32()
		assert.Equal(t, want, got)
	}

	{
		want := uint32(31)
		got, _ := buf.ReadU32()
		assert.Equal(t, want, got)
	}

	{
		want := uint32(30)
		got, _ := buf.ReadU32()
		assert.Equal(t, want, got)
	}

	{
		want := uint8(208)
		got, _ := buf.ReadU8()
		assert.Equal(t, want, got)
	}

	{
		want := uint16(65535)
		got, _ := buf.ReadU16()
		assert.Equal(t, want, got)
	}

	{
		want := []byte{1, 2, 3, 4, 5}
		got, _ := buf.ReadBytes(5)
		assert.Equal(t, want, got)
	}

	{
		buf.ReadZero(3)
	}

	{
		want := "abc"
		got, _ := buf.ReadString(3)
		assert.Equal(t, want, got)
	}

	{
		buf.ReadEOF(1)
	}

	{
		want := "xyz"
		got, _ := buf.ReadStringEOF()
		assert.Equal(t, want, got)
	}

	{
		buf.ReadEOF(1)
	}

	{
		want := uint32(1024)
		got, _ := buf.ReadU24()
		assert.Equal(t, want, got)
	}

	{
		want := buf.Length()
		got := buf.Seek()
		assert.Equal(t, want, got)
	}

}

func TestBufferDatas(t *testing.T) {
	buf := NewBuffer(100)
	buf.WriteU32(22222232)
	buf.WriteString("abc", 3)
	buf.WriteZero(2)

	{
		want := len(buf.Datas())
		got := buf.Length()
		assert.Equal(t, want, got)
	}

	{
		want := []byte{152, 21, 83, 1, 97, 98, 99, 0, 0}
		got := buf.Datas()
		assert.Equal(t, want, got)
	}
}

func TestBufferRead(t *testing.T) {
	data := []byte{152, 21, 83, 1, 97, 98, 99, 0, 0}
	buf := ReadBuffer(data)
	{
		want := uint32(22222232)
		got, _ := buf.ReadU32()
		assert.Equal(t, want, got)
	}

	{
		want := "abc"
		got, _ := buf.ReadString(3)
		assert.Equal(t, want, got)
	}
}

func TestBufferReadError(t *testing.T) {
	{
		data := []byte{152}
		buf := ReadBuffer(data)
		_, err := buf.ReadU8()
		assert.Nil(t, err)
	}

	{
		data := []byte{152}
		buf := ReadBuffer(data)
		want := io.EOF
		_, got := buf.ReadU16()
		assert.Equal(t, want.Error(), got.Error())
	}

	{
		data := []byte{152, 154}
		buf := ReadBuffer(data)
		want := io.EOF
		_, got := buf.ReadU24()
		assert.Equal(t, want.Error(), got.Error())
	}

	{
		data := []byte{152, 154, 155}
		buf := ReadBuffer(data)
		want := io.EOF
		_, got := buf.ReadU32()
		assert.Equal(t, want.Error(), got.Error())
	}

	{
		data := []byte{152, 154, 155}
		buf := ReadBuffer(data)
		want := io.EOF
		got := buf.ReadZero(4)
		assert.Equal(t, want.Error(), got.Error())
	}

	{
		data := []byte{152, 154, 155}
		buf := ReadBuffer(data)
		want := io.EOF
		_, got := buf.ReadString(4)
		assert.Equal(t, want.Error(), got.Error())
	}

	{
		data := []byte{152, 154, 155}
		buf := ReadBuffer(data)
		want := io.EOF
		_, got := buf.ReadStringNUL()
		assert.Equal(t, want.Error(), got.Error())
	}

	{
		data := []byte{152, 154, 155}
		buf := ReadBuffer(data)
		want := io.EOF
		_, got := buf.ReadBytes(4)
		assert.Equal(t, want.Error(), got.Error())
	}
}

func TestBufferReadString(t *testing.T) {
	data := []byte{
		0x98, 0x15, 0x53, 0x01, 0x61, 0x62, 0x63, 0xff,
		0xff, 0x61, 0x62, 0x63, 0x00, 0x00, 0xff, 0xff}
	buf := ReadBuffer(data)

	{
		want := 0
		got := buf.seek
		assert.Equal(t, want, got)
	}

	{
		want := 16
		got := buf.pos
		assert.Equal(t, want, got)
	}

	{
		want := 16
		got := buf.pos
		assert.Equal(t, want, got)
	}

	{
		want := uint32(22222232)
		got, _ := buf.ReadU32()
		assert.Equal(t, want, got)
	}

	{
		want := "abc"
		got, _ := buf.ReadString(3)
		assert.Equal(t, want, got)
	}

	{
		want := uint16(65535)
		got, _ := buf.ReadU16()
		assert.Equal(t, want, got)
	}

	{
		want := "abc"
		got, _ := buf.ReadStringNUL()
		assert.Equal(t, want, got)
	}

	{
		want := 13
		got := buf.seek
		assert.Equal(t, want, got)
		buf.ReadZero(1)
	}

	// here, we inject a ReadStringWithNUL
	// we never got it since here is ReadU16()
	{

		want := "EOF"
		_, err := buf.ReadStringNUL()
		got := err.Error()
		assert.Equal(t, want, got)
	}

	{
		want := 16
		got := buf.seek
		assert.Equal(t, want, got)
	}
}

func TestBufferLenEncode(t *testing.T) {
	buf := NewBuffer(6)

	{
		v := uint64(250)
		buf.WriteLenEncode(v)
	}

	{
		v := uint64(252)
		buf.WriteLenEncode(v)
	}

	{
		v := uint64(1 << 16)
		buf.WriteLenEncode(v)
	}

	{
		buf.WriteLenEncodeNUL()
	}

	{
		v := uint64(1 << 24)
		buf.WriteLenEncode(v)
	}

	{
		v := uint64(1<<24 + 1)
		buf.WriteLenEncode(v)
	}

	read := ReadBuffer(buf.Datas())

	{
		v, err := read.ReadLenEncode()
		assert.Nil(t, err)
		assert.Equal(t, v, uint64(250))
	}

	{
		v, err := read.ReadLenEncode()
		assert.Nil(t, err)
		assert.Equal(t, v, uint64(252))
	}

	{
		v, err := read.ReadLenEncode()
		assert.Nil(t, err)
		assert.Equal(t, v, uint64(1<<16))
	}

	{
		v, err := read.ReadLenEncode()
		assert.Nil(t, err)
		assert.Equal(t, v, uint64(0))
	}

	{
		v, err := read.ReadLenEncode()
		assert.Nil(t, err)
		assert.Equal(t, v, uint64(1<<24))
	}

	{
		v, err := read.ReadLenEncode()
		assert.Nil(t, err)
		assert.Equal(t, v, uint64(1<<24+1))
	}
}

func TestBufferLenEncodeString(t *testing.T) {
	buf := NewBuffer(6)

	s := "BohuTANG"
	{
		v := uint64(len(s))
		buf.WriteLenEncode(v)
	}

	{
		buf.WriteString(s, len(s))
	}

	{
		buf.WriteEOF(3)
	}

	data := []byte{
		0x98, 0x15, 0x53, 0x01, 0x61, 0x62, 0x63, 0xff,
		0xff, 0x61, 0x62, 0x63, 0x00, 0x00, 0xff, 0xff}
	{
		v := uint64(len(data))
		buf.WriteLenEncode(v)
	}

	{
		buf.WriteBytes(data, len(data))
	}

	reader := ReadBuffer(buf.Datas())
	{
		got, err := reader.ReadLenEncodeString()
		assert.Nil(t, err)
		assert.Equal(t, s, got)
	}

	{
		reader.ReadZero(3)
	}

	{
		got, err := reader.ReadLenEncodeBytes()
		assert.Nil(t, err)
		assert.Equal(t, data, got)
	}
}

func TestBufferNULEOF(t *testing.T) {
	buf := NewBuffer(16)
	data1 := "BohuTANG"
	data2 := "radon"

	{
		buf.WriteString(data1, len(data1))
		buf.WriteZero(1)
	}

	{
		buf.WriteString(data2, len(data2))
		buf.WriteZero(1)
	}

	{
		buf.WriteString(data1, len(data1))
		buf.WriteEOF(1)
	}

	{
		buf.WriteString(data2, len(data2))
		buf.WriteEOF(1)
	}

	reader := ReadBuffer(buf.Datas())
	{
		got, _ := reader.ReadStringNUL()
		assert.Equal(t, data1, got)
	}

	{
		got, _ := reader.ReadBytesNUL()
		assert.Equal(t, StringToBytes(data2), got)
	}

	{
		got, _ := reader.ReadStringEOF()
		assert.Equal(t, data1, got)
	}

	{
		got, _ := reader.ReadBytesEOF()
		assert.Equal(t, StringToBytes(data2), got)
	}
}

func TestBufferReset(t *testing.T) {
	buf := NewBuffer(6)
	buf.WriteU32(31)
	buf.WriteU32(30)

	{
		want := uint32(31)
		got, _ := buf.ReadU32()
		assert.Equal(t, want, got)
		assert.Equal(t, buf.seek, 4)
	}

	{
		data := []byte{0x00, 0x00, 0x00, 0x01}
		buf.Reset(data)
		assert.Equal(t, buf.pos, 4)
		assert.Equal(t, buf.seek, 0)
	}
}
