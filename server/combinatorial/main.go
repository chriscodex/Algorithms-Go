package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

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

func combinatorialHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/combinatorial", combinatorialHandler).Methods(http.MethodPost)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}
