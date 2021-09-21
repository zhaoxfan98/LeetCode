package medium

func findNumberOfLIS(nums []int) int {
	NoLIS, LIS := make([]int, len(nums)), make([]int, len(nums))
	max := 1
	for i := 0; i < len(nums); i++ {
		NoLIS[i] = 1
		LIS[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				//LIS[i]被LIS[j]+1直接更新，此时同步更新NoLIS
				if LIS[i] < LIS[j]+1 {
					LIS[i] = LIS[j] + 1
					NoLIS[i] = NoLIS[j]
				} else if LIS[i] == LIS[j]+1 {
					//找到了一个新的符合条件的前驱，此时将值继续累加到方案数当中
					NoLIS[i] += NoLIS[j]
				}
			}
		}
		max = Max(max, LIS[i])
	}
	ans := 0
	for i := 0; i < len(nums); i++ {
		if LIS[i] == max {
			ans += NoLIS[i]
		}
	}
	return ans
}

func Max(i, j int) int {
	if i < j {
		return j
	}
	return i
}
