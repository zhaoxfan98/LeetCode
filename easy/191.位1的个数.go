package easy

//检查第i位时，让n与2^i进行 与 运算 当且仅当n的第i位为1时，运算结果不为0
func hammingWeight(num uint32) (cnt int) {
	for i := 0; i < 32; i++ {
		if 1<<i&num > 0 {
			cnt++
		}
	}
	return
}
