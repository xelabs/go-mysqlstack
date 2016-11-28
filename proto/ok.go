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
	OK_PACKET byte = 0x00
)

type OK struct {
	Header       byte // 0x00
	AffectedRows uint64
	LastInsertID uint64
	StatusFlags  uint16
	Warnings     uint16
}

func NewOK() *OK {
	return &OK{}
}

// https://dev.mysql.com/doc/internals/en/packet-OK_Packet.html
func (o *OK) UnPack(data []byte, capabilityFlags uint32) (err error) {
	buf := common.ReadBuffer(data)

	if o.Header, err = buf.ReadU8(); err != nil {
		return
	}

	if o.Header != OK_PACKET {
		return errors.Errorf("packet.header[%v]!=OK_PACKET", o.Header)
	}

	if o.AffectedRows, err = buf.ReadLenEncode(); err != nil {
		return
	}

	if o.LastInsertID, err = buf.ReadLenEncode(); err != nil {
		return
	}

	if capabilityFlags&consts.CLIENT_PROTOCOL_41 > 0 {
		if o.StatusFlags, err = buf.ReadU16(); err != nil {
			return
		}
		if o.Warnings, err = buf.ReadU16(); err != nil {
			return
		}
	} else if capabilityFlags&consts.CLIENT_TRANSACTIONS > 0 {
		if o.StatusFlags, err = buf.ReadU16(); err != nil {
			return
		}
	}

	return
}
