package easy

func lengthOfLastWord(s string) (res int) {
	n := len(s)
	for i := n - 1; i >= 0; i-- {
		if s[i] != ' ' {
			res++
		} else if res != 0 {
			break
		}
	}
	return
}
