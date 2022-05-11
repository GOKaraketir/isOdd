package isOdd

import "golang.org/x/exp/constraints"

func abs[T constraints.Integer](val T) T {
	if val < 0 {
		return -val
	} else {
		return val
	}
}

func IsOdd[T constraints.Integer](val T) bool {
	return abs(val)%2 == 1
}
