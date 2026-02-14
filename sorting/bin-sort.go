package sorting

import "github.com/rahulparihar-30/dsa-in-golang/Linkedlist"

func Max(a []int) (max int) {
	max = a[0]
	for _, num := range a {
		if num > max {
			max = num
		}
	}
	return max
}

func BinSort(nums []int) {
	largest := Max(nums)

	bins := make([]*Linkedlist.Node, largest+1)

	for _, num := range nums {
		bins[num] = Linkedlist.Insert(bins[num], num)
	}
	index := 0
	for i := 0; i <= largest; i++ {
		curr := bins[i]
		for curr != nil {
			nums[index] = curr.Data
			index++
			curr = curr.Next
		}
	}
}
