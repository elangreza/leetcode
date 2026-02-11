package longestcommonprefix

func longestCommonPrefix(strs []string) string {

	minLength := 200
	for _, str := range strs {
		minLength = min(minLength, len(str))
	}
	var res []byte

	for v := range minLength {
		var char byte
		for i := range len(strs) {
			if char == 0 {
				char = strs[i][v]
			} else if char != strs[i][v] {
				return string(res)
			}
		}
		res = append(res, char)
	}

	return string(res)
}
