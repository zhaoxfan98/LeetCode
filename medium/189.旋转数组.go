package medium

func rotate(nums []int, k int) {
	n := len(nums)
	nums1 := make([]int, n)
	for i := 0; i < n; i++ {
		nums1[(i+k)%n] = nums[i]
	}
	for i := 0; i < n; i++ {
		nums[i] = nums1[i]
	}
}
