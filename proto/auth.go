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
	"crypto/sha1"
	"github.com/XeLabs/go-mysqlstack/common"
	"github.com/XeLabs/go-mysqlstack/consts"
)

type Auth struct {
	charset         uint8
	sequence        uint8
	maxPacketSize   uint32
	authResponseLen uint8
	clientFlags     uint32
	authResponse    []byte
	pluginName      string
	database        string
	user            string
}

func NewAuth() *Auth {
	return &Auth{}
}

func (a *Auth) Database() string {
	return a.database
}

func (a *Auth) ClientFlags() uint32 {
	return a.clientFlags
}

func (a *Auth) Charset() uint8 {
	return a.charset
}

// https://dev.mysql.com/doc/internals/en/connection-phase-packets.html
// #packet-Protocol::HandshakeResponse41
func (a *Auth) UnPack(payload []byte) (err error) {
	buf := common.ReadBuffer(payload)

	if a.clientFlags, err = buf.ReadU32(); err != nil {
		return
	}

	if a.maxPacketSize, err = buf.ReadU32(); err != nil {
		return
	}

	if a.charset, err = buf.ReadU8(); err != nil {
		return
	}

	if err = buf.ReadZero(23); err != nil {
		return
	}

	if a.user, err = buf.ReadStringNUL(); err != nil {
		return
	}

	if (a.clientFlags & consts.CLIENT_SECURE_CONNECTION) > 0 {
		if a.authResponseLen, err = buf.ReadU8(); err != nil {
			return
		}

		if a.authResponse, err = buf.ReadBytes(int(a.authResponseLen)); err != nil {
			return
		}
	} else {
		if a.authResponse, err = buf.ReadBytesNUL(); err != nil {
			return
		}
	}

	if (a.clientFlags & consts.CLIENT_CONNECT_WITH_DB) > 0 {
		if a.database, err = buf.ReadStringNUL(); err != nil {
			return
		}
	}

	if (a.clientFlags & consts.CLIENT_PLUGIN_AUTH) > 0 {
		if a.pluginName, err = buf.ReadStringNUL(); err != nil {
			return
		}
	}

	return nil
}

func (a *Auth) Pack(
	capabilityFlags uint32,
	charset uint8,
	username string,
	password string,
	salt []byte,
	database string,
) []byte {
	buf := common.NewBuffer(256)
	authResponse := nativePassword(password, salt)

	// must always be set
	capabilityFlags |= consts.CLIENT_PROTOCOL_41

	// not supported
	capabilityFlags &= ^consts.CLIENT_SSL
	capabilityFlags &= ^consts.CLIENT_COMPRESS
	capabilityFlags &= ^consts.CLIENT_DEPRECATE_EOF
	capabilityFlags &= ^consts.CLIENT_CONNECT_ATTRS

	if len(database) > 0 {
		capabilityFlags |= consts.CLIENT_CONNECT_WITH_DB
	} else {
		capabilityFlags &= ^consts.CLIENT_CONNECT_WITH_DB
	}

	// 4 capability flags, CLIENT_PROTOCOL_41 always set
	buf.WriteU32(capabilityFlags)

	// 4 max-packet size (none)
	buf.WriteU32(0)

	// 1 character set
	buf.WriteU8(charset)

	// string[23] reserved (all [0])
	buf.WriteZero(23)

	// string[NUL] username
	buf.WriteString(username, len(username))
	buf.WriteZero(1)

	if (capabilityFlags & consts.CLIENT_SECURE_CONNECTION) > 0 {
		// 1 length of auth-response
		// string[n]  auth-response
		buf.WriteU8(uint8(len(authResponse)))
		buf.WriteBytes(authResponse, len(authResponse))
	} else {
		buf.WriteBytes(authResponse, len(authResponse))
		buf.WriteZero(1)
	}
	capabilityFlags &= ^consts.CLIENT_PLUGIN_AUTH_LENENC_CLIENT_DATA

	// string[NUL] database
	if capabilityFlags&consts.CLIENT_CONNECT_WITH_DB > 0 {
		buf.WriteString(database, len(database))
		buf.WriteZero(1)
	}

	// string[NUL] auth plugin name
	buf.WriteString(DefaultAuthPluginName, len(DefaultAuthPluginName))
	buf.WriteZero(1)

	// CLIENT_CONNECT_ATTRS none

	return buf.Datas()
}

// https://dev.mysql.com/doc/internals/en/secure-password-authentication.html#packet-Authentication::Native41
// SHA1( password ) XOR SHA1( "20-bytes random data from server" <concat> SHA1( SHA1( password ) ) )
func nativePassword(password string, salt []byte) []byte {
	if len(password) == 0 {
		return []byte{0x00}
	}

	hash := sha1.New()
	hash.Write([]byte(password))
	hashPass := hash.Sum(nil)

	hash.Reset()
	hash.Write(hashPass)
	doubleHashPass := hash.Sum(nil)

	saltLen := len(salt)
	if saltLen > 20 {
		saltLen = 20
	}
	hash.Reset()
	hash.Write(salt[:saltLen])
	hash.Write(doubleHashPass)
	salts := hash.Sum(nil)

	for i, b := range hashPass {
		hashPass[i] = b ^ salts[i]
	}

	return hashPass
}
