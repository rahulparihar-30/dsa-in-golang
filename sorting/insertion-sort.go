package sorting

func InsertionSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		j := i - 1
		x := arr[i]
		for ; j >= 0 && x < arr[j]; j-- {
			arr[j+1] = arr[j]
		}
		arr[j+1] = x
	}
	return arr
}
