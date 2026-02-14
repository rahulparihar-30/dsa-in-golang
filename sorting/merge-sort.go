package sorting

func Merge(nums []int, lo, mid, hi int) {
	temp := make([]int, hi-lo+1)
	i, j, k := lo, mid+1, 0

	for i <= mid && j <= hi {
		if nums[i] <= nums[j] {
			temp[k] = nums[i]
			i++
		} else {
			temp[k] = nums[j]
			j++
		}
		k++
	}

	for i <= mid {
		temp[k] = nums[i]
		i++
		k++
	}

	for j <= hi {
		temp[k] = nums[j]
		j++
		k++
	}

	// copy back
	for x := 0; x < len(temp); x++ {
		nums[lo+x] = temp[x]
	}
}

func MergeSort(arr []int, lo, hi int) {
	if lo < hi {
		mid := lo + (hi-lo)/2
		MergeSort(arr, lo, mid)
		MergeSort(arr, mid+1, hi)
		Merge(arr, lo, mid, hi)
	}
}
