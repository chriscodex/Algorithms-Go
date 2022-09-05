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
