package longestpalindromicsubstring

// using
func longestPalindromeV0(ss string) string {
	res := ""
	lenRes := 0

	for start, _ := range ss {
		for end, _ := range ss[start:] {
			e := max(end, start)
			word := ss[start : e+1]
			ok := palindrome(word)
			// fmt.Println(word, ok, start, e)
			if ok && len(word) > lenRes {
				lenRes = len(word)
				res = word
			}
		}
	}

	return res
}

func longestPalindromeV3(ss string) string {
	res := ""

	// check with expanding from middle
	f := func(i int, j int) {
		for i >= 0 && j < len(ss) && ss[i] == ss[j] {
			if len(ss[i:j+1]) > len(res) {
				res = ss[i : j+1]
			}
			i--
			j++
		}
	}

	for index := range ss {

		f(index, index)

		// it needs to be check when ss length is even
		f(index, index+1)
	}

	return res
}

func palindrome(s string) bool {
	times := len(s) / 2
	for i := range times {
		if s[i] == s[len(s)-1-i] {
			continue
		}

		return false
	}

	return true
}
