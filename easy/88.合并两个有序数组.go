package easy

import "sort"

func merge(nums1 []int, m int, nums2 []int, n int) {
	for i := m; i < m+n; i++ {
		nums1[i] = nums2[m+n-i-1]
	}
	sort.Ints(nums1)
}
