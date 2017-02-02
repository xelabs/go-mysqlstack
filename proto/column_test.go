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
	"testing"

	"github.com/XeLabs/go-mysqlstack/sqlparser/depends/sqltypes"
	"github.com/stretchr/testify/assert"

	querypb "github.com/XeLabs/go-mysqlstack/sqlparser/depends/query"
)

func TestColumnCount(t *testing.T) {
	payload := []byte{
		0x02,
	}

	want := uint64(2)
	got, err := ColumnCount(payload)
	assert.Nil(t, err)
	assert.Equal(t, want, got)
}

func TestColumn(t *testing.T) {
	want := &querypb.Field{
		Database:     "test",
		Table:        "t1",
		OrgTable:     "t1",
		Name:         "a",
		OrgName:      "a",
		Charset:      11,
		ColumnLength: 11,
		Type:         sqltypes.Int32,
		Flags:        11,
	}

	datas := PackColumn(want)
	got, err := UnpackColumn(datas)
	assert.Nil(t, err)
	assert.Equal(t, want, got)
}
