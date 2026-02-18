package findfirstandlastpositionofelementinsortedarray

func searchRange(arr []int, target int) []int {
	res := []int{-1, -1}

	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)/2
		if target == arr[mid] {
			res[0] = mid
			right = mid - 1
		}

		if target < arr[mid] {
			left = mid + 1
		}

		if target > arr[mid] {
			right = mid - 1
		}
	}

	if res[0] == -1 {
		return res
	}

	left, right = 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)/2
		if target == arr[mid] {
			res[1] = mid
			left = mid + 1
		}

		if target < arr[mid] {
			left = mid + 1
		}

		if target > arr[mid] {
			right = mid - 1
		}
	}

	return res
}
