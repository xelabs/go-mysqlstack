/*
 * go-mysqlstack
 * xelabs.org
 *
 * Copyright (c) XeLabs
 * GPL License
 *
 */

// stream.go
// handles the raw mysql packets without bytes copy
// More than 16MByte packet:
// https://dev.mysql.com/doc/internals/en/sending-more-than-16mbyte.html
package packet

import (
	"bufio"
	"github.com/pkg/errors"
	"io"
)

const PACKET_MAX_SIZE = (1<<24 - 1) // (16MB - 1）
const PACKET_BUFFER_SIZE = 1500     // default MTU

type Stream struct {
	writer io.Writer
	reader *bufio.Reader
	buffer []byte
	offset int
	header []byte
}

func NewStream(c io.ReadWriter) *Stream {
	return &Stream{
		writer: c,
		reader: bufio.NewReaderSize(c, 4096),
		buffer: make([]byte, PACKET_BUFFER_SIZE),
		header: make([]byte, 4),
	}
}

// Read reads the next packet from the reader
// The returned pkt.Payload is only guaranteed to be valid until the next read
func (s *Stream) Read() (pkt Packet, err error) {
	extend := func(size int) {
		buf := make([]byte, size)
		copy(buf, s.buffer[:s.offset])
		s.buffer = buf
	}
	s.offset = 0

	// reset buffer when it reaches PACKET_RESET_SIZE
	if len(s.buffer) > PACKET_MAX_SIZE {
		s.buffer = make([]byte, PACKET_BUFFER_SIZE)
	}

	for {
		// read packet header [32 bit]
		if _, err = io.ReadFull(s.reader, s.header); err != nil {
			err = errors.WithStack(err)
			return
		}

		// payload length [24 bit]
		payLen := int(uint32(s.header[0]) | uint32(s.header[1])<<8 | uint32(s.header[2])<<16)

		// resize the buffer
		total := s.offset + payLen
		if total > len(s.buffer) {
			extend(total)
		}

		// read payload [payLen bytes]
		buf := s.buffer[s.offset : s.offset+payLen]
		if _, err = io.ReadFull(s.reader, buf); err != nil {
			err = errors.WithStack(err)
			return
		}

		s.offset += payLen
		pkt.SequenceID = s.header[3]

		if payLen < PACKET_MAX_SIZE {
			break
		}
	}
	pkt.Payload = s.buffer[:s.offset]

	return
}

// Write writes the packet to writer
func (s *Stream) Write(data []byte) error {
	payLen := len(data) - 4
	sequence := data[3]

	for {
		var size int
		if payLen < PACKET_MAX_SIZE {
			data[0] = byte(payLen)
			data[1] = byte(payLen >> 8)
			data[2] = byte(payLen >> 16)
			data[3] = sequence
			size = payLen
		} else {
			data[0] = 0xff
			data[1] = 0xff
			data[2] = 0xff
			data[3] = sequence
			size = PACKET_MAX_SIZE
		}

		if n, err := s.writer.Write(data[:4+size]); err != nil {
			return errors.WithStack(err)
		} else {
			if n != (4 + size) {
				return ErrBadConn
			}
		}

		if size < PACKET_MAX_SIZE {
			break
		}

		payLen -= size
		data = data[size:]
		sequence++
	}

	return nil
}
