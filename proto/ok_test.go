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
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOK(t *testing.T) {
	{
		buff := common.NewBuffer(32)

		// header
		buff.WriteU8(0x00)
		// affected_rows
		buff.WriteLenEncode(uint64(3))
		// last_insert_id
		buff.WriteLenEncode(uint64(40000000000))

		// status_flags
		buff.WriteU16(0x01)
		// warnings
		buff.WriteU16(0x02)

		want := NewOK()
		want.AffectedRows = 3
		want.LastInsertID = 40000000000
		want.StatusFlags = 1
		want.Warnings = 2

		got := NewOK()
		err := got.UnPack(buff.Datas(), DefaultCapability)
		assert.Nil(t, err)
		assert.Equal(t, want, got)
	}

	// &^ CLIENT_PROTOCOL_41
	{
		buff := common.NewBuffer(32)

		// header
		buff.WriteU8(0x00)
		// affected_rows
		buff.WriteLenEncode(uint64(3))
		// last_insert_id
		buff.WriteLenEncode(uint64(40000000000))

		// status_flags
		buff.WriteU16(0x03)

		want := NewOK()
		want.AffectedRows = 3
		want.LastInsertID = 40000000000
		want.StatusFlags = 3

		got := NewOK()
		err := got.UnPack(buff.Datas(), DefaultCapability&^consts.CLIENT_PROTOCOL_41)
		assert.Nil(t, err)
		assert.Equal(t, want, got)
	}
}

func TestOKError(t *testing.T) {
	{
		buff := common.NewBuffer(32)

		// header
		buff.WriteU8(0x01)
		// affected_rows
		buff.WriteLenEncode(uint64(3))
		// last_insert_id
		buff.WriteLenEncode(uint64(40000000000))

		// status_flags
		buff.WriteU16(0x01)
		// warnings
		buff.WriteU16(0x02)

		got := NewOK()
		err := got.UnPack(buff.Datas(), DefaultCapability)
		assert.NotNil(t, err)
	}
}
