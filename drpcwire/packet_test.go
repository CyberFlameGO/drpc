// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package drpcwire_test

import (
	"testing"

	"github.com/zeebo/assert"

	"storj.io/drpc/drpctest"
	"storj.io/drpc/drpcwire"
)

func TestAppendParse(t *testing.T) {
	requireGoodParse := func(t *testing.T, exp interface{}) func([]byte, interface{}, bool, error) {
		return func(rem []byte, got interface{}, ok bool, err error) {
			t.Helper()
			assert.NoError(t, err)
			assert.That(t, ok)
			assert.Equal(t, 0, len(rem))
			assert.DeepEqual(t, exp, got)
		}
	}

	t.Run("PacketID_RoundTrip_Fuzz", func(t *testing.T) {
		for i := 0; i < 10000; i++ {
			exp := drpctest.RandPacketID()
			requireGoodParse(t, exp)(drpcwire.ParsePacketID(drpcwire.AppendPacketID(nil, exp)))
		}
	})

	t.Run("FrameInfo_RoundTrip_Fuzz", func(t *testing.T) {
		for i := 0; i < 10000; i++ {
			exp := drpctest.RandFrameInfo()
			requireGoodParse(t, exp)(drpcwire.ParseFrameInfo(drpcwire.AppendFrameInfo(nil, exp)))
		}
	})

	t.Run("Header_RoundTrip_Fuzz", func(t *testing.T) {
		for i := 0; i < 10000; i++ {
			exp := drpctest.RandHeader()
			requireGoodParse(t, exp)(drpcwire.ParseHeader(drpcwire.AppendHeader(nil, exp)))
		}
	})

	t.Run("Frame_RoundTrip_Fuzz", func(t *testing.T) {
		for i := 0; i < 1000; i++ {
			exp := drpctest.RandFrame()
			requireGoodParse(t, exp)(drpcwire.ParseFrame(drpcwire.AppendFrame(nil, exp)))
		}
	})
}
