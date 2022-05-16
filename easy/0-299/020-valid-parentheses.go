package main

func isValid(s string) bool {
	stack := []rune{}
	opens := "([{"
	closes := "}])"

	for _, ch := range s {
		for _, open := range opens {
			if ch == open {
				stack = append(stack, ch)
			}
		}
		for _, clos := range closes {
			if ch == clos {
				if len(stack) == 0 || !isMatch(stack[len(stack)-1], ch) {
					return false
				}
				stack = stack[:len(stack)-1]
			}
		}
	}
	return len(stack) == 0
}

func isMatch(open, clos rune) bool {
	if open == '{' && clos == '}' {
		return true
	}
	if open == '(' && clos == ')' {
		return true
	}
	if open == '[' && clos == ']' {
		return true
	}
	return false
}
