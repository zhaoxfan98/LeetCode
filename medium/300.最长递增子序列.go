package medium

//动态规划
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

//二分查找
func lengthOfLIS1(nums []int) int {
	top := make([]int, len(nums))
	//牌堆初始化为0
	piles := 0
	for i := 0; i < len(nums); i++ {
		//要处理的扑克牌
		poker := nums[i]

		//搜索左侧边界的二分查找
		left, right := 0, piles
		for left < right {
			mid := (left + right) / 2
			if top[mid] > poker {
				right = mid
			} else if top[mid] < poker {
				left = mid + 1
			} else {
				right = mid
			}
		}

		//没找到合适的牌堆 新建
		if left == piles {
			piles++
		}
		top[left] = poker
	}

	return piles
}
