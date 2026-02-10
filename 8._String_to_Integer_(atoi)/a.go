package a

import (
	"math"
	"strings"
	"unicode"
)

func myAtoi(s string) int {

	s = strings.TrimLeft(s, " ")

	isNegative := false

	i := 0
	if s[i] == '-' {
		isNegative = true
		i++
	} else if s[i] == '+' {
		isNegative = false
		i++
	}

	res := 0

	for i < len(s) && unicode.IsDigit(rune(s[i])) {
		a := int(s[i] - '0')
		res = (res * 10) + a
		i++
	}

	if isNegative {
		res = -res
	}

	if res > math.MaxInt32 {
		return math.MaxInt32
	}

	if res < math.MinInt32 {
		return math.MinInt32
	}

	return res
}

// 42
// 4
// 40 + 2
