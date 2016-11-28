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
	catalog    string
	schema     string
	table      string
	org_table  string
	name       string
	org_name   string
	charset    uint16
	columnLen  uint32
	fieldType  uint8
	fieldFlags uint16
	decimals   uint8
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
	//lenenc_str catalog
	if c.catalog, err = buff.ReadLenEncodeString(); err != nil {
		return
	}

	// lenenc_str schema
	if c.schema, err = buff.ReadLenEncodeString(); err != nil {
		return
	}

	// lenenc_str table
	if c.table, err = buff.ReadLenEncodeString(); err != nil {
		return
	}

	// lenenc_str org_table
	if c.org_table, err = buff.ReadLenEncodeString(); err != nil {
		return
	}

	// lenenc_str name
	if c.name, err = buff.ReadLenEncodeString(); err != nil {
		return
	}

	// lenenc_str org_name
	if c.org_name, err = buff.ReadLenEncodeString(); err != nil {
		return
	}

	// lenenc_int length of fixed-length fields [0c]
	if _, err = buff.ReadLenEncode(); err != nil {
		return
	}

	// 2 character set
	if c.charset, err = buff.ReadU16(); err != nil {
		return
	}
	// 4 column length
	if c.columnLen, err = buff.ReadU32(); err != nil {
		return
	}

	// 1 type
	if c.fieldType, err = buff.ReadU8(); err != nil {
		return
	}

	// 2 flags
	if c.fieldFlags, err = buff.ReadU16(); err != nil {
		return
	}

	//1 decimals
	if c.decimals, err = buff.ReadU8(); err != nil {
		return
	}

	return
}
