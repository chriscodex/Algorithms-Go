package main

import "fmt"

func main() {
	num := fact(4)
	fmt.Println(num)
}

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}
