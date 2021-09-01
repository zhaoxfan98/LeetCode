package easy

//超时
func countPrimes1(n int) (cnt int) {
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			cnt++
		}
	}
	return
}

func isPrime(x int) bool {
	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

/*		埃氏筛
	枚举没有考虑到数与数之间的关联性，考虑这样一个事实：
如果x是质数，那么大于x的x的倍数2x,3x一定不是质数。还可继续优化，
对于一个质数 xx，如果按上文说的我们从 2x2x 开始标记其实是冗余的，应该直接从 x*x 开始标记
因为 2x,3x,… 这些数一定在x之前就被其他数的倍数标记过了
*/
func countPrimes(n int) (cnt int) {
	isPrime := make([]bool, n)
	for i := range isPrime {
		isPrime[i] = true
	}
	for i := 2; i < n; i++ {
		if isPrime[i] {
			cnt++
			for j := 2 * i; j < n; j += i {
				isPrime[j] = false
			}
		}
	}
	return
}
