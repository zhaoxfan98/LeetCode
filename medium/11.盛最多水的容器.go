package medium

func maxArea(height []int) int {
	i, n := 0, len(height)-1
	mArea := 0
	//并将指向较短线段的指针向较长线段那端移动一步。
	for i < n {
		if height[i] < height[n] {
			if mArea < height[i]*(n-i) {
				mArea = height[i] * (n - i)
			}
			i++
		} else {
			if mArea < height[n]*(n-i) {
				mArea = height[n] * (n - i)
			}
			n--
		}
	}
	return mArea
}
