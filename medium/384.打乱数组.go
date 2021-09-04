package medium

import "math/rand"

type Solution struct {
	nums []int
}

func Constructor(nums []int) Solution {
	return Solution{nums}
}

func (this *Solution) Reset() []int {
	return this.nums
}

func (this *Solution) Shuffle() []int {
	nums := make([]int, len(this.nums))
	copy(nums, this.nums)
	rand.Shuffle(len(nums), func(i, j int) {
		nums[i], nums[j] = nums[j], nums[i]
	})
	return nums
}

func (this *Solution) Shuffle2() []int {
	nums := make([]int, len(this.nums))
	buf := make([]int, len(this.nums))
	copy(buf, this.nums)
	for i := range nums {
		j := rand.Intn(len(buf))
		nums[i] = buf[j]
	}

	return nums
}
