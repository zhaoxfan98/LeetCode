package easy

var symbolValues = map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}

func romanToInt(s string) (ans int) {
	n := len(s)
	for i := range s {
		value := symbolValues[s[i]]
		if i < n-1 && value < symbolValues[s[i+1]] {
			ans -= value
		} else {
			ans += value
		}
	}
	return
}

// 通常情况下，罗马数字中小的数字在大的数字的右边。
// 若输入的字符串满足该情况，那么可以将每个字符视作一个单独的值，累加每个字符对应的数值即可。

// 若存在小的数字在大的数字的左边的情况，根据规则需要减去小的数字。
// 对于这种情况，我们也可以将每个字符视作一个单独的值，若一个数字右侧的数字比它大，则将该数字的符号取反。
