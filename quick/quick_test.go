package main

import (
	"math/rand"
	"reflect"
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

	// Set config
	cfg := &quick.Config{
		// 試行回数
		MaxCount: 10,
		// Seed
		Rand: rand.New(rand.NewSource(2)),
		// 独自定義した引数の生成関数
		Values: func(args []reflect.Value, rand *rand.Rand) {
			args[0] = reflect.ValueOf(MyInt{rand.Int(), "x"})
			args[1] = reflect.ValueOf(MyInt{rand.Int(), "y"})
		},
	}

	if err := quick.Check(f, cfg); err != nil {
		// Use specified error type
		if ce, ok := err.(*quick.CheckError); ok {
			t.Errorf("Try count = %d, In %#v, Out %s\n", ce.Count, ce.In, ce)
		} else {
			t.Error(err)
		}
	}
}
