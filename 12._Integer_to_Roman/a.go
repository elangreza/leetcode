package a

func intToRoman(num int) string {
	// divider := 1
	// for num >= (divider * 10) {
	// 	divider = divider * 10
	// }

	table := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	tableName := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	// 3749 > 1000
	//
	res := ""
	for i, t := range table {
		for num >= t {
			res = res + tableName[i]
			num = num - t
		}
	}

	// for divider > 0 {
	// 	left := (num / divider) * divider
	// 	for i, t := range table {
	// 		if left > t {
	// 			fmt.Println(left, (left), t, i)
	// 			break
	// 		}
	// 	}

	// 	// calculate for next iteration
	// 	num = num - (left)
	// 	divider = divider / 10
	// }

	return res
}

// 40
// 10 10 10 10
//
