package medium

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[j]+1, dp[i])
			}
		}
	}
	max := 1
	for _, v := range dp {
		if v > max {
			max = v
		}
	}
	return max

}

// func max(i, j int) int {
// 	if i < j {
// 		return j
// 	}
// 	return i
// }
