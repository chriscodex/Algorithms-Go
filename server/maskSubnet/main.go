package main

import (
	"math"
	"strconv"
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

func maskSubNet(num int) string {
	p := powTwo(num)
	if p <= 31 && p >= 24 {
		s := ""
		s += strconv.Itoa(bitess(31, p))
		s += ".0.0.0"
		return s
	}
	if p <= 23 && p >= 16 {
		s := "255."
		s += strconv.Itoa(bitess(23, p))
		s += ".0.0"
		return s
	}
	if p <= 15 && p >= 8 {
		s := "255.255."
		s += strconv.Itoa(bitess(15, p))
		s += ".0"
		return s
	}
	if p <= 7 && p >= 0 {
		s := "255.255.255."
		s += strconv.Itoa(bitess(7, p))
		return s
	}
	return ""
}
