package plaindromenumber

func isPalindromeV2(x int) bool {
	if x < 0 {
		return false
	}

	divider := 1
	for (divider * 10) <= x {
		divider = divider * 10
	}

	for divider > 0 {
		if x/divider != x%10 {
			return false
		}
		x = x % divider
		x = x / 10
		divider = divider / 100
	}

	return true
}
