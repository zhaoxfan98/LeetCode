package easy

func maxSubArray(nums []int) int {
	dp := []int{}
	dp = append(dp, nums[0])
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > dp[i-1]+nums[i] {
			dp = append(dp, nums[i])
		} else {
			dp = append(dp, dp[i-1]+nums[i])
		}
	}
	//找到dp数组中的最大值即可
	max := dp[0]
	for _, v := range dp {
		if max < v {
			max = v
		}
	}
	return max
}
