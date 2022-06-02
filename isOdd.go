package isOdd

import "golang.org/x/exp/constraints"

func IsOdd[T constraints.Integer](val T) bool {
	return val%2 == 1
}
