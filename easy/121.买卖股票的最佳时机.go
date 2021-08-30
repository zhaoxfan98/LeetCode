package easy

/*
	确定状态
	找到转移公式
	确定初始条件以及边界条件
	计算结果

定义二维数值dp[len][2],其中dp[i][0]表示第i+1天结束的时候没持有股票的最大利润，
dp[i][1]表示第i+1天结合的时候持有股票的最大利润

如果要求第i+1天结合的时候没持有股票的利润最大，那么有两种情况
1. i+1天我没有没买也没卖，那么利润就是第i天没持有股票的最大利润dp[i-1][0]
2. i+1我们卖了股票，那么最大利润就是那么最大利润就是第 i 天持有股票的最大利润加上第 i+1 天卖出股票的最大利润
dp[i-1][1]+prices[i]

很明显我们可以得出
dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
*/
// func maxProfit1(prices []int) int {

// }

//采用双指针方法
func maxProfit1(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	//赋值得小心   如果不是赋值prices[0]则有可能min比prices中所有值小 会导致错误的结果  例如min=0会导致结果为7
	min, max := prices[0], 0
	for i := 1; i < len(prices); i++ {
		if min > prices[i] {
			min = prices[i]
		}
		if max < prices[i]-min {
			max = prices[i] - min
		}
	}

	return max
}
