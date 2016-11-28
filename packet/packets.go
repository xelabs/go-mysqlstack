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
	"github.com/XeLabs/go-mysqlstack/common"
	"github.com/XeLabs/go-mysqlstack/proto"
	"github.com/pkg/errors"
	"io"
)

type Packet struct {
	SequenceID byte
	Payload    []byte
}

type Packets struct {
	seq    uint8
	stream *Stream
}

func NewPackets(rw io.ReadWriter) *Packets {
	return &Packets{
		stream: NewStream(rw),
	}
}

// Read reads packet from stream
func (p *Packets) Next() (v []byte, e error) {
	pkt, err := p.stream.Read()
	if err != nil {
		e = err
		return
	}

	if pkt.SequenceID != p.seq {
		e = errors.Errorf("pkt.read.seq[%v]!=pkt.actual.seq[%v]", pkt.SequenceID, p.seq)
		return
	}
	p.seq++

	return pkt.Payload, nil
}

// Write writes the packet to stream
// It packs as:
// [header]
// [payload]
func (p *Packets) Write(payload []byte) error {
	payLen := len(payload)
	pkt := common.NewBuffer(128)

	// body length(24bits)
	pkt.WriteU24(uint32(payLen))

	// SequenceID
	pkt.WriteU8(p.seq)

	// body
	pkt.WriteBytes(payload, payLen)
	if err := p.stream.Write(pkt.Datas()); err != nil {
		return err
	}
	p.seq++

	return nil
}

func (p *Packets) WriteCommand(command byte, payload []byte) error {
	// reset packet sequence
	p.seq = 0
	pkt := common.NewBuffer(128)

	// body length(24bits):
	// command length + payload length
	payLen := len(payload)
	pkt.WriteU24(uint32(1 + payLen))

	// SequenceID
	pkt.WriteU8(p.seq)

	// command
	pkt.WriteU8(command)

	// body
	pkt.WriteBytes(payload, payLen)
	if err := p.stream.Write(pkt.Datas()); err != nil {
		return err
	}
	p.seq++

	return nil
}

// ResetSeq reset sequence to zero
func (p *Packets) ResetSeq() {
	p.seq = 0
}

func (p *Packets) ParseERR(data []byte, capabilityFlags uint32) (e *proto.ERR, err error) {
	e = proto.NewERR()
	err = e.UnPack(data, capabilityFlags)

	return
}

func (p *Packets) ParseOK(data []byte, capabilityFlags uint32) (ok *proto.OK, err error) {
	ok = proto.NewOK()
	err = ok.UnPack(data, capabilityFlags)

	return
}
