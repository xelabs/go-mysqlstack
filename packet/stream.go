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
	"net"
)

const (
	// PACKET_BUFFER_SIZE is how much we buffer for reading.
	PACKET_BUFFER_SIZE = 32 * 1024
)

type Stream struct {
	pktMaxSize int
	header     []byte
	reader     *bufio.Reader
	writer     *bufio.Writer
}

func NewStream(conn net.Conn, pktMaxSize int) *Stream {
	return &Stream{
		pktMaxSize: pktMaxSize,
		header:     []byte{0, 0, 0, 0},
		reader:     bufio.NewReaderSize(conn, PACKET_BUFFER_SIZE),
		writer:     bufio.NewWriterSize(conn, PACKET_BUFFER_SIZE),
	}
}

// Read reads the next packet from the reader
// The returned pkt.Datas is only guaranteed to be valid until the next read
func (s *Stream) Read() (*Packet, error) {
	// Header.
	if _, err := io.ReadFull(s.reader, s.header); err != nil {
		return nil, err
	}

	// Length.
	pkt := &Packet{}
	pkt.SequenceID = s.header[3]
	length := int(uint32(s.header[0]) | uint32(s.header[1])<<8 | uint32(s.header[2])<<16)
	if length == 0 {
		return pkt, nil
	}

	// Datas.
	data := make([]byte, length)
	if _, err := io.ReadFull(s.reader, data); err != nil {
		return nil, err
	}
	pkt.Datas = data

	// Single packet.
	if length < s.pktMaxSize {
		return pkt, nil
	}

	// There is more than one packet, read them all.
	next, err := s.Read()
	if err != nil {
		return nil, err
	} else {
		pkt.SequenceID = next.SequenceID
		pkt.Datas = append(pkt.Datas, next.Datas...)
		return pkt, nil
	}
}

// Write writes the packet to writer
func (s *Stream) Write(data []byte) error {
	if err := s.Append(data); err != nil {
		return err
	}

	if err := s.Flush(); err != nil {
		return err
	}
	return nil
}

func (s *Stream) Append(data []byte) error {
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
		s.writer.Write(data[:4+size])
		if size < s.pktMaxSize {
			break
		}

		payLen -= size
		data = data[size:]
		sequence++
	}
	return nil
}

func (s *Stream) Flush() error {
	return s.writer.Flush()
}
