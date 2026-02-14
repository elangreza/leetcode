package removeduplicatesfromsortedarray

import "fmt"

func removeDuplicatesV0(nums []int) int {
	if len(nums) == 1 {
		return 1
	}

	j := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			fmt.Println(j, i, nums[j], nums[i])
			nums[j] = nums[i]
			j++
		}
	}

	fmt.Println(nums)

	return j
}
