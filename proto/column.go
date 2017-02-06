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

	querypb "github.com/XeLabs/go-mysqlstack/sqlparser/depends/query"
	"github.com/XeLabs/go-mysqlstack/sqlparser/depends/sqltypes"
)

func ColumnCount(payload []byte) (count uint64, err error) {
	buff := common.ReadBuffer(payload)
	if count, err = buff.ReadLenEncode(); err != nil {
		return
	}
	return
}

// http://dev.mysql.com/doc/internals/en/com-query-response.html#packet-Protocol::ColumnDefinition41
func UnpackColumn(payload []byte) (field *querypb.Field, err error) {
	field = &querypb.Field{}
	buff := common.ReadBuffer(payload)
	// Catalog is ignored, always set to "def"
	if _, err = buff.ReadLenEncodeString(); err != nil {
		return
	}

	// lenenc_str Schema
	if field.Database, err = buff.ReadLenEncodeString(); err != nil {
		return
	}

	// lenenc_str Table
	if field.Table, err = buff.ReadLenEncodeString(); err != nil {
		return
	}

	// lenenc_str Org_Table
	if field.OrgTable, err = buff.ReadLenEncodeString(); err != nil {
		return
	}

	// lenenc_str Name
	if field.Name, err = buff.ReadLenEncodeString(); err != nil {
		return
	}

	// lenenc_str Org_Name
	if field.OrgName, err = buff.ReadLenEncodeString(); err != nil {
		return
	}

	// lenenc_int length of fixed-length fields [0c], skip
	if _, err = buff.ReadLenEncode(); err != nil {
		return
	}

	// 2 character set
	charset, err := buff.ReadU16()
	if err != nil {
		return
	}
	field.Charset = uint32(charset)

	// 4 column length
	if field.ColumnLength, err = buff.ReadU32(); err != nil {
		return
	}

	// 1 type
	t, err := buff.ReadU8()
	if err != nil {
		return
	}

	// 2 flags
	flags, err := buff.ReadU16()
	if err != nil {
		return
	}
	field.Flags = uint32(flags)

	// Convert MySQL type
	if field.Type, err = sqltypes.MySQLToType(int64(t), int64(field.Flags)); err != nil {
		return
	}

	// 1 Decimals
	decimals, err := buff.ReadU8()
	if err != nil {
		return
	}
	field.Decimals = uint32(decimals)

	// 2 Filler and Default Values is ignored

	return
}

func PackColumn(field *querypb.Field) []byte {
	typ, flags := sqltypes.TypeToMySQL(field.Type)
	if field.Flags != 0 {
		flags = int64(field.Flags)
	}

	buf := common.NewBuffer(256)

	// lenenc_str Catalog, always 'def'
	buf.WriteLenEncodeString("def")

	// lenenc_str Schema
	buf.WriteLenEncodeString(field.Database)

	// lenenc_str Table
	buf.WriteLenEncodeString(field.Table)

	// lenenc_str Org_Table
	buf.WriteLenEncodeString(field.OrgTable)

	// lenenc_str Name
	buf.WriteLenEncodeString(field.Name)

	// lenenc_str Org_Name
	buf.WriteLenEncodeString(field.OrgName)

	// lenenc_int length of fixed-length fields [0c]
	buf.WriteLenEncode(uint64(0x0c))

	// 2 character set
	buf.WriteU16(uint16(field.Charset))

	// 4 column length
	buf.WriteU32(field.ColumnLength)

	// 1 type
	buf.WriteU8(byte(typ))

	// 2 flags
	buf.WriteU16(uint16(flags))

	//1 Decimals
	buf.WriteU8(uint8(field.Decimals))

	// 2 filler [00] [00]
	buf.WriteU16(uint16(0))

	return buf.Datas()
}
