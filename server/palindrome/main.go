package main

/* Build a server where the client sends a word and the server responds if it is a palindrome or not */

func palindrome(s string) bool {
	str := ""
	for i := len(s) - 1; i >= 0; i-- {
		str += string(s[i])
	}
	return str == s
}
