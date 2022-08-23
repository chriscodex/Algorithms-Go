package main

// Recursive Algorithm - Calculate the greatest common factor of 2 numbers

import "fmt"

var a1 []int
var a2 []int

func main() {
	n := 412
	m := 184
	a1 = append(a1, n)
	a2 = append(a2, m)
	c1, _ := maxcd(a1, a2, 0)
	fmt.Printf("MCD (%d - %d): %d", n, m, c1[len(c1)-1])
}

func maxcd(a1 []int, a2 []int, c int) ([]int, []int) {
	if a1[c] == a2[c] {
		return a1, a2
	}
	if a1[c] > a2[c] {
		a1 = append(a1, a1[c]-a2[c])
		a2 = append(a2, a2[c])
		c++
		return maxcd(a1, a2, c)
	} else {
		a1 = append(a1, a1[c])
		a2 = append(a2, a2[c]-a1[c])
		c++
		return maxcd(a1, a2, c)
	}
}
