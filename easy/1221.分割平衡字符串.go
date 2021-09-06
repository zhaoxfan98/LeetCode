package easy

func balancedStringSplit(s string) int {
	n, num := 0, 0
	for _, v := range s {
		if v == 'R' {
			n++
		} else {
			n--
		}
		if n == 0 {
			num++
		}
	}
	return num
}
