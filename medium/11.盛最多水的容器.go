package medium

func maxArea(height []int) int {
	i, n := 0, len(height)-1
	mArea := 0
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
