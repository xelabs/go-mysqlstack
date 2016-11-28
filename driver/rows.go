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
	"github.com/XeLabs/go-mysqlstack/common"
	"github.com/XeLabs/go-mysqlstack/packet"
	"github.com/XeLabs/go-mysqlstack/proto"
	"github.com/pkg/errors"
	"strconv"
)

type Rows struct {
	buffer  *common.Buffer
	packets *packet.Packets
	payload []byte
	columns []proto.Column
	err     error
}

func NewRows(packets *packet.Packets) *Rows {
	return &Rows{
		packets: packets,
		buffer:  common.NewBuffer(8),
	}
}

func (r *Rows) String() (s string) {
	s, r.err = r.buffer.ReadLenEncodeString()
	return
}

func (r *Rows) Int() (v int) {
	v, r.err = strconv.Atoi(r.String())
	return
}

// http://dev.mysql.com/doc/internals/en/com-query-response.html#packet-ProtocolText::ResultsetRow
func (r *Rows) Next() bool {
	if r.payload, r.err = r.packets.Next(); r.err != nil {
		return false
	}

	switch r.payload[0] {
	case proto.EOF_PACKET:
		return false

	case proto.ERR_PACKET:
		r.err = errors.New("next.packet.error")
		return false
	}

	r.buffer.Reset(r.payload)

	return true
}

func (r *Rows) LastError() error {
	return r.err
}
