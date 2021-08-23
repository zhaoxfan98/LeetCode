package medium

func isValidSudoku(board [][]byte) bool {
	var row, col, block [9]uint16
	var cur uint16
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			//当前数字的 二进制数位 位置
			cur = 1 << (board[i][j] - '1')
			bi := i/3 + j/3*3 //块索引号
			//使用 与运算符 查重
			if (row[i]&cur)|(col[j]&cur)|(block[bi]&cur) != 0 {
				return false
			}
			// 在对应的位图上，加上当前数字
			row[i] |= cur
			col[j] |= cur
			block[bi] |= cur
		}
	}
	return true
}
