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

func TestERR(t *testing.T) {
	{
		buff := common.NewBuffer(32)

		// header
		buff.WriteU8(0xff)
		// error_code
		buff.WriteU16(0x01)
		// sql_state_marker
		buff.WriteString("a", 1)
		// sql_state
		buff.WriteString("ABCDE", 5)
		buff.WriteString("ERROR", 5)

		want := NewERR()
		want.Header = 0xff
		want.ErrorCode = 0x1
		want.SQLStateMarker = "a"
		want.SQLState = "ABCDE"
		want.ErrorMessage = "ERROR"

		got := NewERR()
		err := got.UnPack(buff.Datas(), DefaultCapability)
		assert.Nil(t, err)
		assert.Equal(t, want, got)
	}

	// &^ CLIENT_PROTOCOL_41
	{
		buff := common.NewBuffer(32)

		// header
		buff.WriteU8(0xff)
		// error_code
		buff.WriteU16(0x01)
		// sql_state_marker
		buff.WriteString("ERROR", 5)

		want := NewERR()
		want.Header = 0xff
		want.ErrorCode = 0x1
		want.ErrorMessage = "ERROR"

		got := NewERR()
		err := got.UnPack(buff.Datas(), DefaultCapability&^consts.CLIENT_PROTOCOL_41)
		assert.Nil(t, err)
		assert.Equal(t, want, got)
	}
}

func TestERRError(t *testing.T) {
	{
		buff := common.NewBuffer(32)

		// header
		buff.WriteU8(0x01)

		got := NewERR()
		err := got.UnPack(buff.Datas(), DefaultCapability)
		assert.NotNil(t, err)
	}
}
