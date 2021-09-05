package easy

func fib1(n int) int {

	/*
		//SB 递归方法
		if n == 1 || n == 2 {
			return 1
		}
		return fib(n-1) + fib(n-2)
	*/
	/*
		//带备忘录的递归解法
		memo := make([]int, n+1)
		return helper(memo, n)
	*/
	/*
		//dp数组的迭代解法
		if n == 0 {
			return 0
		}
		dp := make([]int, n+1)
		dp[0], dp[1] = 0, 1
		//状态转移
		for i := 2; i <= n; i++ {
			dp[i] = dp[i-1] + dp[i-2]
		}
		return dp[n]
	*/
	//状态压缩
	if n < 1 {
		return 0
	}
	if n == 2 || n == 1 {
		return 1
	}
	sum, prev, curr := 0, 1, 1
	for i := 3; i <= n; i++ {
		sum = prev + curr
		prev = curr
		curr = sum
	}
	return curr
}

func helper(memo []int, n int) int {
	if n == 0 || n == 1 {
		return n
	}
	//已经计算 不用再计算了
	if memo[n] != 0 {
		return memo[n]
	}
	memo[n] = helper(memo, n-1) + helper(memo, n-2)
	return memo[n]
}
