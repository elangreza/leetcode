package zigzagconversion

// Example 1:

// Input: s = "PAYPALISHIRING", numRows = 3
// Output: "PAHNAPLSIIGYIR"
// Explanation:
// P   A   H   N
// A P L S I I G
// Y   I   R

// 1 2 3 4 5 6 7 8 9

// Input: s = "PAYPALISHIRING", numRows = 4
// Output: "PINALSIGYAHRPI"
// Explanation:
// P     I    N
// A   L S  I G
// Y A   H R
// P     I

func convert(s string, numRows int) string {
	if len(s) < 3 || numRows == 1 {
		return s
	}

	arrs := make([][]byte, numRows)
	direction := 1

	pos := 0
	for i := 0; i < len(s); i++ {
		arrs[pos] = append(arrs[pos], s[i])

		if pos == 0 {
			direction = 1
		}

		if pos == (numRows - 1) {
			direction = -1
		}

		pos = pos + direction
	}

	res := ""
	for _, arr := range arrs {
		res += string(arr)
	}

	return res
}
