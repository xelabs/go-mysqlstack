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
	"github.com/XeLabs/go-mysqlstack/consts"
	"github.com/pkg/errors"
)

const (
	ERR_PACKET byte = 0xff
)

type ERR struct {
	Header         byte // always 0xff
	ErrorCode      uint16
	SQLStateMarker string
	SQLState       string
	ErrorMessage   string
}

func NewERR() *ERR {
	return &ERR{}
}

// https://dev.mysql.com/doc/internals/en/packet-ERR_Packet.html
func (e *ERR) UnPack(data []byte, capabilityFlags uint32) (err error) {
	buf := common.ReadBuffer(data)

	if e.Header, err = buf.ReadU8(); err != nil {
		return
	}

	if e.Header != ERR_PACKET {
		return errors.Errorf("packet.header[%v]!=ERR_PACKET", e.Header)
	}

	if e.ErrorCode, err = buf.ReadU16(); err != nil {
		return
	}

	if capabilityFlags&consts.CLIENT_PROTOCOL_41 > 0 {
		if e.SQLStateMarker, err = buf.ReadString(1); err != nil {
			return
		}
		if e.SQLState, err = buf.ReadString(5); err != nil {
			return
		}
	}

	msgLen := len(data) - buf.Seek()
	if e.ErrorMessage, err = buf.ReadString(msgLen); err != nil {
		return
	}

	return
}
