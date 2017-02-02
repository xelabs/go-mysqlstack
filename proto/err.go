/*
 * go-mysqlstack
 * xelabs.org
 *
 * Copyright (c) XeLabs
 * GPL License
 *
 */

package proto

import (
	"github.com/XeLabs/go-mysqlstack/common"
	"github.com/pkg/errors"
)

const (
	ERR_PACKET byte = 0xff
)

type ERR struct {
	Header       byte // always 0xff
	ErrorCode    uint16
	SQLState     string
	ErrorMessage string
}

// https://dev.mysql.com/doc/internals/en/packet-ERR_Packet.html
func UnPackERR(data []byte) (e *ERR, err error) {
	e = &ERR{}
	buf := common.ReadBuffer(data)

	if e.Header, err = buf.ReadU8(); err != nil {
		return
	}

	if e.Header != ERR_PACKET {
		err = errors.Errorf("packet.header[%v]!=ERR_PACKET", e.Header)
		return
	}

	if e.ErrorCode, err = buf.ReadU16(); err != nil {
		return
	}

	// Skip SQLStateMarker
	if _, err = buf.ReadString(1); err != nil {
		return
	}

	if e.SQLState, err = buf.ReadString(5); err != nil {
		return
	}

	msgLen := len(data) - buf.Seek()
	if e.ErrorMessage, err = buf.ReadString(msgLen); err != nil {
		return
	}

	return
}

func PackERR(e *ERR) []byte {
	buf := common.NewBuffer(64)

	buf.WriteU8(ERR_PACKET)

	// error code
	buf.WriteU16(e.ErrorCode)

	// sql-state marker #
	buf.WriteU8('#')

	// sql-state (?) 5 ascii bytes
	if e.SQLState == "" {
		e.SQLState = "HY000"
	}
	if len(e.SQLState) != 5 {
		panic("sqlState has to be 5 characters long")
	}
	buf.WriteString(e.SQLState)

	// error msg
	buf.WriteString(e.ErrorMessage)

	return buf.Datas()
}
