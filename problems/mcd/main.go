package main

// Calculate the greatest common factor of 2 numbers

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
