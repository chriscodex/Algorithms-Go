package main

import (
	"strconv"
)

var rr []int

func decimalToBinary(n int) string {
	c := n / 2
	if c == 1 {
		rr = append(rr, 1)
		str := ""
		for i := len(rr) - 1; i >= 0; i-- {
			str += strconv.Itoa(rr[i])
		}
		rr = []int{}
		return str
	}
	rr = append(rr, n%2)
	return decimalToBinary(c)
}
