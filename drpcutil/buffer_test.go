// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package drpcutil

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/zeebo/assert"
	"storj.io/drpc/drpctest"
	"storj.io/drpc/drpcwire"
)

func TestBuffer(t *testing.T) {
	run := func(size int) func(t *testing.T) {
		return func(t *testing.T) {
			var exp []byte
			var got bytes.Buffer

			buffer := NewBuffer(&got, size)
			for i := 0; i < 1000; i++ {
				fr := drpctest.RandFrame()
				exp = drpcwire.AppendFrame(exp, fr)
				assert.NoError(t, buffer.Write(fr))
			}
			assert.NoError(t, buffer.Flush())
			assert.DeepEqual(t, exp, got.Bytes())

			// just ensures that the calls did not grow any internal buffers
			assert.Equal(t, cap(buffer.buf), size)
			assert.Equal(t, cap(buffer.tmp), drpcwire.MaxPacketSize)
		}
	}

	t.Run("0", run(0))
	t.Run(fmt.Sprint(drpcwire.MaxPacketSize), run(drpcwire.MaxPacketSize))
	t.Run("1MB", run(1024*1024))
}
