package main

import (
	"fmt"

	"github.com/rahulparihar-30/dsa-in-golang/sorting"
)

func sort() {
	nums := []int{90, 10, 50, 60, 70, 30, 20}
	//sorting.QuickSort(nums, 0, len(nums)-1)
	sorting.RadixSort(nums)
	fmt.Println(nums)
}

func main() {
	sort()
}
