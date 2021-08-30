package easy

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	return ArrayToBST(nums, 0, len(nums)-1)
}

func ArrayToBST(nums []int, start, end int) *TreeNode {
	if end < start {
		return nil
	}
	mid := (end + start) / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = ArrayToBST(nums, start, mid-1)
	root.Right = ArrayToBST(nums, mid+1, end)
	return root
}
