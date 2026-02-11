package longestcommonprefix

func longestCommonPrefixV0(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}

	maxIter := len(strs[0])
	for _, v := range strs[1:] {
		maxIter = min(len(v), maxIter)
	}

	max := 0
	for index := range maxIter {
		var a byte = strs[0][index]
		for _, str := range strs[1:] {
			b := str[index]
			if a != b {
				return strs[0][:index]
			}
		}
		max++
	}

	return strs[0][:max]
}
