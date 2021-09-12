package hard

import "strings"

var res [][]string

func solveNQueens(n int) [][]string {
	res = [][]string{}
	board := make([][]string, n)
	for i := 0; i < n; i++ {
		board[i] = make([]string, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			board[i][j] = "."
		}
	}
	backtrack(board, 0, n)
	return res
}

func backtrack(board [][]string, row, n int) {
	//确定结束条件
	if row == n {
		temp := make([]string, n)
		for i := 0; i < n; i++ {
			temp[i] = strings.Join(board[i], "")
		}
		res = append(res, temp)
		return
	}
	for col := 0; col < n; col++ {
		// 排除不合法选择
		if !isValid(board, row, col) {
			continue
		}
		// 做选择
		board[row][col] = "Q"
		// 进入下一行决策
		backtrack(board, row+1, n)
		// 撤销选择
		board[row][col] = "."
	}
}

func isValid(board [][]string, row, col int) (res bool) {
	n := len(board)
	for i := 0; i < row; i++ {
		if board[i][col] == "Q" {
			return false
		}
	}
	for i := 0; i < n; i++ {
		if board[row][i] == "Q" {
			return false
		}
	}
	//左上
	for i, j := row, col; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == "Q" {
			return false
		}
	}
	//右上
	for i, j := row, col; i >= 0 && j < n; i, j = i-1, j+1 {
		if board[i][j] == "Q" {
			return false
		}
	}
	return true
}
