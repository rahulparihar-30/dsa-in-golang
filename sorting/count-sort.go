package sorting

func CountSort(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}
	maxx := nums[0]
	for _, num := range nums {
		if num > maxx {
			maxx = num
		}
	}
	frequency := make([]int, maxx+1)

	for _, num := range nums {
		frequency[num]++
	}

	i := 0

	for index, cnt := range frequency {
		for cnt > 0 {
			nums[i] = index
			cnt--
			i++
		}
	}
	return nums
}
