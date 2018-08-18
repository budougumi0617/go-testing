// +build integration

package tags_test

import (
	"testing"

	"github.com/budougumi0617/go-testing/t/tags"
)

func TestSub(t *testing.T) {
	a, b, want := 1, 2, -1
	if got := tags.Sub(a, b); got != want {
		t.Fatalf("want = %d, got = %d", want, got)
	}
}
