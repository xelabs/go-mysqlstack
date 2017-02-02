/*
 * go-mysqlstack
 * xelabs.org
 *
 * Copyright (c) XeLabs
 * GPL License
 *
 */

package consts

// https://dev.mysql.com/doc/internals/en/capability-flags.html
// include/mysql_com.h
const (
	// new more secure password
	CLIENT_LONG_PASSWORD = 1

	// Found instead of affected rows
	CLIENT_FOUND_ROWS = uint32(1 << 1)

	// Get all column flags
	CLIENT_LONG_FLAG = uint32(1 << 2)

	// One can specify db on connect
	CLIENT_CONNECT_WITH_DB = uint32(1 << 3)

	// Don't allow database.table.column
	CLIENT_NO_SCHEMA = uint32(1 << 4)

	// Can use compression protocol
	CLIENT_COMPRESS = uint32(1 << 5)

	// Odbc client
	CLIENT_ODBC = uint32(1 << 6)

	// Can use LOAD DATA LOCAL
	CLIENT_LOCAL_FILES = uint32(1 << 7)

	// Ignore spaces before '('
	CLIENT_IGNORE_SPACE = uint32(1 << 8)

	// New 4.1 protocol
	CLIENT_PROTOCOL_41 = uint32(1 << 9)

	// This is an interactive client
	CLIENT_INTERACTIVE = uint32(1 << 10)

	// Switch to SSL after handshake
	CLIENT_SSL = uint32(1 << 11)

	// IGNORE sigpipes
	CLIENT_IGNORE_SIGPIPE = uint32(1 << 12)

	// Client knows about transactions
	CLIENT_TRANSACTIONS = uint32(1 << 13)

	// Old flag for 4.1 protocol
	CLIENT_RESERVED = uint32(1 << 14)

	// Old flag for 4.1 authentication
	CLIENT_SECURE_CONNECTION = uint32(1 << 15)

	// Enable/disable multi-stmt support
	CLIENT_MULTI_STATEMENTS = uint32(1 << 16)

	// Enable/disable multi-results
	CLIENT_MULTI_RESULTS = uint32(1 << 17)

	// Multi-results in PS-protocol
	CLIENT_PS_MULTI_RESULTS = uint32(1 << 18)

	// Client supports plugin authentication
	CLIENT_PLUGIN_AUTH = uint32(1 << 19)

	// Client supports connection attributes
	CLIENT_CONNECT_ATTRS = uint32(1 << 20)

	//  Enable authentication response packet to be larger than 255 bytes
	CLIENT_PLUGIN_AUTH_LENENC_CLIENT_DATA = uint32(1 << 21)

	// Don't close the connection for a connection with expired password
	CLIENT_CAN_HANDLE_EXPIRED_PASSWORDS = uint32(1 << 22)

	// Capable of handling server state change information. Its a hint to the
	// server to include the state change information in Ok packet.
	CLIENT_SESSION_TRACK = uint32(1 << 23)

	//Client no longer needs EOF packet
	CLIENT_DEPRECATE_EOF = uint32(1 << 24)
)
