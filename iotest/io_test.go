package main

import (
	"bytes"
	"io"
	"testing"
	"testing/iotest"
)

func TestDataErrReader(t *testing.T) {
	orign := []byte("Hello\nbyte.Reader\n")
	type want struct {
		n   int
		buf string
	}
	tests := []struct {
		subject string
		r       io.Reader
		size    int
		wants   []want
	}{
		{
			"DataErrReader",
			iotest.DataErrReader(bytes.NewReader(orign)),
			5,
			[]want{
				{5, "Hello"},
				{5, "\nbyte"},
				{5, ".Read"},
				{3, "er\n\x00\x00"}, // return with io.EOF
			},
		},
		{
			"Normal Reader",
			bytes.NewReader(orign),
			5,
			[]want{
				{5, "Hello"},
				{5, "\nbyte"},
				{5, ".Read"},
				{3, "er\n\x00\x00"},
				{0, "\x00\x00\x00\x00\x00"}, // return with io.EOF
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.subject, func(t *testing.T) {
			for i, want := range tt.wants {
				buf := make([]byte, tt.size)
				size, err := tt.r.Read(buf)
				if size != want.n {
					t.Fatalf("want %d, but got = %d\n", want.n, size)
				}
				if string(buf) != want.buf {
					t.Fatalf("want %#v, but got = %#v\n", want.buf, string(buf))
				}
				if i == len(tt.wants)-1 && err != io.EOF {
					t.Fatalf("want io.EOF, but got = %#v\n", err)
				}
			}
		})
	}
}
