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

type Conn interface {
	Cleanup()
	Close() error
	Closed() bool
	NextPacket() ([]byte, error)
	Query(sql string) (rows Rows, err error)
	Exec(sql string) (rows Rows, err error)
}

type conn struct {
	netConn  net.Conn
	greeting *proto.Greeting
	auth     *proto.Auth
	packets  *packet.Packets
}

func (c *conn) handleErrorPacket(data []byte) error {
	if data[0] == proto.ERR_PACKET {
		pkt, e := c.packets.ParseERR(data, c.greeting.Capability)
		if e != nil {
			c.Cleanup()
			return e
		}
		return errors.New(pkt.ErrorMessage)
	}

	return nil
}

func NewConn(username, password, address, database string) (c *conn, err error) {
	var payload []byte

	c = &conn{}
	if c.netConn, err = net.Dial("tcp", address); err != nil {
		return nil, errors.WithStack(err)
	}

	c.auth = proto.NewAuth()
	c.greeting = proto.NewGreeting(0)
	c.packets = packet.NewPackets(c.netConn)

	{

		// greeting read
		if payload, err = c.packets.Next(); err != nil {
			c.Cleanup()
			return
		}
		if err = c.handleErrorPacket(payload); err != nil {
			return
		}

		// greeting unpack
		if err = c.greeting.UnPack(payload); err != nil {
			c.Cleanup()
			return
		}
	}

	{
		// auth pack
		payload := c.auth.Pack(
			c.greeting.Capability,
			c.greeting.Charset,
			username,
			password,
			c.greeting.Salt,
			database,
		)

		// auth write
		if err = c.packets.Write(payload); err != nil {
			c.Cleanup()
			return
		}
	}

	{
		// read
		if payload, err = c.packets.Next(); err != nil {
			c.Cleanup()
			return
		}

		if err = c.handleErrorPacket(payload); err != nil {
			return
		}
	}

	return c, nil
}

func (c *conn) NextPacket() ([]byte, error) {
	return c.packets.Next()
}

func (c *conn) Cleanup() {
	if c.netConn != nil {
		c.netConn.Close()
		c.netConn = nil
	}
}

// Close closes the connection
func (c *conn) Close() error {
	if c.netConn != nil {
		if err := c.packets.WriteCommand(consts.COM_QUIT, nil); err != nil {
			return err
		}
		c.Cleanup()
	}

	return nil
}

// Closed checks the connection broken or not
func (c *conn) Closed() bool {
	return c.netConn == nil
}
