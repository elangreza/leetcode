package a

func plusOne(digits []int) []int {
	digits[len(digits)-1] = digits[len(digits)-1] + 1

	for i := len(digits) - 1; i > 0; i-- {
		digit := digits[i]
		digits[i] = digit % 10
		digits[i-1] = digits[i-1] + digit/10
	}

	if digits[0] > 9 {
		left := digits[0] / 10
		digits[0] = digits[0] % 10
		digits = append([]int{left}, digits...)
	}

	return digits
}
