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
	"net"

	"github.com/XeLabs/go-mysqlstack/common"
	"github.com/XeLabs/go-mysqlstack/consts"
	"github.com/XeLabs/go-mysqlstack/packet"
	"github.com/XeLabs/go-mysqlstack/proto"
	"github.com/pkg/errors"

	querypb "github.com/XeLabs/go-mysqlstack/sqlparser/depends/query"
	"github.com/XeLabs/go-mysqlstack/sqlparser/depends/sqltypes"
)

type Conn interface {
	Cleanup()
	Close() error
	Closed() bool
	NextPacket() ([]byte, error)

	// Query gets the row iterator
	Query(sql string) (Rows, error)
	Exec(sql string) error

	// FetchAll fetchs all results
	FetchAll(sql string, maxrows int) (*sqltypes.Result, error)
}

type conn struct {
	netConn  net.Conn
	auth     *proto.Auth
	greeting *proto.Greeting
	packets  *packet.Packets
}

func (c *conn) handleErrorPacket(data []byte) error {
	if data[0] == proto.ERR_PACKET {
		pkt, e := c.packets.ParseERR(data)
		if e != nil {
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

	defer func() {
		if err != nil {
			c.Cleanup()
		}
	}()

	c.auth = proto.NewAuth()
	c.greeting = proto.NewGreeting(0)
	c.packets = packet.NewPackets(c.netConn)

	{

		// greeting read
		if payload, err = c.packets.Next(); err != nil {
			return
		}

		// check greeting packet
		if err = c.handleErrorPacket(payload); err != nil {
			return
		}

		// unpack greeting packet
		if err = c.greeting.UnPack(payload); err != nil {
			return
		}

		// check greating Capability
		if c.greeting.Capability&consts.CLIENT_PROTOCOL_41 == 0 {
			err = errors.New("cannot.connect.to.servers.earlier.than.4.1")
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
			return
		}
	}

	{
		// read
		if payload, err = c.packets.Next(); err != nil {
			return
		}

		if err = c.handleErrorPacket(payload); err != nil {
			return
		}
	}

	return c, nil
}

func (c *conn) query(command byte, sql string) (r Rows, err error) {
	var ok *proto.OK
	var columns []*querypb.Field

	defer func() {
		if err != nil {
			c.Cleanup()
		}
	}()

	// query
	if err = c.packets.WriteCommand(command, common.StringToBytes(sql)); err != nil {
		return
	}
	// columns
	if columns, ok, err = c.packets.ReadColumns(); err != nil {
		return
	}

	rows := NewTextRows(c)
	rows.rowsAffected = ok.AffectedRows
	rows.insertID = ok.LastInsertID
	rows.fields = columns

	return rows, nil
}

// Query execute the query and return the row iterator
func (c *conn) Query(sql string) (Rows, error) {
	return c.query(consts.COM_QUERY, sql)
}

func (c *conn) Ping() (err error) {
	var rows Rows

	defer func() {
		if err != nil {
			c.Cleanup()
		}
	}()

	if rows, err = c.query(consts.COM_PING, ""); err != nil {
		return
	}

	if err = rows.Close(); err != nil {
		return
	}
	return
}

func (c *conn) InitDB(db string) (err error) {
	var rows Rows

	defer func() {
		if err != nil {
			c.Cleanup()
		}
	}()

	if rows, err = c.query(consts.COM_INIT_DB, db); err != nil {
		return
	}

	if err = rows.Close(); err != nil {
		return
	}
	return
}

// Exec executes the query and drain the results
func (c *conn) Exec(sql string) (err error) {
	var rows Rows

	defer func() {
		if err != nil {
			c.Cleanup()
		}
	}()

	if rows, err = c.query(consts.COM_QUERY, sql); err != nil {
		return
	}

	if err = rows.Close(); err != nil {
		return
	}
	return
}

func (c *conn) FetchAll(sql string, maxrows int) (r *sqltypes.Result, err error) {
	var ok *proto.OK
	var columns []*querypb.Field

	defer func() {
		if err != nil {
			c.Cleanup()
		}
	}()

	// Write query request
	if err = c.packets.WriteCommand(consts.COM_QUERY, common.StringToBytes(sql)); err != nil {
		return
	}
	// Read column definition
	if columns, ok, err = c.packets.ReadColumns(); err != nil {
		return
	}

	r = &sqltypes.Result{
		Fields:       columns,
		RowsAffected: ok.AffectedRows,
		InsertID:     ok.LastInsertID,
	}

	// Read results
	rows := NewTextRows(c)
	rows.fields = columns
	for rows.Next() {
		if len(r.Rows) == maxrows {
			break
		}
		row, err := rows.RowValues()
		if err != nil {
			return nil, err
		}
		r.Rows = append(r.Rows, row)
	}
	// Check last error
	if err = rows.Close(); err != nil {
		return
	}

	return
}

// NextPacket used to get the next packet
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
