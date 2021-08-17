package easy

func moveZeroes(nums []int) {
	fast := 1
	length := len(nums)
	for slow := 0; slow < length; slow++ {
		if nums[slow] == 0 {
			for ; fast < length; fast++ {
				if nums[fast] != 0 {
					nums[slow] = nums[fast]
					nums[fast] = 0
					break
				}
			}
		}
		fast++
	}
}
