package validparentheses

func isValid(s string) bool {

	if len(s) <= 1 {
		return false
	}
	x := []rune{}
	for _, v := range s {
		switch v {
		case '(', '[', '{':
			x = append(x, v)
			continue
		case ')':
			if len(x) > 0 && x[len(x)-1] != '(' {
				return false
			}
		case ']':
			if len(x) > 0 && x[len(x)-1] != '[' {
				return false
			}
		case '}':
			if len(x) > 0 && x[len(x)-1] != '{' {
				return false
			}
		}
		if len(x) > 0 {
			x = x[:len(x)-1]
		}
	}

	return len(x) <= 1
}

func isValidV2(str string) bool {

	stack := make([]rune, 0)

	ms := make(map[rune]rune)
	ms['{'] = '}'
	ms['['] = ']'
	ms['('] = ')'

	for _, s := range str {
		if _, ok := ms[s]; ok {
			stack = append(stack, s)
		} else {
			if len(stack) == 0 {
				return false
			}

			// get last
			last := stack[len(stack)-1]
			// cut last of the stack
			stack = stack[:len(stack)-1]

			if ms[last] != s {
				return false
			}
		}
	}

	return len(stack) == 0
}

func isValidV3(s string) bool {
	stack := []rune{} // Using a slice as a stack

	// Map to store corresponding closing brackets
	bracketMap := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	for _, char := range s {
		// If it's an opening bracket, push it onto the stack
		if _, ok := bracketMap[char]; ok {
			stack = append(stack, char)
		} else { // It's a closing bracket
			// If the stack is empty, no matching opening bracket exists
			if len(stack) == 0 {
				return false
			}

			// Pop the last opening bracket from the stack
			lastOpen := stack[len(stack)-1]
			stack = stack[:len(stack)-1] // Pop operation

			// Check if the popped opening bracket matches the current closing bracket
			if bracketMap[lastOpen] != char {
				return false
			}
		}
	}

	// After iterating, if the stack is empty, all brackets are matched
	return len(stack) == 0
}
