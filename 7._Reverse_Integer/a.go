package a

import "math"

func reverse(x int) int {
	isMin := false
	if x < 0 {
		isMin = true
		x = -1 * x
	}

	divider := 1
	for (divider * 10) <= x {
		divider = divider * 10
	}

	res := 0
	for divider > 0 {
		buff := x % 10
		x = x / 10
		res = res + (buff * divider)
		divider = divider / 10
	}

	if res > math.MaxUint32 {
		return 0
	}

	if isMin {
		return -res
	}

	return res
}
