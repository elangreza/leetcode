package a

func lengthOfLongestSubstring2(s string) int {
	slow, fast := 0, 0
	res := 0
	m := make(map[byte]int)
	for fast < len(s) {
		v, ok := m[s[fast]]
		if ok {
			slow = max(slow, v)
			// m[s[fast]] = fast + 1
			// fast++
			// continue
		}
		m[s[fast]] = fast + 1
		res = max(res, fast-slow+1)
		fast++
	}

	return res
}
