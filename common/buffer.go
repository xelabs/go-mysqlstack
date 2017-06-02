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
	"bytes"
	"io"
)

var (
	ErrIOEOF = io.EOF
)

type Buffer struct {
	pos  int
	seek int
	cap  int
	buf  []byte
}

func NewBuffer(cap int) *Buffer {
	return &Buffer{pos: 0,
		cap: cap,
		buf: make([]byte, cap),
	}
}

func ReadBuffer(b []byte) *Buffer {
	return &Buffer{
		buf: b,
		pos: len(b),
	}
}

func (b *Buffer) Reset(data []byte) {
	b.buf = data
	b.pos = len(data)
	b.seek = 0
}

func (b *Buffer) Datas() []byte {
	return b.buf[:b.pos]
}

func (b *Buffer) Length() int {
	return b.pos
}

func (b *Buffer) Seek() int {
	return b.seek
}

func (b *Buffer) extend(n int) {
	if (b.pos + n) > b.cap {
		// allocate double what's needed, for future growth
		b.cap = (b.pos + n) * 2
		t := make([]byte, b.cap)
		copy(t, b.buf)
		b.buf = t
	}
}

func (b *Buffer) WriteU8(v uint8) {
	b.extend(1)
	b.buf[b.pos] = v
	b.pos++
}

func (b *Buffer) ReadU8() (v uint8, err error) {
	if (b.seek + 1) > b.pos {
		err = ErrIOEOF
		return
	}

	v = uint8(b.buf[b.seek])
	b.seek++
	return
}

func (b *Buffer) WriteU16(v uint16) {
	b.extend(2)
	b.buf[b.pos] = byte(v)
	b.buf[b.pos+1] = byte(v >> 8)
	b.pos += 2
}

func (b *Buffer) ReadU16() (v uint16, err error) {
	if (b.seek + 2) > b.pos {
		err = ErrIOEOF
		return
	}

	v = uint16(b.buf[b.seek]) |
		uint16(b.buf[b.seek+1])<<8
	b.seek += 2
	return
}

func (b *Buffer) WriteU24(v uint32) {
	b.extend(3)
	b.buf[b.pos] = byte(v)
	b.buf[b.pos+1] = byte(v >> 8)
	b.buf[b.pos+2] = byte(v >> 16)
	b.pos += 3
}

func (b *Buffer) ReadU24() (v uint32, err error) {
	if (b.seek + 3) > b.pos {
		err = ErrIOEOF
		return
	}

	v = uint32(b.buf[b.seek]) |
		uint32(b.buf[b.seek+1])<<8 |
		uint32(b.buf[b.seek+2])<<16
	b.seek += 3
	return
}

func (b *Buffer) WriteU32(v uint32) {
	b.extend(4)
	b.buf[b.pos] = byte(v)
	b.buf[b.pos+1] = byte(v >> 8)
	b.buf[b.pos+2] = byte(v >> 16)
	b.buf[b.pos+3] = byte(v >> 24)
	b.pos += 4
}

func (b *Buffer) ReadU32() (v uint32, err error) {
	if (b.seek + 4) > b.pos {
		err = ErrIOEOF
		return
	}

	v = uint32(b.buf[b.seek]) |
		uint32(b.buf[b.seek+1])<<8 |
		uint32(b.buf[b.seek+2])<<16 |
		uint32(b.buf[b.seek+3])<<24
	b.seek += 4
	return
}

func (b *Buffer) WriteU64(v uint64) {
	b.extend(8)
	b.buf[b.pos] = byte(v)
	b.buf[b.pos+1] = byte(v >> 8)
	b.buf[b.pos+2] = byte(v >> 16)
	b.buf[b.pos+3] = byte(v >> 24)
	b.buf[b.pos+4] = byte(v >> 32)
	b.buf[b.pos+5] = byte(v >> 40)
	b.buf[b.pos+6] = byte(v >> 48)
	b.buf[b.pos+7] = byte(v >> 56)
	b.pos += 8
}

func (b *Buffer) ReadU64() (v uint64, err error) {
	if (b.seek + 8) > b.pos {
		err = ErrIOEOF
		return
	}

	v = uint64(b.buf[b.seek]) |
		uint64(b.buf[b.seek+1])<<8 |
		uint64(b.buf[b.seek+2])<<16 |
		uint64(b.buf[b.seek+3])<<24 |
		uint64(b.buf[b.seek+4])<<32 |
		uint64(b.buf[b.seek+5])<<40 |
		uint64(b.buf[b.seek+6])<<48 |
		uint64(b.buf[b.seek+7])<<56
	b.seek += 8
	return
}

// https://dev.mysql.com/doc/internals/en/integer.html#length-encoded-integer
func (b *Buffer) WriteLenEncode(v uint64) {
	switch {
	case v < 251:
		b.WriteU8(uint8(v))

	case v >= 251 && v < (1<<16):
		b.WriteU8(0xfc)
		b.WriteU16(uint16(v))

	case v >= (1<<16) && v < (1<<24):
		b.WriteU8(0xfd)
		b.WriteU24(uint32(v))

	default:
		b.WriteU8(0xfe)
		b.WriteU64(v)
	}
}

// 0xfb is represents a NULL in a ProtocolText::ResultsetRow
func (b *Buffer) WriteLenEncodeNUL() {
	b.WriteU8(0xfb)
}

func (b *Buffer) ReadLenEncode() (v uint64, err error) {
	var u8 uint8
	var u16 uint16
	var u24 uint32

	if u8, err = b.ReadU8(); err != nil {
		return
	}

	switch u8 {
	case 0xfb:
		v = 0
		return

	case 0xfc:
		if u16, err = b.ReadU16(); err != nil {
			return
		}
		v = uint64(u16)
		return

	case 0xfd:
		if u24, err = b.ReadU24(); err != nil {
			return
		}
		v = uint64(u24)
		return

	case 0xfe:
		if v, err = b.ReadU64(); err != nil {
			return
		}
		return

	default:
		return uint64(u8), nil
	}
}

func (b *Buffer) WriteLenEncodeString(s string) {
	l := len(s)
	b.WriteLenEncode(uint64(l))
	b.WriteString(s)
}

func (b *Buffer) ReadLenEncodeString() (s string, err error) {
	var l uint64

	if l, err = b.ReadLenEncode(); err != nil {
		return
	}

	if s, err = b.ReadString(int(l)); err != nil {
		return
	}
	return
}

func (b *Buffer) WriteLenEncodeBytes(v []byte) {
	l := len(v)
	b.WriteLenEncode(uint64(l))
	b.WriteBytes(v)
}

func (b *Buffer) ReadLenEncodeBytes() (v []byte, err error) {
	var l uint64

	if l, err = b.ReadLenEncode(); err != nil {
		return
	}

	if l > 0 {
		if v, err = b.ReadBytes(int(l)); err != nil {
			return
		}
	} else {
		v = nil
	}
	return
}

func (b *Buffer) WriteEOF(n int) {
	b.extend(n)
	for i := 0; i < n; i++ {
		b.buf[b.pos] = 0xfe
		b.pos++
	}
}

func (b *Buffer) ReadEOF(n int) (err error) {
	return b.ReadZero(n)
}

func (b *Buffer) WriteZero(n int) {
	b.extend(n)
	for i := 0; i < n; i++ {
		b.buf[b.pos] = 0
		b.pos++
	}
}

func (b *Buffer) ReadZero(n int) (err error) {
	if (b.seek + n) > b.pos {
		err = ErrIOEOF
		return
	}
	b.seek += n
	return
}

func (b *Buffer) WriteString(s string) {
	n := len(s)
	b.extend(n)
	copy(b.buf[b.pos:], s)
	b.pos += n
}

func (b *Buffer) ReadString(n int) (s string, err error) {
	if (b.seek + n) > b.pos {
		err = ErrIOEOF
		return
	}

	s = string(b.buf[b.seek:(b.seek + n)])
	b.seek += n
	return
}

// ReadStringNUL reads until the first NUL in the buffer
// returning a string containing the data up to and not including the NUL
func (b *Buffer) ReadStringNUL() (s string, err error) {
	var v []byte

	if v, err = b.readBytesWithToken(0x00); err != nil {
		return
	}
	s = string(v)
	return
}

// ReadStringEOF reads until the first EOF in the buffer
// returning a string containing the data up to and not including the EOF
func (b *Buffer) ReadStringEOF() (s string, err error) {
	var v []byte

	if v, err = b.readBytesWithToken(0xfe); err != nil {
		return
	}
	s = string(v)
	return
}

// ReadBytesNUL reads until the first NUL in the buffer
// returning a byte slice containing the data up to and not including the NUL
func (b *Buffer) ReadBytesNUL() (v []byte, err error) {
	return b.readBytesWithToken(0x00)
}

// ReadBytesEOF reads until the first EOF in the buffer
// returning a byte slice containing the data up to and not including the EOF
func (b *Buffer) ReadBytesEOF() (v []byte, err error) {
	return b.readBytesWithToken(0xfe)
}

func (b *Buffer) readBytesWithToken(token uint8) (v []byte, err error) {
	i := bytes.IndexByte(b.buf[b.seek:], token)
	end := b.seek + i + 1
	if i < 0 {
		b.seek = len(b.buf)
		err = ErrIOEOF
		return
	}
	v = b.buf[b.seek : end-1]
	b.seek = end
	return
}

func (b *Buffer) WriteBytes(bs []byte) {
	n := len(bs)
	b.extend(n)
	copy(b.buf[b.pos:], bs)
	b.pos += n
}

func (b *Buffer) ReadBytes(n int) (v []byte, err error) {
	if n == 0 {
		return nil, nil
	}

	if (b.seek + n) > b.pos {
		err = ErrIOEOF
		return
	}

	v = b.buf[b.seek:(b.seek + n)]
	b.seek += n
	return
}
