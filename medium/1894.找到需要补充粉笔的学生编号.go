package medium

func chalkReplacer(chalk []int, k int) int {
	sum := 0
	for i, v := range chalk {
		sum += v
		if k < sum {
			return i
		}
	}
	k %= sum
	for i, v := range chalk {
		if k < v {
			return i
		}
		k -= v
	}
	return 0
}
