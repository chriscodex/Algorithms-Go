package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

type decimalToBin struct {
	Number int `json:"number"`
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/decimalToBinary", decimalToBinaryHandler).Methods(http.MethodPost)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}
