package romantointeger

func romanToIntV0(s string) int {
	res := 0
	length := len(s) - 1
	for i, j := length, length; i >= 0; i-- {
		j = i - 1
		switch string(s[i]) {
		case "I":
			res += 1
		case "V":
			if j > -1 && string(s[j]) == "I" {
				res += 4
				i--
				continue
			}
			res += 5
		case "X":
			if j > -1 && string(s[j]) == "I" {
				res += 9
				i--
				continue
			}
			res += 10
		case "L":
			if j > -1 && string(s[j]) == "X" {
				res += 40
				i--
				continue
			}
			res += 50
		case "C":
			if j > -1 && string(s[j]) == "X" {
				res += 90
				i--
				continue
			}
			res += 100
		case "D":
			if j > -1 && string(s[j]) == "C" {
				res += 40
				i--
				continue
			}
			res += 500
		case "M":
			if j > -1 && string(s[j]) == "C" {
				res += 900
				i--
				continue
			}
			res += 1000
		}
	}

	return res
}

func romanToIntV1(s string) int {

	m := map[string]int{
		"I":   1,
		"II":  2,
		"III": 3,
		"IV":  4,
		"V":   5,
		"IX":  9,
		"X":   10,
		"XX":  20,
		"XXX": 30,
		"XL":  40,
		"L":   50,
		"XC":  90,
		"C":   100,
		"CC":  200,
		"CCC": 300,
		"CD":  400,
		"D":   500,
		"CM":  900,
		"M":   1000,
		"MM":  2000,
		"MMM": 3000,
	}

	res := 0
	i := len(s) - 1
	j := len(s) - 1
	lastVal := 0
	for j > -1 {
		val, ok := m[string(s[j:i+1])]
		if ok {
			lastVal = val
			j--
			continue
		} else {
			i = j
			res = res + lastVal
			lastVal = 0
		}
	}

	if lastVal > 0 {
		return res + lastVal
	}

	return res
}

func romanToInt(s string) int {

	m := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	res := 0
	i := len(s) - 1
	lastVal := 0
	for i > -1 {
		val, ok := m[s[i]]
		if ok {
			if val >= lastVal {
				res = res + val
			} else {
				res = res - val
			}
			lastVal = val
		}
		i--
	}

	return res
}
