/*
 * go-mysqlstack
 * xelabs.org
 *
 * Copyright (c) XeLabs
 * GPL License
 *
 */

package driver

import (
	"github.com/XeLabs/go-mysqlstack/consts"
	"github.com/XeLabs/go-mysqlstack/packet"
	"github.com/XeLabs/go-mysqlstack/proto"
	"github.com/pkg/errors"
	"net"
)

type Conn struct {
	netConn  net.Conn
	greeting *proto.Greeting
	auth     *proto.Auth
	packets  *packet.Packets
}

func NewConn(username, password, protocol, address, database string) (*Conn, error) {
	netconn, err := net.Dial(protocol, address)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	conn := &Conn{
		netConn:  netconn,
		greeting: proto.NewGreeting(0),
		auth:     proto.NewAuth(),
		packets:  packet.NewPackets(netconn),
	}

	{
		// greeting read
		payload, err := conn.packets.Next()
		if err != nil {
			return nil, err
		}

		// greeting unpack
		err = conn.greeting.UnPack(payload)
		if err != nil {
			return nil, err
		}
	}

	{
		// auth pack
		payload := conn.auth.Pack(
			conn.greeting.Capability,
			conn.greeting.Charset,
			username,
			password,
			conn.greeting.Salt,
			database,
		)

		// auth write
		err := conn.packets.Write(payload)
		if err != nil {
			return nil, err
		}
	}

	{
		// read
		payload, err := conn.packets.Next()
		if err != nil {
			return nil, err
		}

		if payload[0] != proto.OK_PACKET {
			pkt, err := conn.packets.ParseERR(payload, conn.greeting.Capability)
			if err != nil {
				return nil, err
			}
			return nil, errors.New(pkt.ErrorMessage)
		}
	}

	return conn, nil
}

// Close closes the connection
func (c *Conn) Close() error {
	if c.netConn != nil {
		if err := c.packets.WriteCommand(consts.COM_QUIT, nil); err != nil {
			return err
		}

		if c.netConn != nil {
			if err := c.netConn.Close(); err != nil {
				return errors.WithStack(err)
			}
			c.netConn = nil
		}
	}

	return nil
}
