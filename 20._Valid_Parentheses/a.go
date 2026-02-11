package validparentheses

func isValid(s string) bool {
	st := []rune{}
	for _, v := range s {
		switch v {
		case '(', '{', '[':
			st = append(st, v)
		case '}', ')', ']':
			switch v {
			case ']':
				if st[len(st)-1] != '[' {
					return false
				}
			case ')':
				if st[len(st)-1] != '(' {
					return false
				}
			case '}':
				if st[len(st)-1] != '{' {
					return false
				}
			}
			st = st[:len(st)-1]
		}
	}

	return true
}
