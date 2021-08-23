package easy

//由于同一个数字在两个数组中都可能出现多次，因此需要用哈希表存储每个数组出现的次数
//对于一个数字，其在交集中出现的次数等于该数字在两个数组中出现次数的最小值

func intersect(nums1 []int, nums2 []int) []int {
	//为了降低空间复杂度，首先遍历较短的数组并在哈希表中记录每个数字以及对应出现的次数
	//然后遍历较长的数组得到交集。
	// if len(nums1) > len(nums2) {
	//     intersect(nums2, nums1)
	// }
	//没有降低空间复杂度	？？
	m := map[int]int{}
	for _, num := range nums1 {
		m[num]++
	}

	intersect := []int{}
	for _, num := range nums2 {
		if m[num] > 0 {
			intersect = append(intersect, num)
			m[num]--
		}
	}
	return intersect
}
