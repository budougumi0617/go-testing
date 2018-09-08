package main

import (
	"bytes"
	"io"
	"testing"
	"testing/iotest"
)

func TestIoTest(t *testing.T) {
	orign := []byte("Hello\nbyte.Reader\n")
	type want struct {
		n         int
		buf       string
		wantError error
	}
	tests := []struct {
		subject string
		r       io.Reader
		size    int
		wants   []want
	}{
		{
			"Normal Reader",
			bytes.NewReader(orign),
			5,
			[]want{
				{5, "Hello", nil},
				{5, "\nbyte", nil},
				{5, ".Read", nil},
				{3, "er\n\x00\x00", nil},
				{0, "\x00\x00\x00\x00\x00", io.EOF}, // return with io.EOF
			},
		},
		{
			"DataErrReader",
			iotest.DataErrReader(bytes.NewReader(orign)),
			5,
			[]want{
				{5, "Hello", nil},
				{5, "\nbyte", nil},
				{5, ".Read", nil},
				{3, "er\n\x00\x00", io.EOF}, // return with io.EOF
			},
		},
		{
			"HalfReader",
			iotest.HalfReader(bytes.NewReader(orign)),
			5,
			[]want{
				// len(5)のbufferでReadしても、半分の3バイトしか読み込んでくれない
				{3, "Hel\x00\x00", nil},
				{3, "lo\n\x00\x00", nil},
				{3, "byt\x00\x00", nil},
				{3, "e.R\x00\x00", nil},
				{3, "ead\x00\x00", nil},
				{3, "er\n\x00\x00", nil},
				{0, "\x00\x00\x00\x00\x00", io.EOF},
			},
		},
		{
			"OneByteReader",
			iotest.OneByteReader(bytes.NewReader(orign)),
			5,
			[]want{
				// 1バイトしか読み込まない
				{1, "H\x00\x00\x00\x00", nil},
				{1, "e\x00\x00\x00\x00", nil},
				{1, "l\x00\x00\x00\x00", nil},
				{1, "l\x00\x00\x00\x00", nil},
				{1, "o\x00\x00\x00\x00", nil},
				{1, "\n\x00\x00\x00\x00", nil},
				{1, "b\x00\x00\x00\x00", nil},
				{1, "y\x00\x00\x00\x00", nil},
				{1, "t\x00\x00\x00\x00", nil},
				{1, "e\x00\x00\x00\x00", nil},
				{1, ".\x00\x00\x00\x00", nil},
				{1, "R\x00\x00\x00\x00", nil},
				{1, "e\x00\x00\x00\x00", nil},
				{1, "a\x00\x00\x00\x00", nil},
				{1, "d\x00\x00\x00\x00", nil},
				{1, "e\x00\x00\x00\x00", nil},
				{1, "r\x00\x00\x00\x00", nil},
				{1, "\n\x00\x00\x00\x00", nil},
				{0, "\x00\x00\x00\x00\x00", io.EOF},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.subject, func(t *testing.T) {
			for _, want := range tt.wants {
				buf := make([]byte, tt.size)
				size, err := tt.r.Read(buf)
				if size != want.n {
					t.Fatalf("want %d, but got = %d\n", want.n, size)
				}
				if string(buf) != want.buf {
					t.Fatalf("want %#v, but got = %#v\n", want.buf, string(buf))
				}
				if err != want.wantError {
					t.Fatalf("want io.EOF, but got = %#v\n", err)
				}
			}
		})
	}
}
