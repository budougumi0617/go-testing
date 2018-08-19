package helper_test

import (
	"testing"
)

// Sum returns sum int values.
func Sum(a, b int) int {
	return a + b
}

func TestSum(t *testing.T) {
	a, b, want := 1, 2, 3
	if got := Sum(a, b); got != want {
		errorf(t, want, got)
		errorfHelper(t, want, got)
	}
}
