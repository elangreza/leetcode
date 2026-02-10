package a

func lengthOfLongestSubstring(s string) int {
	maxLength := 0
	start := 0
	// map is created to track the position of each char
	// key is the char, and value is the index
	m := make(map[rune]int)

	for end, char := range s {
		// 3. we want to check if repeated character is occurred
		// we want to check char is available in the map
		prevIndex, ok := m[char]
		// 4. if the char is found
		// we update the start position
		if ok {
			start = max(start, prevIndex+1)
		}
		// 1. calculate the maxLength condition
		// this covered first iteration, since start and end is in the same position
		// we want to get exactly one iteration
		// let's say input is " " the result will always be 1
		maxLength = max(maxLength, end-start+1)
		// 2. second we want to save char in the map
		// save char as key
		// and save end value or index as a value
		m[char] = end
	}

	return maxLength
}

// func lengthOfLongestSubstring(s string) int {
// 	maxLength := 0
// 	m := make(map[rune]int)
// 	start := 0
// 	for end, char := range s {
// 		prevIndex, ok := m[char]
// 		if ok {
// 			start = max(start, prevIndex+1)
// 		}
// 		m[char] = end
// 		// condition for covering exactly one char
// 		// if end is not moved yet, but the length can be calculated as one
// 		maxLength = max(maxLength, end-start+1)
// 	}

// 	return maxLength
// }

// abba
// set all
