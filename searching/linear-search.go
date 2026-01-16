package searching

// Time Complexity : O(n)
// Space Complexity : O(1)
func LinearSearch(nums [10]int, target int) (bool, int) {
	for index, num := range nums {
		if num == target {
			return true, index
		}
	}
	return false, -1
}
