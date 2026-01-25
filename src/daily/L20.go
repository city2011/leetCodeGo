package main

func isValid(s string) bool {
	q := make([]rune, 0)

	match := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, c := range s {
		if m, ok := match[c]; ok {
			if len(q) == 0 || q[len(q)-1] != m {
				return false
			}
			q = q[:len(q)-1]
		} else {
			q = append(q, c)
		}
	}

	return len(q) == 0
}
