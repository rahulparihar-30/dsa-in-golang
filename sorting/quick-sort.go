package sorting

func hoarsePartition(nums []int, low, high int) int {
	pivot := nums[low]
	i := low - 1
	j := high + 1

	for {
		for {
			i++
			if nums[i] >= pivot {
				break
			}
		}
		for {
			j--
			if nums[j] <= pivot {
				break
			}
		}
		// swap
		if i >= j {
			return j
		}
		nums[i], nums[j] = nums[j], nums[i]

	}

}

func QuickSort(arr []int, low, high int) {
	if low < high {
		pivIdx := hoarsePartition(arr, low, high)
		QuickSort(arr, low, pivIdx)
		QuickSort(arr, pivIdx+1, high)
	}
}
