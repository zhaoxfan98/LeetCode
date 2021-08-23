package easy

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i] += 1
			return digits
		} else {
			digits[i] = 0
		}
	}
	//跳出了for循环  说明数字全部是9
	temp := make([]int, len(digits)+1)
	temp[0] = 1
	return temp
}
