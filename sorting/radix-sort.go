package sorting

import "github.com/rahulparihar-30/dsa-in-golang/Linkedlist"

func CountDigits(num int) (count int) {
	if num == 0 {
		return 1
	}
	count = 0
	if num < 0 {
		num = -num
	}

	for num > 0 {
		num = num / 10
		count++
	}
	return count
}

func RadixSort(nums []int) {
	largest := CountDigits(Max(nums))
	exp := 1
	for pass := 1; pass <= largest; pass++ {
		bins := make([]*Linkedlist.Node, 10)

		for _, num := range nums {
			digit := (num / exp) % 10
			bins[digit] = Linkedlist.Insert(bins[digit], num)
		}

		index := 0

		for _, bin := range bins {
			curr := bin
			for curr != nil {
				nums[index] = curr.Data
				index++
				curr = curr.Next
			}
		}
		exp *= 10
	}
}
