// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package drpcwire

import (
	"bufio"
	"io"

	"storj.io/drpc"
)

func FrameScanner(data []byte, atEOF bool) (int, []byte, error) {
	rem, _, ok, err := ParseFrame(data)
	switch advance := len(data) - len(rem); {
	case err != nil, !ok:
		return 0, nil, err
	case advance < 0, len(data) < advance:
		return 0, nil, drpc.InternalError.New("scanner issue with advance value")
	default:
		return advance, data[:advance], nil
	}
}

func NewScanner(r io.Reader) *bufio.Scanner {
	scanner := bufio.NewScanner(r)
	scanner.Buffer(make([]byte, 4*1024), MaxPacketSize)
	scanner.Split(FrameScanner)
	return scanner
}
