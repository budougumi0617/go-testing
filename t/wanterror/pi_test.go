package wanterror_test

import (
	"errors"
	"reflect"
	"testing"

	we "github.com/budougumi0617/go-testing/t/wanterror"
)

func TestPositiveInt(t *testing.T) {
	tests := []struct {
		name      string
		in        int
		want      int
		wantError bool
		err       error
	}{
		{"Basic", 4, 4, false, nil},
		{"HasError", -1, 0, true, errors.New("Negative value")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pt := we.PositiveInt(tt.in)
			got, err := pt.Value()
			if !tt.wantError && err != nil {
				t.Fatalf("want no err, but has error %#v", err)
			}

			if tt.wantError && !reflect.DeepEqual(err, tt.err) {
				t.Fatalf("want %#v, but %#v", tt.err, err)
			}

			if !tt.wantError && got != tt.want {
				t.Fatalf("want %q, but %q", tt.want, got)
			}
		})
	}
}
