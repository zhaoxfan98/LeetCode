package easy

import "sort"

func arrayPairSum(nums []int) int {
	sum := 0
	sort.Ints(nums)
	for idx, val := range nums {
		if idx%2 == 0 {
			sum += val
		}
	}
	return sum
}
