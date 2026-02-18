package a

// IF left half is sorted (arr[left] <= arr[mid]):
//     IF target is between arr[left] and arr[mid]:
//         → search left
//     ELSE:
//         → search right
// ELSE:  // right half must be sorted
//     IF target is between arr[mid] and arr[right]:
//         → search right
//     ELSE:
//         → search left

func search(nums []int, target int) int {
	return binarySearch(nums, 0, len(nums), target)
}

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

	searchLeft := false

	if arr[left] <= arr[mid] {
		if target >= arr[left] && target < arr[mid] {
			// search left
			searchLeft = true
		} else {
			// search right
		}
	} else {
		// right
		if target >= arr[mid] && target <= arr[right-1] {
			// search right
		} else {
			// search left
			searchLeft = true
		}
	}

	// check left side
	if searchLeft {
		return binarySearch(arr, left, mid, target)
	}

	// check right side
	return binarySearch(arr, mid+1, right, target)
}
