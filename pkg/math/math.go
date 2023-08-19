package math

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Integer | constraints.Float
}

func Max[T Number](x, y T) T {
	if x > y {
		return x
	}

	return y
}

func Min[T Number](x, y T) T {
	if (x < y) {
		return x
	}

	return y
}

