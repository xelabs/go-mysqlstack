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
	"github.com/stretchr/testify/assert"
	"testing"
)

// TEST EFFECTS:
// writes normal packet
//
// TEST PROCESSES:
// 1. write datas more than PACKET_BUFFER_SIZE
// 2. write checks
// 3. read checks
func TestStream(t *testing.T) {
	rBuf := NewMockConn()
	wBuf := NewMockConn()

	rStream := NewStream(rBuf)
	wStream := NewStream(wBuf)

	packet := common.NewBuffer(PACKET_BUFFER_SIZE)
	payload := common.NewBuffer(PACKET_BUFFER_SIZE)

	for i := 0; i < 1234; i++ {
		payload.WriteU8(byte(i))
	}

	packet.WriteU24(uint32(payload.Length()))
	packet.WriteU8(1)
	packet.WriteBytes(payload.Datas(), payload.Length())

	// write checks
	{
		err := wStream.Write(packet.Datas())
		assert.Nil(t, err)

		want := packet.Datas()
		got := wBuf.Datas()
		assert.Equal(t, want, got)
	}

	// read checks
	{
		rBuf.Write(wBuf.Datas())
		ptk, err := rStream.Read()
		assert.Nil(t, err)

		assert.Equal(t, ptk.SequenceID, byte(0x01))
		assert.Equal(t, ptk.Payload, payload.Datas())
	}
}

// TEST EFFECTS:
// write packet whoes payload length equals PACKET_MAX_SIZE
//
// TEST PROCESSES:
// 1. write payload whoes length equals PACKET_MAX_SIZE
// 2. read checks
// 3. write checks
func TestStreamWriteMax(t *testing.T) {
	rBuf := NewMockConn()
	wBuf := NewMockConn()

	rStream := NewStream(rBuf)
	wStream := NewStream(wBuf)

	packet := common.NewBuffer(PACKET_BUFFER_SIZE)
	expect := common.NewBuffer(PACKET_BUFFER_SIZE)
	payload := common.NewBuffer(PACKET_BUFFER_SIZE)

	// fill payload with PACKET_MAX_SIZE(16777215)
	{
		for i := 0; i < 4194303; i++ {
			payload.WriteU32(uint32(i))
		}
		payload.WriteU24(24)
	}
	packet.WriteU24(uint32(payload.Length()))
	packet.WriteU8(1)
	packet.WriteBytes(payload.Datas(), payload.Length())

	// write checks
	{
		err := wStream.Write(packet.Datas())
		assert.Nil(t, err)

		// check length
		{
			want := packet.Length() + 4
			got := len(wBuf.Datas())
			assert.Equal(t, want, got)
		}

		// check chunks
		{
			// first chunk
			expect.WriteU24(uint32(PACKET_MAX_SIZE))
			expect.WriteU8(1)
			expect.WriteBytes(payload.Datas(), PACKET_MAX_SIZE)

			// second chunk
			expect.WriteU24(0)
			expect.WriteU8(2)

			want := expect.Datas()
			got := wBuf.Datas()
			assert.Equal(t, want, got)
		}
	}

	// read checks
	{
		rBuf.Write(wBuf.Datas())
		ptk, err := rStream.Read()
		assert.Nil(t, err)

		assert.Equal(t, ptk.SequenceID, byte(0x02))
		assert.Equal(t, ptk.Payload, payload.Datas())
	}
}

// TEST EFFECTS:
// write packet whoes payload length more than PACKET_MAX_SIZE
//
// TEST PROCESSES:
// 1. write payload whoes length (PACKET_MAX_SIZE + 8)
// 2. read checks
// 3. write checks
func TestStreamWriteOverMax(t *testing.T) {
	rBuf := NewMockConn()
	wBuf := NewMockConn()

	rStream := NewStream(rBuf)
	wStream := NewStream(wBuf)

	packet := common.NewBuffer(PACKET_BUFFER_SIZE)
	expect := common.NewBuffer(PACKET_BUFFER_SIZE)
	payload := common.NewBuffer(PACKET_BUFFER_SIZE)

	// fill payload with PACKET_MAX_SIZE(16777215)
	{
		for i := 0; i < 4194303; i++ {
			payload.WriteU32(uint32(i))
		}
		payload.WriteU24(24)
	}
	// fill with 8bytes
	payload.WriteU32(32)
	payload.WriteU32(32)

	packet.WriteU24(uint32(payload.Length()))
	packet.WriteU8(1)
	packet.WriteBytes(payload.Datas(), payload.Length())

	// write checks
	{
		err := wStream.Write(packet.Datas())
		assert.Nil(t, err)

		// check length
		{
			want := packet.Length() + 4
			got := len(wBuf.Datas())
			assert.Equal(t, want, got)
		}

		// check chunks
		{
			// first chunk
			expect.WriteU24(uint32(PACKET_MAX_SIZE))
			expect.WriteU8(1)
			expect.WriteBytes(payload.Datas(), PACKET_MAX_SIZE)

			// second chunk
			left := (packet.Length() - 4) - PACKET_MAX_SIZE
			expect.WriteU24(uint32(left))
			expect.WriteU8(2)
			expect.WriteBytes(payload.Datas()[PACKET_MAX_SIZE:], left)

			want := expect.Datas()
			got := wBuf.Datas()
			assert.Equal(t, want, got)
		}
	}

	// read checks
	{
		rBuf.Write(wBuf.Datas())
		ptk, err := rStream.Read()
		assert.Nil(t, err)

		assert.Equal(t, ptk.SequenceID, byte(0x02))
		assert.Equal(t, ptk.Payload, payload.Datas())
		_, err = rStream.Read()
		assert.NotNil(t, err)
	}
}
