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
	"github.com/stretchr/testify/assert"
	"testing"
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
	want := &Column{
		Catalog:    "def",
		Schema:     "sbtest",
		Table:      "t1",
		Org_Table:  "t1",
		Name:       "a",
		Org_Name:   "a",
		Charset:    63,
		ColumnLen:  11,
		FieldType:  3,
		FieldFlags: 20483,
	}

	got := &Column{}
	err := got.UnPack(want.Pack())
	assert.Nil(t, err)
	assert.Equal(t, want, got)
}
