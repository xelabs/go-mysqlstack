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
	payload := []byte{
		0x03, 0x64, 0x65, 0x66, 0x06, 0x73, 0x62, 0x74,
		0x65, 0x73, 0x74, 0x02, 0x74, 0x31, 0x02, 0x74,
		0x31, 0x01, 0x61, 0x01, 0x61, 0x0c, 0x3f, 0x00,
		0x0b, 0x00, 0x00, 0x00, 0x03, 0x03, 0x50, 0x00,
		0x00, 0x00,
	}

	want := &Column{
		catalog:    "def",
		schema:     "sbtest",
		table:      "t1",
		org_table:  "t1",
		name:       "a",
		org_name:   "a",
		charset:    63,
		columnLen:  11,
		fieldType:  3,
		fieldFlags: 20483,
	}

	got := &Column{}
	err := got.UnPack(payload)
	assert.Nil(t, err)
	assert.Equal(t, want, got)
}
