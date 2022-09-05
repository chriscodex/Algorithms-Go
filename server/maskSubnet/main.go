package main

import (
	"math"
)

func bitess(res int, p int) int {
	un := res - p
	var dec int
	for i := 7; i >= 7-un; i-- {
		dec += int(math.Pow(2, float64(i)))
	}
	return dec
}

func powTwo(n int) int {
	for i := 0; i <= 31; i++ {
		p := int(math.Pow(2, float64(i)))
		if p >= n {
			return i
		}
	}
	return 0
}
