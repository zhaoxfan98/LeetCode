package easy

import "sort"

func firstBadVersion(n int) int {
	//Search 函数采用二分法搜索找到 [0, n) 区间内最小的满足 f (i)==true 的值 i
	return sort.Search(n, func(version int) bool { return isBadVersion(version) })
}

func isBadVersion(version int) bool {
	return true
}
