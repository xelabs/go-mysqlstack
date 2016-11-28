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
)

type Greeting struct {
	protocolVersion uint8
	Charset         uint8
	status          uint16
	Capability      uint32
	connectionID    uint32
	Salt            []byte
	serverVersion   string
	authPluginName  string
}

func NewGreeting(connectionID uint32) *Greeting {
	greeting := &Greeting{
		protocolVersion: 10,
		serverVersion:   "Radon 5.7",
		connectionID:    connectionID,
		Charset:         consts.CHARSET_UTF8,
		status:          consts.SERVER_STATUS_AUTOCOMMIT,
	}

	greeting.Capability = DefaultCapability
	greeting.status |= consts.SERVER_STATUS_AUTOCOMMIT
	greeting.Salt = []byte{
		0x77, 0x63, 0x6a, 0x6d, 0x61, 0x22, 0x23, 0x27, // first part
		0x38, 0x26, 0x55, 0x58, 0x3b, 0x5d, 0x44, 0x78, 0x53, 0x73, 0x6b, 0x41,
		0x00}

	return greeting
}

func (g *Greeting) Status() uint16 {
	return g.status
}

// https://dev.mysql.com/doc/internals/en/connection-phase-packets.html#packet-Protocol::HandshakeV10
func (g *Greeting) Pack() []byte {
	// greeting buffer
	buf := common.NewBuffer(256)
	capabilityLo := uint16(g.Capability)
	capabilityHi := uint16(uint32(g.Capability) >> 16)

	// 1: [0a] protocol version
	buf.WriteU8(g.protocolVersion)

	// string[NUL]: server version
	buf.WriteString(g.serverVersion, len(g.serverVersion))
	buf.WriteZero(1)

	// 4: connection id
	buf.WriteU32(g.connectionID)

	// string[8]: auth-plugin-data-part-1
	buf.WriteBytes(g.Salt, 8)

	// 1: [00] filler
	buf.WriteZero(1)

	// 2: capability flags (lower 2 bytes)
	buf.WriteU16(capabilityLo)

	// 1: character set
	buf.WriteU8(consts.CHARSET_UTF8)

	// 2: status flags
	buf.WriteU16(g.status)

	// 2: capability flags (upper 2 bytes)
	buf.WriteU16(capabilityHi)

	// 1: length of auth-plugin-data-part-1
	if (g.Capability & consts.CLIENT_PLUGIN_AUTH) > 0 {
		buf.WriteU8(uint8(len(g.Salt)))
	} else {
		buf.WriteZero(1)
	}

	// string[10]: reserved (all [00])
	buf.WriteZero(10)

	// string[$len]: auth-plugin-data-part-2 ($len=MAX(13, length of auth-plugin-data - 8))
	if (g.Capability & consts.CLIENT_SECURE_CONNECTION) > 0 {
		buf.WriteBytes(g.Salt[8:], len(g.Salt)-8)
	}

	// string[NUL]    auth-plugin name
	if (g.Capability & consts.CLIENT_PLUGIN_AUTH) > 0 {
		pluginName := "mysql_native_password"
		buf.WriteString(pluginName, len(pluginName))
		buf.WriteZero(1)
	}

	return buf.Datas()
}

func (g *Greeting) UnPack(payload []byte) (err error) {
	buf := common.ReadBuffer(payload)

	// 1: [0a] protocol version
	if g.protocolVersion, err = buf.ReadU8(); err != nil {
		return
	}

	// string[NUL]: server version
	if g.serverVersion, err = buf.ReadStringNUL(); err != nil {
		return
	}

	// 4: connection id
	if g.connectionID, err = buf.ReadU32(); err != nil {
		return
	}

	// string[8]: auth-plugin-data-part-1
	var salt8 []byte
	if salt8, err = buf.ReadBytes(8); err != nil {
		return
	}
	copy(g.Salt, salt8)

	// 1: [00] filler
	if err = buf.ReadZero(1); err != nil {
		return
	}

	// 2: capability flags (lower 2 bytes)
	var CL uint16
	if CL, err = buf.ReadU16(); err != nil {
		return
	}

	// 1: character set
	if g.Charset, err = buf.ReadU8(); err != nil {
		return
	}

	// 2: status flags
	if g.status, err = buf.ReadU16(); err != nil {
		return
	}

	// 2: capability flags (upper 2 bytes)
	var CH uint16
	if CH, err = buf.ReadU16(); err != nil {
		return
	}
	g.Capability = (uint32(CH) << 16) | (uint32(CL))

	// 1: length of auth-plugin-data-part-1
	var SLEN uint8
	if (g.Capability & consts.CLIENT_PLUGIN_AUTH) > 0 {
		if SLEN, err = buf.ReadU8(); err != nil {
			return
		}
	} else {
		if err = buf.ReadZero(1); err != nil {
			return
		}
	}

	// string[10]: reserved (all [00])
	if err = buf.ReadZero(10); err != nil {
		return
	}

	// string[$len]: auth-plugin-data-part-2 ($len=MAX(13, length of auth-plugin-data - 8))
	if (g.Capability & consts.CLIENT_SECURE_CONNECTION) > 0 {
		var salt []byte

		read := int(SLEN) - 8
		if read < 0 {
			read = 13
		}

		if salt, err = buf.ReadBytes(read); err != nil {
			return
		}
		copy(g.Salt[8:], salt)
	}

	// string[NUL]    auth-plugin name
	if (g.Capability & consts.CLIENT_PLUGIN_AUTH) > 0 {
		if g.authPluginName, err = buf.ReadStringNUL(); err != nil {
			return
		}
	}

	return nil
}
