package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	stack := []rune{}
	m := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	for _, e := range s {
		if value, ok := m[e]; ok {
			stack = append(stack, value)
			continue
		}
		l := len(stack) - 1
		if l < 0 || stack[l] != e {
			return false
		}
		stack = stack[:l]
	}
	return len(stack) == 0
}

type parentheses struct {
	Input string `json:"input"`
}

func vpHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	par := parentheses{}

	err := decoder.Decode(&par)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/validparentheses", vpHandler).Methods(http.MethodPost)

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
