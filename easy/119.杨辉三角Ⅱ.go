package easy

//第n行的第i个数等于第n-1行的第i-1个数和第i个数之和

func getRow(rowIndex int) []int {
	C := make([][]int, rowIndex+1)
	for i := range C {
		C[i] = make([]int, i+1)
		C[i][0], C[i][i] = 1, 1
		for j := 1; j < i; j++ {
			C[i][j] = C[i-1][j-1] + C[i-1][j]
		}
	}
	return C[rowIndex]
}

//注意到对第 i+1i+1 行的计算仅用到了第 ii 行的数据，因此可以使用滚动数组的思想优化空间复杂度
func getRow1(rowIndex int) []int {
	var pre, cur []int
	for i := 0; i <= rowIndex; i++ {
		cur = make([]int, i+1)
		cur[0], cur[i] = 1, 1
		for j := 1; j < i; j++ {
			cur[j] = pre[j-1] + pre[j]
		}
		pre = cur
	}
	return pre
}
