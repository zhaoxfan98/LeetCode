package medium

func rand10() int {
	a := rand7()
	b := rand7()
	if a > 4 && b < 4 {
		return rand10()
	} else {
		return (a+b)%10 + 1
	}

}

func rand7() int {
	return 0
}
