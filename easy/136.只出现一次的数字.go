package easy

func singleNumber(nums []int) int {
	single := 0
	//数组中的全部元素的异或运算结果即为数组中只出现一次的数字
	for _, num := range nums {
		single ^= num
	}
	return single
}
