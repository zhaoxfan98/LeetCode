package medium

func removeDuplicates(nums []int) int {
	//既然相同数字最多只能保存两个，又因为数组有序，相同的数字必然连续，
	//只需判断数字与它相隔前两个位置的数字是否相同，相同则不存，不同才存
	n := len(nums)
	if n < 3 {
		return n
	}
	i := 0
	for j := 2; j < n; j++ {
		if nums[i] != nums[j] {
			nums[i+2] = nums[j]
			i = i + 1
		}
	}
	return i + 2
}
