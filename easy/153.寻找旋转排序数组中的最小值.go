package easy

func findMin(nums []int) int {
	temp := nums[0]
	for i := len(nums) - 1; i >= 0; i-- {
		if temp > nums[i] {
			temp = nums[i]
		}
	}
	return temp
}
