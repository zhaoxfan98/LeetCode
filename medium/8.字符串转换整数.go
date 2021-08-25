package medium

import (
	"fmt"
	"math"
)

func myAtoi(s string) int {
	//先把空格去掉
	index := 0
	//这种方法判断空格有问题
	// for _, v := range s {
	//     if v == ' ' {
	//         index ++
	//     }
	// }
	for index < len(s) && s[index] == ' ' {
		index++
	}
	//极端情况 " " ''
	if index >= len(s) {
		return 0
	}
	//判断正负号
	sign := 1
	if s[index] == '-' {
		sign = -1
		index++
	} else if s[index] == '+' {
		index++
	}
	//遇到非数字return
	result := 0
	for ; index < len(s); index++ {
		if s[index] > '9' || s[index] < '0' {
			break
		}
		num := int(s[index] - '0')
		fmt.Println(num)
		result = result*10 + num
		//判断是否越界
		if sign*result < math.MinInt32 { //整数超过 32 位有符号整数范围
			return math.MinInt32
		} else if sign*result > math.MaxInt32 {
			return math.MaxInt32
		}
	}
	//带上符号返回结果
	return result * sign
}
