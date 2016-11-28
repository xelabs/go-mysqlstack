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

func TestGreetingUnPack(t *testing.T) {
	want := NewGreeting(4)
	got := NewGreeting(4)

	// normal
	{
		want.authPluginName = "mysql_native_password"
		err := got.UnPack(want.Pack())
		assert.Nil(t, err)
		assert.Equal(t, want, got)
	}

	// 1. off consts.CLIENT_PLUGIN_AUTH
	{
		want.Capability = want.Capability &^ consts.CLIENT_PLUGIN_AUTH
		want.authPluginName = "mysql_native_password"
		err := got.UnPack(want.Pack())
		assert.Nil(t, err)
		assert.Equal(t, want, got)
	}

	// 2. off consts.CLIENT_SECURE_CONNECTION
	{
		want.Capability &= ^consts.CLIENT_SECURE_CONNECTION
		want.authPluginName = "mysql_native_password"
		err := got.UnPack(want.Pack())
		assert.Nil(t, err)
		assert.Equal(t, want, got)
	}

	// 3. off consts.CLIENT_PLUGIN_AUTH && consts.CLIENT_SECURE_CONNECTION
	{
		want.Capability &= (^consts.CLIENT_PLUGIN_AUTH ^ consts.CLIENT_SECURE_CONNECTION)
		want.authPluginName = "mysql_native_password"
		err := got.UnPack(want.Pack())
		assert.Nil(t, err)
		assert.Equal(t, want, got)
	}
}
