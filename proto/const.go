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
)

const (
	DefaultAuthPluginName = "mysql_native_password"
	DefaultCapability     = consts.CLIENT_LONG_PASSWORD |
		consts.CLIENT_LONG_FLAG |
		consts.CLIENT_CONNECT_WITH_DB |
		consts.CLIENT_PROTOCOL_41 |
		consts.CLIENT_TRANSACTIONS |
		consts.CLIENT_MULTI_STATEMENTS |
		consts.CLIENT_PLUGIN_AUTH |
		consts.CLIENT_SECURE_CONNECTION
)

var (
	DefaultSalt = []byte{
		0x77, 0x63, 0x6a, 0x6d, 0x61, 0x22, 0x23, 0x27, // first part
		0x38, 0x26, 0x55, 0x58, 0x3b, 0x5d, 0x44, 0x78, 0x53, 0x73, 0x6b, 0x41,
		0x00}
)
