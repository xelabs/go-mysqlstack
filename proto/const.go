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
	"github.com/XeLabs/go-mysqlstack/sqldb"
)

const (
	DefaultAuthPluginName = "mysql_native_password"

	DefaultServerCapability = sqldb.CLIENT_LONG_PASSWORD |
		sqldb.CLIENT_LONG_FLAG |
		sqldb.CLIENT_CONNECT_WITH_DB |
		sqldb.CLIENT_PROTOCOL_41 |
		sqldb.CLIENT_TRANSACTIONS |
		sqldb.CLIENT_MULTI_STATEMENTS |
		sqldb.CLIENT_PLUGIN_AUTH |
		sqldb.CLIENT_DEPRECATE_EOF |
		sqldb.CLIENT_SECURE_CONNECTION

	DefaultClientCapability = sqldb.CLIENT_LONG_PASSWORD |
		sqldb.CLIENT_LONG_FLAG |
		sqldb.CLIENT_PROTOCOL_41 |
		sqldb.CLIENT_TRANSACTIONS |
		sqldb.CLIENT_MULTI_STATEMENTS |
		sqldb.CLIENT_PLUGIN_AUTH |
		sqldb.CLIENT_DEPRECATE_EOF |
		sqldb.CLIENT_SECURE_CONNECTION
)

var (
	DefaultSalt = []byte{
		0x77, 0x63, 0x6a, 0x6d, 0x61, 0x22, 0x23, 0x27, // first part
		0x38, 0x26, 0x55, 0x58, 0x3b, 0x5d, 0x44, 0x78, 0x53, 0x73, 0x6b, 0x41}
)
