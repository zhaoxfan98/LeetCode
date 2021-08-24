package easy

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	cnt := [26]int{}
	for _, ch := range s {
		cnt[ch-'a']++
	}
	for _, ch := range t {
		if cnt[ch-'a'] != 0 {
			cnt[ch-'a']--
		} else {
			return false
		}
	}
	return true
}
