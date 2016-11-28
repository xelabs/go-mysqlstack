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
	"github.com/XeLabs/go-mysqlstack/consts"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAuth(t *testing.T) {
	auth := NewAuth()
	{
		data := []byte{
			0x8d, 0xa6, 0xff, 0x01, 0x00, 0x00, 0x00, 0x01,
			0x21, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x72, 0x6f, 0x6f, 0x74, 0x00, 0x14, 0x0e, 0xb4,
			0xdd, 0xb5, 0x5b, 0x64, 0xf8, 0x54, 0x40, 0xfd,
			0xf3, 0x45, 0xfa, 0x37, 0x12, 0x20, 0x20, 0xda,
			0x38, 0xaa, 0x61, 0x62, 0x63, 0x00, 0x6d, 0x79,
			0x73, 0x71, 0x6c, 0x5f, 0x6e, 0x61, 0x74, 0x69,
			0x76, 0x65, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77,
			0x6f, 0x72, 0x64, 0x00}

		auth.UnPack(data)
		want := &Auth{
			charset:         33,
			maxPacketSize:   16777216,
			authResponseLen: 20,
			authResponse: []byte{
				0x0e, 0xb4, 0xdd, 0xb5, 0x5b, 0x64, 0xf8, 0x54,
				0x40, 0xfd, 0xf3, 0x45, 0xfa, 0x37, 0x12, 0x20,
				0x20, 0xda, 0x38, 0xaa},
			pluginName:  "mysql_native_password",
			database:    "abc",
			user:        "root",
			clientFlags: 33531533,
		}
		got := auth
		assert.Equal(t, want, got)
	}

	{
		want := "abc"
		got := auth.Database()
		assert.Equal(t, want, got)
	}

	{
		want := uint32(33531533)
		got := auth.ClientFlags()
		assert.Equal(t, want, got)
	}

	{
		want := uint8(33)
		got := auth.Charset()
		assert.Equal(t, want, got)
	}
}

func TestAuthUnpackError(t *testing.T) {
	auth := NewAuth()
	{
		data := []byte{
			0x8d, 0xa6, 0xff,
		}
		err := auth.UnPack(data)
		want := "EOF"
		got := err.Error()
		assert.Equal(t, want, got)
	}
}

func TestAuthUnPack(t *testing.T) {
	want := NewAuth()
	want.charset = 0x02
	want.authResponseLen = 20
	want.clientFlags = DefaultCapability
	want.authResponse = nativePassword("sbtest", DefaultSalt)
	want.database = "sbtest"
	want.user = "sbtest"
	want.pluginName = DefaultAuthPluginName

	got := NewAuth()
	err := got.UnPack(want.Pack(
		DefaultCapability,
		0x02,
		"sbtest",
		"sbtest",
		DefaultSalt,
		"sbtest",
	))
	assert.Nil(t, err)
	assert.Equal(t, want, got)
}

func TestAuthWithoutPWD(t *testing.T) {
	want := NewAuth()
	want.charset = 0x02
	want.authResponseLen = 1
	want.clientFlags = DefaultCapability
	want.authResponse = nativePassword("", DefaultSalt)
	want.database = "sbtest"
	want.user = "sbtest"
	want.pluginName = DefaultAuthPluginName

	got := NewAuth()
	err := got.UnPack(want.Pack(
		DefaultCapability,
		0x02,
		"sbtest",
		"",
		DefaultSalt,
		"sbtest",
	))
	assert.Nil(t, err)
	assert.Equal(t, want, got)
}

func TestAuthWithoutDB(t *testing.T) {
	want := NewAuth()
	want.charset = 0x02
	want.authResponseLen = 20
	want.clientFlags = DefaultCapability &^ consts.CLIENT_CONNECT_WITH_DB
	want.authResponse = nativePassword("sbtest", DefaultSalt)
	want.user = "sbtest"
	want.pluginName = DefaultAuthPluginName

	got := NewAuth()
	err := got.UnPack(want.Pack(
		DefaultCapability,
		0x02,
		"sbtest",
		"sbtest",
		DefaultSalt,
		"",
	))
	assert.Nil(t, err)
	assert.Equal(t, want, got)
}

func TestAuthWithoutSecure(t *testing.T) {
	want := NewAuth()
	want.charset = 0x02
	want.authResponseLen = 20
	want.clientFlags = DefaultCapability &^ consts.CLIENT_SECURE_CONNECTION &^ consts.CLIENT_PLUGIN_AUTH_LENENC_CLIENT_DATA
	want.authResponse = nativePassword("sbtest", DefaultSalt)
	want.user = "sbtest"
	want.database = "sbtest"
	want.pluginName = DefaultAuthPluginName

	got := NewAuth()
	err := got.UnPack(want.Pack(
		DefaultCapability&^consts.CLIENT_SECURE_CONNECTION,
		0x02,
		"sbtest",
		"sbtest",
		DefaultSalt,
		"sbtest",
	))
	got.authResponseLen = 20
	assert.Nil(t, err)
	assert.Equal(t, want, got)
}
