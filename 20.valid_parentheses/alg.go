package _0_valid_parentheses

func isValid(s string) bool {
	if len(s)&1 > 0 {
		return false
	}
	symbolMap := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	var stack []byte
	for i := 0; i < len(s); i++ {
		if symbolMap[s[i]] > 0 {
			if len(stack) == 0 || symbolMap[s[i]] != stack[len(stack)-1] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}
