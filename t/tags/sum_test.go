// +build !integration

package tags_test

import (
	"testing"

	"github.com/budougumi0617/go-testing/t/tags"
)

func TestSum(t *testing.T) {
	a, b, want := 1, 2, 3
	if got := tags.Sum(a, b); got != want {
		t.Fatalf("want = %d, got = %d", want, got)
	}
}
