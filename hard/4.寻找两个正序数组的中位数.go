package hard

import "sort"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	nums := make([]int, m+n)
	for i, v := range nums1 {
		nums[i] = v
	}
	for i, v := range nums2 {
		nums[i+m] = v
	}
	sort.Ints(nums)
	mid := (m + n) / 2
	if (m+n)%2 != 0 {
		return float64(nums[mid])
	} else {
		return float64(nums[mid-1]+nums[mid]) / 2.0
	}
}
