package main

type Numbs struct {
	Num1 int `json:"num_1"`
	Num2 int `json:"num_2"`
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
