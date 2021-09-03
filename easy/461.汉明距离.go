package easy

import "math/bits"

//使用异或运算，记为⊕，当且仅当输入位不同时输出为1
func hammingDistance(x int, y int) int {
	//大多数编程语言都内置了计算二进制表达中1的数量的函数
	return bits.OnesCount(uint(x ^ y))
}
