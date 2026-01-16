package searching

// Time Complexity : O(log n)
// Space Complexity : O(n)
func BinarySearchRecursive(nums [10]int, start, end, target int) (bool, int) {
	if start < end {
		mid := int(start + (end-start)/2)

		if nums[mid] == target {
			return true, mid
		} else if nums[mid] < target {
			return BinarySearchRecursive(nums, mid+1, end, target)
		} else {
			return BinarySearchRecursive(nums, start, mid-1, target)
		}
	}
	return false, -1
}

// Time Complexity : O(log n)
// Space Complexity : O(1)

func BinarySearchIterative(nums [10]int, start, end, target int) (bool, int) {
	for start < end {
		mid := int(start + (end-start)/2)

		if nums[mid] == target {
			return true, mid
		} else if nums[mid] < target {
			return BinarySearchIterative(nums, mid+1, end, target)
		} else {
			return BinarySearchIterative(nums, start, mid-1, target)
		}
	}
	return false, -1
}
