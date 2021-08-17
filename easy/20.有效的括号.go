package easy

//左括号类型入栈，右括号类型匹配栈顶元素，如果不是对应关系，则肯定false，匹配上之后移除栈顶元素

func isValid(s string) bool {
	if s == "" {
		return true
	}
	var stack []uint8
	m := map[uint8]uint8{
		'}': '{',
		')': '(',
		']': '[',
	}
	for i := 0; i < len(s); i++ {
		if s[i] == '{' || s[i] == '[' || s[i] == '(' {
			stack = append(stack, s[i])
		} else {
			if len(stack) == 0 {
				return false
			}

			if m[s[i]] != stack[len(stack)-1] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
