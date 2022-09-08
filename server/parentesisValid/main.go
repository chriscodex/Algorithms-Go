package main

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	stack := []rune{}
	m := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	for _, e := range s {
		if value, ok := m[e]; ok {
			stack = append(stack, value)
			continue
		}
	}
	return len(stack) == 0
}
