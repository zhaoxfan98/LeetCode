package easy

func climbStairs(n int) int {
	//爬到第N级台阶的方法等于爬到第N-1级台阶的方法与爬到N-2级台阶的方法之和
	if n <= 2 {
		return n
	}
	n_1, n_2 := 1, 2
	for i := 3; i <= n; i++ {
		n_1, n_2 = n_2, n_1+n_2
	}
	return n_2
}
