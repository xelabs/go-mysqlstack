/*
 * go-mysqlstack
 * xelabs.org
 *
 * Copyright (c) XeLabs
 * GPL License
 *
 */

package packet

import (
	"errors"
)

var (
	ErrBadConn       = errors.New("connection.was.bad")
	ErrMalformPacket = errors.New("Malform.packet.error")
)
