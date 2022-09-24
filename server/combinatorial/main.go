package main

import "github.com/gorilla/mux"

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

func comb(m int, n int) int {
	divs := factorial(m-n) * factorial(n)
	return factorial(m) / divs
}

func main() {
	router := mux.NewRouter()
}
