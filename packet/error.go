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
	"github.com/pkg/errors"
)

var (
	ErrBadConn       = errors.New("connection.was.bad")
	ErrMalformPacket = errors.New("Malform.packet.error")
)
