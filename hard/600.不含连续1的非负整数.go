package hard

//暴力 超时
// func findIntegers(n int) (cnt int) {
// 	for i := 0; i <= n; i++ {
// 		if i&(i>>1) == 0 {
// 			cnt++
// 		}
// 	}
// 	return
// }

//动态规划
func findIntegers(n int) (cnt int) {
	//预处理第i层满二叉树的路径数量
	dp := make([]int, 31)
	dp[0], dp[1] = 1, 1
	for i := 2; i < 31; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	//pre记录上一层的根节点值
	pre := 0
	for i := 29; i >= 0; i-- {
		val := 1 << i
		//判断当前子树是否有右子树
		if (n & val) != 0 {
			//有右子树  先将左子树（满二叉树）的路径加到结果中
			cnt += dp[i+1]
			//处理右子树
			if pre == 1 {
				//上一层为1，之后要处理的右子树根节点肯定也为1 此时连续两个1，不满足要求
				break
			}
			//标记当前根节点为1
			pre = 1
		} else {
			//无右子树 此时不能保证左子树是否为满二叉树 所以要在下一层再继续判断
			pre = 0
		}
		if i == 0 {
			cnt++
		}
	}

	return
}
