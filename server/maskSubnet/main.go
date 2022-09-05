package main

import (
	"encoding/json"
	"log"
	"math"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

type subnet struct {
	Hosts int `json:"hosts"`
}

func subnetmaskHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	hosts := subnet{}
	err := decoder.Decode(&hosts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	subnetmask := maskSubNet(hosts.Hosts)

}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/subnetmask", subnetmaskHandler).Methods(http.MethodPost)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}
