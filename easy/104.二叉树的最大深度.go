package easy

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//递归
func maxDepth_recur(root *TreeNode) int {
	if root == nil {
		return 0
	} else {
		return max(maxDepth_recur(root.Left), maxDepth_recur(root.Right)) + 1
	}
}

//DFS
func maxDepth_dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth_dfs(root.Left), maxDepth_dfs(root.Right)) + 1
}

//BFS
func maxDepth_bfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{}
	queue = append(queue, root)
	ans := 0
	for len(queue) > 0 {
		sz := len(queue)
		for sz > 0 {
			node := queue[0]
			queue := queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			sz--
		}
		ans++
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
