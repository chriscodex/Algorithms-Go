package main

import "fmt"

var arr []string

func main() {
	sli := "abcde"
	printReversed(sli, 1)
}

func printReversed(s string, i int) {
	if i > len(s) {
		chain := ""
		for _, e := range arr {
			chain += e
		}
		fmt.Println(chain)
		return
	}
	arr = append(arr, string(s[len(s)-i]))
	printReversed(s[:len(s)-i], i)
}
