package binarysearch

// 1 2 3 4 5 6 7 8 9

func binarySearch(arr []int, left, right, target int) int {

	// check if overflow
	if left >= right {
		return -1
	}

	mid := left + (right-left)/2
	// check if ok
	if arr[mid] == target {
		return mid
	}

	// check left side
	if target < arr[mid] {
		return binarySearch(arr, left, mid, target)
	}

	// check right side
	return binarySearch(arr, mid+1, right, target)
}

func binarySearchIterative(arr []int, target int) int {
	left, right := 0, len(arr)
	for right > left {
		mid := left + (right-left)/2

		if arr[mid] == target {
			return mid
		}

		if target < arr[mid] {
			right = mid
		}

		if target > arr[mid] {
			left = mid + 1
		}
	}

	return -1
}

// arr = [10, 20, 20, 20, 30]
// target = 20
// expected result [1, 3]
func binarySearchIterativeWithDuplicate(arr []int, target int) []int {

	res := []int{-1, -1}

	if len(arr) == 0 {
		return res
	}

	// calculate left most
	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			res[0] = mid
			right = mid - 1
			continue
		}

		if target < arr[mid] {
			right = mid
		}

		if target > arr[mid] {
			left = mid + 1
		}
	}

	if res[0] == -1 {
		return res
	}

	// calculate left most
	left, right = 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			res[1] = mid
			left = mid + 1
			continue
		}

		if target < arr[mid] {
			right = mid - 1
		}

		if target > arr[mid] {
			left = mid + 1
		}
	}

	return res
}
