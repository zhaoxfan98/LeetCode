package easy

import "strings"

func isPalindrome(s string) bool {
	//筛选
	var str string
	for i := 0; i < len(s); i++ {
		if isalnum(s[i]) {
			str += string(s[i])
		}
	}
	n := len(str)
	//判断
	str = strings.ToUpper(str)
	for i := 0; i < n/2; i++ {
		if str[i] != str[n-i-1] {
			return false
		}
	}
	return true
}

func isalnum(ch byte) bool {
	return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
}
