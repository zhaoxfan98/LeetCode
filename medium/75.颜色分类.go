package medium

func sortColors(nums []int) {
	if len(nums) <= 1 {
		return
	}

	r, w, b := 0, 0, 0
	for _, v := range nums {
		switch v {
		case 0:
			r++
		case 1:
			w++
		case 2:
			b++
		}
	}
	for i := 0; i < r; i++ {
		nums[i] = 0
	}

	for i := r; i < r+w; i++ {
		nums[i] = 1
	}
	for i := r + w; i < len(nums); i++ {
		nums[i] = 2
	}
}
