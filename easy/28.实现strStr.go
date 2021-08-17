package easy

func strStr(haystack string, needle string) int {
	L := len(needle)
	n := len(haystack)
	for start := 0; start < n-L+1; start++ {
		if haystack[start:start+L] == needle {
			return start
		}
	}
	return -1
}
