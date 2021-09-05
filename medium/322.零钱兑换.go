package medium

import (
	"math"
)

/*暴力解法
dp(n)的定义：输入一个目标金额n，返回凑出目标金额n的最少硬币数量

*/
func coinChange1(coins []int, amount int) int {
	return dp1(coins, amount)
}

func dp1(coins []int, amount int) int {
	//base case
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}

	res := math.MaxInt32
	for coin := range coins {
		//计算子问题的解
		subP := dp1(coins, amount-coin)
		//子问题无解则跳过
		if subP == -1 {
			continue
		}
		//在子问题中选择最优解 然后加一
		res = min(res, subP+1)
	}
	if res == math.MaxInt32 {
		return -1
	} else {
		return res
	}
}

//带备忘录的递归
var memo []int

func coinChange2(coins []int, amount int) int {
	memo = make([]int, amount+1)
	//初始化memo
	memo[0] = 0
	for i := 1; i < amount+1; i++ {
		memo[i] = -100
	}
	return dp2(coins, amount)
}

func dp2(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}
	//查备忘录 防止重复计算
	if memo[amount] != -100 {
		return memo[amount]
	}
	res := math.MaxInt32
	for _, coin := range coins {
		//計算子问题的结果
		subP := dp2(coins, amount-coin)
		//子问题无解则跳过
		if subP == -1 {
			continue
		}
		//在子问题中选择最优解 然后加一
		res = min(res, subP+1)
	}
	// **把计算结果存入备忘录 不然无法跳出子问题的递归调用**
	if res == math.MaxInt32 {
		memo[amount] = -1
	} else {
		memo[amount] = res
	}
	return memo[amount]
}

//dp数组的迭代解法
func coinChange3(coins []int, amount int) int {
	dp := make([]int, amount+1)
	//初始化 数组大小为amount+1， 初始值也为amount+1
	for i:=1; i<amount+1; i++ {
		dp[i] = amount+1
	}

	//base case
	dp[0] = 0
	//外层for循环遍历所有状态的所有取值
	for i:=0; i<len(dp); i++ {
		//内层循环求所有选择的最小值
		for _, coin := range coins {
			//子问题无解 跳过
			if i - coin < 0 {
				continue
			}
			dp[i] = min(dp[i], dp[i-coin]+1)
		}
	}
	if dp[amount] == amount+1 {
		return -1
	} else {
		return dp[amount]
	}
}
