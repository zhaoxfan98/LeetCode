package easy

import "strings"

func firstUniqChar(s string) int {
	for i, v := range s {
		str := string(v)
		if strings.Count(s, str) == 1 {
			return i
		}
	}
	return -1
}

//使用哈希表存储频数 两次遍历
func firstUniqChar1(s string) int {
	cnt := [26]int{}
	for _, ch := range s {
		cnt[ch-'a']++
	}
	for i, ch := range s {
		if cnt[ch-'a'] == 1 {
			return i
		}
	}
	return -1
}
