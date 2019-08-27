// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package drpctest

import (
	"math"
	"math/rand"

	"storj.io/drpc/drpcwire"
)

func RandBytes(n int) []byte {
	out := make([]byte, n)
	for i := range out {
		out[i] = byte(rand.Intn(256))
	}
	return out
}

func RandUint64() uint64 {
	return uint64(rand.Int63n(math.MaxInt64))<<1 + uint64(rand.Intn(2))
}

func RandBool() bool {
	return rand.Intn(2) == 0
}

func RandPacketID() drpcwire.PacketID {
	return drpcwire.PacketID{
		StreamID:  RandUint64() | 1,
		MessageID: RandUint64() | 1,
	}
}

func RandPayloadKind() drpcwire.PayloadKind {
	return drpcwire.PayloadKind(rand.Intn(int(drpcwire.PayloadKind_Largest)-1) + 1)
}

// payloadMaxSize maps a payload kind to 1 more than the maximum number of bytes that can
// be sent with a packet with that kind.
var payloadMaxSize = map[drpcwire.PayloadKind]func() int{
	drpcwire.PayloadKind_Invoke:    func() int { return rand.Intn(1023) + 1 },
	drpcwire.PayloadKind_Message:   func() int { return rand.Intn(1023) + 1 },
	drpcwire.PayloadKind_Error:     func() int { return rand.Intn(1023) + 1 },
	drpcwire.PayloadKind_Cancel:    func() int { return 0 },
	drpcwire.PayloadKind_Close:     func() int { return 0 },
	drpcwire.PayloadKind_CloseSend: func() int { return 0 },
}

var kindCanContinue = map[drpcwire.PayloadKind]bool{
	drpcwire.PayloadKind_Invoke:    true,
	drpcwire.PayloadKind_Message:   true,
	drpcwire.PayloadKind_Error:     true,
	drpcwire.PayloadKind_Cancel:    false,
	drpcwire.PayloadKind_Close:     false,
	drpcwire.PayloadKind_CloseSend: false,
}

func RandFrameInfo() drpcwire.FrameInfo {
	kind := RandPayloadKind()
	return drpcwire.FrameInfo{
		Length:       uint16(payloadMaxSize[kind]()),
		Continuation: kindCanContinue[kind] && RandBool(),
		Starting:     !kindCanContinue[kind] || RandBool(),
		PayloadKind:  kind,
	}
}

func RandHeader() drpcwire.Header {
	return drpcwire.Header{
		FrameInfo: RandFrameInfo(),
		PacketID:  RandPacketID(),
	}
}

func RandFrame() drpcwire.Frame {
	hdr := RandHeader()
	return drpcwire.Frame{
		Header: hdr,
		Data:   RandBytes(int(hdr.Length)),
	}
}

func RandPacket() drpcwire.Packet {
	kind := RandPayloadKind()
	return drpcwire.Packet{
		PacketID:    RandPacketID(),
		PayloadKind: kind,
		Data:        RandBytes(100 * payloadMaxSize[kind]()),
	}
}
