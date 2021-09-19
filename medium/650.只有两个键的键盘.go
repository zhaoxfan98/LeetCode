package medium

func minSteps(n int) int {
	//递归
	if n == 1 {
		return 0
	}
	maxfactor := 1
	//找公因式
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			//大于1的最小公因数
			maxfactor = i
			break
		}
	}
	//代表n是质数 返回本身
	if maxfactor == 1 {
		return n
	}
	return maxfactor + minSteps(n/maxfactor)
}
