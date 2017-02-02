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

	"github.com/XeLabs/go-mysqlstack/common"
	"github.com/stretchr/testify/assert"
)

func TestERR(t *testing.T) {
	{
		buff := common.NewBuffer(32)

		// header
		buff.WriteU8(0xff)
		// error_code
		buff.WriteU16(0x01)
		// sql_state_marker
		buff.WriteString("#")
		// sql_state
		buff.WriteString("ABCDE")
		buff.WriteString("ERROR")

		want := &ERR{}
		want.Header = 0xff
		want.ErrorCode = 0x1
		want.SQLState = "ABCDE"
		want.ErrorMessage = "ERROR"

		got, err := UnPackERR(buff.Datas())
		assert.Nil(t, err)
		assert.Equal(t, want, got)
	}

	{
		want := &ERR{}
		want.Header = 0xff
		want.ErrorCode = 0x1
		want.SQLState = "ABCDE"
		want.ErrorMessage = "ERROR"
		datas := PackERR(want)

		got, err := UnPackERR(datas)
		assert.Nil(t, err)
		assert.Equal(t, want, got)
	}
}

func TestERRError(t *testing.T) {
	{
		buff := common.NewBuffer(32)

		// header
		buff.WriteU8(0x01)

		_, err := UnPackERR(buff.Datas())
		assert.NotNil(t, err)
	}
}
