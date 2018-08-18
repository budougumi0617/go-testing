package helper_test

import (
	"testing"
)

func errorf(tb testing.TB, want, got int) {
	tb.Errorf("want = %d, got = %d", want, got)
}

func errorfHelper(tb testing.TB, want, got int) {
	tb.Helper()
	tb.Errorf("want = %d, got = %d", want, got)
}
