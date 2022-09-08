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
}
