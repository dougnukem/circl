// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !amd64,!386,!ppc64le

package shake

import "encoding/binary"

type storageBuf [maxRate]byte

func (b *storageBuf) asBytes() *[maxRate]byte {
    return (*[maxRate]byte)(b)
}

// xorInGeneric xors the bytes in buf into the state; it
// makes no non-portable assumptions about memory layout
// or alignment.
func xorIn(d *Shake, buf []byte) {
	n := len(buf) / 8

	for i := 0; i < n; i++ {
		a := binary.LittleEndian.Uint64(buf)
		d.a[i] ^= a
		buf = buf[8:]
	}
}

// copyOutGeneric copies ulint64s to a byte buffer.
func copyOut(d *Shake, b []byte) {
	for i := 0; len(b) >= 8; i++ {
		binary.LittleEndian.PutUint64(b, d.a[i])
		b = b[8:]
	}
}
