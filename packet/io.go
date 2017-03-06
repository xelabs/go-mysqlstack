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
	"bufio"
	"io"
)

const PACKET_BUFFER_SIZE = 1500 // default MTU

type ReaderWriter struct {
	left   int
	buffer []byte
	offset int

	reader *bufio.Reader
	writer io.Writer
}

func NewReaderWriter(c io.ReadWriter) *ReaderWriter {
	return &ReaderWriter{
		writer: c,
		reader: bufio.NewReaderSize(c, 4096),
		buffer: make([]byte, PACKET_BUFFER_SIZE),
	}
}

func (rw *ReaderWriter) Read(need int) ([]byte, error) {
	extend := func(size int) {
		buffer := make([]byte, ((size/PACKET_BUFFER_SIZE)+1)*PACKET_BUFFER_SIZE)
		copy(buffer, rw.buffer[rw.offset:rw.offset+rw.left])
		rw.buffer = buffer
		rw.offset = 0
	}

	// read more datas to the buffer
	if rw.left < need {
		// extend the buffer
		if len(rw.buffer) < (rw.offset + need) {
			extend(need)
		}

		n, err := rw.reader.Read(rw.buffer[rw.offset+rw.left:])
		if err != nil {
			return nil, err
		}

		if (rw.left + n) < need {
			return nil, io.ErrUnexpectedEOF
		}

		rw.left += n
	}

	datas := rw.buffer[rw.offset : rw.offset+need]
	rw.left -= need
	rw.offset += need

	return datas, nil
}

func (rw *ReaderWriter) Write(datas []byte) error {
	var err error
	var n int

	if n, err = rw.writer.Write(datas); err != nil {
		return err
	} else {
		if n != len(datas) {
			return ErrBadConn
		}
	}

	return nil
}
