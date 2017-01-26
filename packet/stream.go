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
	"io"

	"github.com/XeLabs/go-mysqlstack/common"
	"github.com/pkg/errors"
)

const PACKET_BUFFER_SIZE = 1500 // default MTU

type Stream struct {
	writer     io.Writer
	reader     *bufio.Reader
	buffer     []byte
	offset     int
	header     []byte
	pktMaxSize int
}

func NewStream(c io.ReadWriter, pktMaxSize int) *Stream {
	return &Stream{
		writer:     c,
		reader:     bufio.NewReaderSize(c, 4096),
		buffer:     make([]byte, PACKET_BUFFER_SIZE),
		header:     make([]byte, 4),
		pktMaxSize: pktMaxSize,
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

	// reset buffer when it reaches pktMaxSize
	if len(s.buffer) > s.pktMaxSize {
		s.buffer = make([]byte, s.pktMaxSize)
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

		if payLen < s.pktMaxSize {
			break
		}
	}
	pkt.Payload = s.buffer[:s.offset]

	return
}

// Write writes the packet to writer
func (s *Stream) Write(data []byte) error {
	buf := common.NewBuffer(64)
	if err := s.Append(buf, data); err != nil {
		return err
	}

	if err := s.Flush(buf.Datas()); err != nil {
		return err
	}

	return nil
}

func (s *Stream) Append(buff *common.Buffer, data []byte) error {
	payLen := len(data) - 4
	sequence := data[3]

	for {
		var size int
		if payLen < s.pktMaxSize {
			size = payLen
		} else {
			size = s.pktMaxSize
		}
		data[0] = byte(size)
		data[1] = byte(size >> 8)
		data[2] = byte(size >> 16)
		data[3] = sequence

		// append to buffer
		buff.WriteBytes(data[:4+size])
		if size < s.pktMaxSize {
			break
		}

		payLen -= size
		data = data[size:]
		sequence++
	}

	return nil
}

func (s *Stream) Flush(data []byte) error {
	if n, err := s.writer.Write(data); err != nil {
		return errors.WithStack(err)
	} else {
		if n != len(data) {
			return ErrBadConn
		}
	}

	return nil
}
