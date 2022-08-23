package main

/*Write a program that outputs all possibilities to put + or - or nothing between the numbers 1,2,…,9 (in this order) such that
the result is 100. For example 1 + 2 + 3 - 4 + 5 + 6 + 78 + 9 = 100.*/

import (
	"fmt"
)

var perm []string

func main() {
	inp := []string{"1", "2", "3", "+", "-", " "}
	//num := 33 1 + 32
	generate(len(inp), inp)
	perm = deleter(perm)
	for _, e := range perm {
		fmt.Println(e)
	}
}

// Heap Algorithm
func generate(n int, arr []string) {
	if n == 1 {
		s := ""
		for _, e := range arr {
			s += e
		}
		perm = append(perm, s)
		fmt.Println(arr)
	}
	for i := 0; i < n; i++ {
		generate(n-1, arr)
		if n%2 == 0 {
			aux := arr[i]
			arr[i] = arr[n-1]
			arr[n-1] = aux
		} else {
			aux2 := arr[0]
			arr[0] = arr[n-1]
			arr[n-1] = aux2
		}
	}
}

func deleter(arr []string) []string {
	flag := false
	for _, e := range arr {
		if string(e[len(e)-1]) == "+" || string(e[len(e)-1]) == "-" || string(e[len(e)-1]) == " " {
			flag = true
		}
		if string(e[0]) == "+" || string(e[0]) == "-" || string(e[0]) == " " {
			flag = true
		}
	}
	if !flag {
		return arr
	}
	for i, e := range arr {
		if string(e[len(e)-1]) == "+" || string(e[len(e)-1]) == "-" || string(e[len(e)-1]) == " " {
			arr[i] = e[:(len(e) - 1)]
		}
		if string(e[0]) == "+" || string(e[0]) == "-" || string(e[0]) == " " {
			arr[i] = e[1:]
		}
	}
	return deleter(arr)
}
