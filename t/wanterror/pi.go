package wanterror

import "errors"

// PositiveInt is too simple...
type PositiveInt int

// Value returns posive value.
func (pi *PositiveInt) Value() (int, error) {
	if int(*pi) < 0 {
		return -1, errors.New("Negative value")
	}
	return int(*pi), nil
}
