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
		want.ErrorMessage = "ERROR"
		datas := PackERR(want)

		got, err := UnPackERR(datas)
		assert.Nil(t, err)
		assert.Equal(t, want, got)
	}
}

func TestERRUnPackError(t *testing.T) {
	// header error
	{
		buff := common.NewBuffer(32)

		// header
		buff.WriteU8(0x01)

		_, err := UnPackERR(buff.Datas())
		assert.NotNil(t, err)
	}

	// NULL
	f0 := func(buff *common.Buffer) {
	}

	// Write error header.
	f1 := func(buff *common.Buffer) {
		buff.WriteU8(0xff)
	}

	// Write error code.
	f2 := func(buff *common.Buffer) {
		buff.WriteU16(0x01)
	}

	// Write SQLStateMarker.
	f3 := func(buff *common.Buffer) {
		buff.WriteU8('#')
	}

	// Write SQLState.
	f4 := func(buff *common.Buffer) {
		buff.WriteString("xxxxx")
	}

	buff := common.NewBuffer(32)
	fs := []func(buff *common.Buffer){f0, f1, f2, f3, f4}
	for i := 0; i < len(fs); i++ {
		_, err := UnPackERR(buff.Datas())
		assert.NotNil(t, err)
		fs[i](buff)
	}
}
