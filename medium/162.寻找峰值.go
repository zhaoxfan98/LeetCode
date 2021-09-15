package medium

func findPeakElement(nums []int) int {
	start, end := 0, len(nums)-1
	if len(nums) == 1 {
		return 0
	}
	if len(nums) == 2 {
		if nums[0] > nums[1] {
			return 0
		}
		return 1
	}
	for start < end {
		mid := start + (end-start)/2
		if nums[mid] > nums[mid+1] {
			//左半边一定有峰值
			end = mid
		} else {
			start = mid + 1
		}
	}
	return start
}
