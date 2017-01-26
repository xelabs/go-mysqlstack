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
)

type Column struct {
	Catalog      string
	Schema       string
	Table        string
	Org_Table    string
	Name         string
	Org_Name     string
	Charset      uint16
	ColumnLen    uint32
	FieldType    uint8
	FieldFlags   uint16
	Decimals     uint8
	Filler       uint16
	DefaultValue []byte
}

func ColumnCount(payload []byte) (count uint64, err error) {
	buff := common.ReadBuffer(payload)
	if count, err = buff.ReadLenEncode(); err != nil {
		return
	}
	return
}

// http://dev.mysql.com/doc/internals/en/com-query-response.html#packet-Protocol::ColumnDefinition41
func (c *Column) UnPack(payload []byte) (err error) {
	buff := common.ReadBuffer(payload)
	//lenenc_str Catalog
	if c.Catalog, err = buff.ReadLenEncodeString(); err != nil {
		return
	}

	// lenenc_str Schema
	if c.Schema, err = buff.ReadLenEncodeString(); err != nil {
		return
	}

	// lenenc_str Table
	if c.Table, err = buff.ReadLenEncodeString(); err != nil {
		return
	}

	// lenenc_str Org_Table
	if c.Org_Table, err = buff.ReadLenEncodeString(); err != nil {
		return
	}

	// lenenc_str Name
	if c.Name, err = buff.ReadLenEncodeString(); err != nil {
		return
	}

	// lenenc_str Org_Name
	if c.Org_Name, err = buff.ReadLenEncodeString(); err != nil {
		return
	}

	// lenenc_int length of fixed-length fields [0c]
	if _, err = buff.ReadLenEncode(); err != nil {
		return
	}

	// 2 character set
	if c.Charset, err = buff.ReadU16(); err != nil {
		return
	}
	// 4 column length
	if c.ColumnLen, err = buff.ReadU32(); err != nil {
		return
	}

	// 1 type
	if c.FieldType, err = buff.ReadU8(); err != nil {
		return
	}

	// 2 flags
	if c.FieldFlags, err = buff.ReadU16(); err != nil {
		return
	}

	//1 Decimals
	if c.Decimals, err = buff.ReadU8(); err != nil {
		return
	}

	//2 Filler
	if c.Filler, err = buff.ReadU16(); err != nil {
		return
	}

	// Default Values
	if buff.Seek() < buff.Length() {
		if c.DefaultValue, err = buff.ReadLenEncodeBytes(); err != nil {
			return
		}
	}

	return
}

func (c *Column) Pack() []byte {
	buf := common.NewBuffer(256)

	//lenenc_str Catalog
	buf.WriteLenEncodeString(c.Catalog)
	// lenenc_str Schema
	buf.WriteLenEncodeString(c.Schema)
	// lenenc_str Table
	buf.WriteLenEncodeString(c.Table)
	// lenenc_str Org_Table
	buf.WriteLenEncodeString(c.Org_Table)
	// lenenc_str Name
	buf.WriteLenEncodeString(c.Name)
	// lenenc_str Org_Name
	buf.WriteLenEncodeString(c.Org_Name)
	// lenenc_int length of fixed-length fields [0c]
	buf.WriteLenEncode(uint64(0x0c))
	// 2 character set
	buf.WriteU16(c.Charset)
	// 4 column length
	buf.WriteU32(c.ColumnLen)
	// 1 type
	buf.WriteU8(c.FieldType)
	// 2 flags
	buf.WriteU16(c.FieldFlags)
	//1 Decimals
	buf.WriteU8(c.Decimals)
	// 2 filler [00] [00]
	buf.WriteU16(c.Filler)
	if c.DefaultValue != nil {
		buf.WriteLenEncodeBytes(c.DefaultValue)
	}

	return buf.Datas()
}
