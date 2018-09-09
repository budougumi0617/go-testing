package main

import (
	"testing"
	"testing/quick"
)

type MyInt struct {
	V int
	X string
}

func div(n, m MyInt) int {
	return n.V / m.V
}

func TestDiv(t *testing.T) {
	// div関数は戻り値がintなので、評価結果をboolで返すラップ関数を用意する
	f := func(x, y MyInt) bool {
		return div(x, y) == (x.V / y.V)
	}
	if err := quick.Check(f, nil); err != nil {
		if ce, ok := err.(*quick.CheckError); ok {
			t.Errorf("Try count = %d, In %#v, Out %s\n", ce.Count, ce.In, ce)
		} else {
			t.Error(err)
		}
	}
}
