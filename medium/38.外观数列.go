package medium

import (
	"strconv"
)

func countAndSay(n int) string {
	//初始化一个存储结果的数组
	arr := "1"
	//n=1为特例
	if n == 1 {
		return arr
	}
	//第一层循环 循环中的第1次将生成n=i+1时的外观数列
	for i := 1; i < n; i++ {
		word := arr[0] //当前搜查的数字
		num := 1       //当前搜查的数字在当前arr序列中出现的个数

		temp := "" //存储新生成序列的字符串

		//第二层循环 循环中的第j次将扫描元素arr[j]
		for j := 1; j < len(arr); j++ {
			if word == arr[j] {
				num++
			} else if word != arr[j] {
				//出现新的数字
				num_char := strconv.Itoa(num)
				temp += num_char     //将个数加入新序列中
				temp += string(word) //将搜寻的数字加入序列中

				word = arr[j] //reset  将word设为新的被搜寻数字
				num = 1       //reset
			}
		}
		//边界处理
		//上面的循环结束后，序列中排在最后一个的数字虽然被找到，但并没有被加入到新序列中，因此要把它也补进去
		num_char := strconv.Itoa(num)
		temp += num_char     //将个数加入新序列中
		temp += string(word) //将搜寻的数字加入序列中
		arr = temp           //将新生成的序列赋给arr
	}

	return arr
}

//迭代  太难懂了 只能一步一步调试
func countAndSay1(n int) string {
	s := "1"
	for i := 0; i < n-1; i++ {
		j, tmp := 0, ""
		for k, c := range s {
			if c != rune(s[j]) {
				//数量+字符
				tmp += strconv.Itoa(k-j) + string(s[j])
				//从出现不同的下一个字符继续搜索
				j = k
			}
		}
		//边界处理 上面的循环结束后，序列中排在最后一个的数字虽然被找到，但并没有被加入到新序列中，因此要把它也补进去
		s = tmp + strconv.Itoa(len(s)-j) + string(s[j])
	}
	return s
}
