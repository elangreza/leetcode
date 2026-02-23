package a

import "fmt"

// s = "   fly me   to   the moon  "

func lengthOfLastWord(s string) int {

	// iterate from the back
	x := []byte{}
	for i := len(s) - 1; i > -1; i-- {
		fmt.Println(string(s[i]))
		if s[i] != ' ' {
			x = append(x, s[i])
		}
		if s[i] == ' ' && len(x) > 0 {
			break
		}
	}

	return len(x)
}
