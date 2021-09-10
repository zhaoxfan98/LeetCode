package medium

const mod = 1000000007

func numSub(s string) (cnt int) {
	index := 0
	for _, v := range s {
		if v == '1' {
			cnt++
			index++
		} else {
			cnt += index * (index - 1) / 2
			index = 0
		}
	}
	if index == 0 {
		return cnt % mod
	} else {
		cnt += index * (index - 1) / 2
		return cnt % mod
	}
}
