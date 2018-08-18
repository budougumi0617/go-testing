package table_test

import (
	"testing"

	"github.com/budougumi0617/go-testing/t/parallel"
)

func TestSum(t *testing.T) {
	tests := []struct {
		name       string
		a, b, want int
	}{
		{"Simple", 0, 1, 1},
		{"Minus", -1, -1, -2},
		{"Both", -3, 2, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parallel.Sum(tt.a, tt.b); got != tt.want {
				t.Fatalf("want = %d, got = %d", tt.want, got)
			}
		})
	}
}
