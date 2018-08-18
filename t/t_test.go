package t

import (
	"testing"
)

// Sum returns two int values.
func Sum(a, b int) int {
	return a + b
}

func TestSum(t *testing.T) {
	a, b, want := 1, 2, 3
	if got := Sum(a, b); got != want {
		t.Fatalf("want = %d, got = %d", want, got)
	}
}
