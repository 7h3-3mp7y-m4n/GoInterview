package stack

func isValid(s string) bool {
	stack := []rune{}
	pair := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, ch := range s {
		switch ch {
		case '(', '[', '{':
			stack = append(stack, ch)
		case ')', ']', '}':
			if len(stack) == 0 || stack[len(stack)-1] != pair[ch] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
