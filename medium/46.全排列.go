package medium

var res [][]int

// func permute(nums []int) [][]int {
// 	//记录路径
// 	var track []int
// 	backtrack(nums, track)
// 	return res
// }

//路径：记录在track中

//选择列表：nums中不存在于track的那些元素
//结束条件：nums中的元素全部在track中出现
//  该方法只能保存最后一条路径  ？？？？
// func backtrack(nums, track []int) {
// 	//触发结束条件
// 	if len(track) == len(nums) {
// 		temp := make([]int, len(track))
// 		copy(temp, track)
// 		res = append(res, temp)
// 		return
// 	}
// 	flag := true
// 	for i := 0; i < len(nums); i++ {
// 		//排除不合法的选择
// 		for _, v := range track {
// 			if nums[i] == v {
// 				flag = false
// 			}
// 		}
// 		if !flag {
// 			continue
// 		}
// 		//做选择
// 		track = append(track, nums[i])
// 		//进入下一层决策树
// 		backtrack(nums, track)
// 		//取消选择
// 		track = track[:len(track)-1]
// 	}
// }

func permute(nums []int) [][]int {
	res := [][]int{}
	visited := map[int]bool{}

	var dfs func(path []int)
	dfs = func(path []int) {
		//触发结束条件
		if len(path) == len(nums) {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}
		for _, n := range nums {
			//排除不合法的选择
			if visited[n] {
				continue
			}
			//做选择
			path = append(path, n)
			visited[n] = true
			//进入下一层决策树
			dfs(path)
			//取消选择
			path = path[:len(path)-1]
			visited[n] = false
		}
	}

	dfs([]int{})
	return res
}
