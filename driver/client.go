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
	"time"

	"github.com/XeLabs/go-mysqlstack/common"
	"github.com/XeLabs/go-mysqlstack/packet"
	"github.com/XeLabs/go-mysqlstack/proto"
	"github.com/XeLabs/go-mysqlstack/sqldb"

	querypb "github.com/XeLabs/go-mysqlstack/sqlparser/depends/query"
	"github.com/XeLabs/go-mysqlstack/sqlparser/depends/sqltypes"
)

var _ Conn = &conn{}

type Conn interface {
	Ping() error
	Quit()
	Close() error
	Closed() bool
	Cleanup()

	// ConnectionID is the connection id at greeting
	ConnectionID() uint32
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
		return c.packets.ParseERR(data)
	}
	return nil
}

func (c *conn) handShake(username, password, database string) error {
	var err error
	var data []byte

	//Parses the initial handshake from the server.
	{
		// greeting read
		if data, err = c.packets.Next(); err != nil {
			return err
		}

		// check greeting packet
		if err = c.handleErrorPacket(data); err != nil {
			return err
		}

		// unpack greeting packet
		if err = c.greeting.UnPack(data); err != nil {
			return err
		}

		// check greating Capability
		if c.greeting.Capability&sqldb.CLIENT_PROTOCOL_41 == 0 {
			err = sqldb.NewSQLError(sqldb.CR_VERSION_ERROR, "cannot connect to servers earlier than 4.1")
			return err
		}
	}

	{
		// auth pack
		data := c.auth.Pack(
			proto.DefaultClientCapability,
			c.greeting.Charset,
			username,
			password,
			c.greeting.Salt,
			database,
		)

		// auth write
		if err = c.packets.Write(data); err != nil {
			return err
		}
	}

	{
		// read
		if data, err = c.packets.Next(); err != nil {
			return err
		}

		if err = c.handleErrorPacket(data); err != nil {
			return err
		}
	}
	return nil
}

// NewConn used to create a new client connection.
// The timeout is 30 seconds.
func NewConn(username, password, address, database string) (*conn, error) {
	var err error
	c := &conn{}
	timeout := time.Duration(30) * time.Second
	if c.netConn, err = net.DialTimeout("tcp", address, timeout); err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			c.Cleanup()
		}
	}()
	// Set timeouts, make the handshake timeout if the underflying connection blocked.
	// This timeout only used in handshake, we will disable(set zero time) it at last.
	c.netConn.SetReadDeadline(time.Now().Add(timeout))
	defer c.netConn.SetReadDeadline(time.Time{})

	c.auth = proto.NewAuth()
	c.greeting = proto.NewGreeting(0)
	c.packets = packet.NewPackets(c.netConn)
	if err = c.handShake(username, password, database); err != nil {
		return nil, err
	}
	return c, nil
}

func (c *conn) query(command byte, sql string) (Rows, error) {
	var ok *proto.OK
	var myerr, err error
	var columns []*querypb.Field
	var colNumber int

	// if err != nil means the connection is broken(packet error)
	defer func() {
		if err != nil {
			c.Cleanup()
		}
	}()

	// Query.
	if err = c.packets.WriteCommand(command, common.StringToBytes(sql)); err != nil {
		return nil, err
	}

	// Read column number.
	ok, colNumber, myerr, err = c.packets.ReadComQueryResponse()
	if err != nil {
		return nil, err
	}
	if myerr != nil {
		return nil, myerr
	}

	if colNumber > 0 {
		if columns, err = c.packets.ReadColumns(colNumber); err != nil {
			return nil, err
		}

		// Read EOF.
		if (c.greeting.Capability & sqldb.CLIENT_DEPRECATE_EOF) == 0 {
			if err = c.packets.ReadEOF(); err != nil {
				return nil, err
			}
		}
	}
	rows := NewTextRows(c)
	rows.rowsAffected = ok.AffectedRows
	rows.insertID = ok.LastInsertID
	rows.fields = columns
	return rows, nil
}

// ConnectionID is the connection id at greeting
func (c *conn) ConnectionID() uint32 {
	return c.greeting.ConnectionID
}

// Query execute the query and return the row iterator
func (c *conn) Query(sql string) (Rows, error) {
	return c.query(sqldb.COM_QUERY, sql)
}

func (c *conn) Ping() error {
	rows, err := c.query(sqldb.COM_PING, "")
	if err != nil {
		return err
	}

	if err := rows.Close(); err != nil {
		return err
	}
	return nil
}

func (c *conn) InitDB(db string) error {
	rows, err := c.query(sqldb.COM_INIT_DB, db)
	if err != nil {
		return err
	}

	if err := rows.Close(); err != nil {
		return err
	}

	return nil
}

// Exec executes the query and drain the results
func (c *conn) Exec(sql string) error {
	rows, err := c.query(sqldb.COM_QUERY, sql)
	if err != nil {
		return err
	}

	if err := rows.Close(); err != nil {
		c.Cleanup()
	}
	return nil
}

func (c *conn) FetchAll(sql string, maxrows int) (*sqltypes.Result, error) {
	var r *sqltypes.Result

	rows, err := c.query(sqldb.COM_QUERY, sql)
	if err != nil {
		return nil, err
	}

	r = &sqltypes.Result{
		Fields:       rows.Fields(),
		RowsAffected: rows.RowsAffected(),
		InsertID:     rows.LastInsertID(),
	}

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
	if len(r.Rows) > 0 {
		r.RowsAffected = uint64(len(r.Rows))
	}

	// Check last error
	if err := rows.Close(); err != nil {
		c.Cleanup()
		return nil, err
	}
	return r, nil
}

// NextPacket used to get the next packet
func (c *conn) NextPacket() ([]byte, error) {
	return c.packets.Next()
}

func (c *conn) Command(command byte) error {
	rows, err := c.query(command, "")
	if err != nil {
		return err
	}

	if err := rows.Close(); err != nil {
		c.Cleanup()
	}
	return nil
}

func (c *conn) Quit() {
	c.packets.WriteCommand(sqldb.COM_QUIT, nil)
}

func (c *conn) Cleanup() {
	if c.netConn != nil {
		c.Quit()
		c.netConn.Close()
		c.netConn = nil
	}
}

// Close closes the connection
func (c *conn) Close() error {
	if c != nil && c.netConn != nil {
		c.Cleanup()
	}
	return nil
}

// Closed checks the connection broken or not
func (c *conn) Closed() bool {
	return c.netConn == nil
}
