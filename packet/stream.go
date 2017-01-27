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
	"io"

	"github.com/XeLabs/go-mysqlstack/common"
)

type Stream struct {
	pktMaxSize int
	rw         *ReaderWriter
}

func NewStream(c io.ReadWriter, pktMaxSize int) *Stream {
	return &Stream{
		pktMaxSize: pktMaxSize,
		rw:         NewReaderWriter(c),
	}
}

// Read reads the next packet from the reader
// The returned pkt.Payload is only guaranteed to be valid until the next read
func (s *Stream) Read() (pkt Packet, err error) {
	var datas []byte

	for {
		var header []byte
		var body []byte

		// read packet header [32 bit]
		if header, err = s.rw.Read(4); err != nil {
			return
		}

		// payload length [24 bit]
		payLen := int(uint32(header[0]) | uint32(header[1])<<8 | uint32(header[2])<<16)
		pkt.SequenceID = header[3]

		if body, err = s.rw.Read(payLen); err != nil {
			return
		}

		// more chunks
		if datas != nil {
			datas = append(datas, body...)
		} else {
			datas = body
		}

		if payLen < s.pktMaxSize {
			break
		}
	}
	pkt.Payload = datas

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

func (s *Stream) Flush(datas []byte) error {
	return s.rw.Write(datas)
}
