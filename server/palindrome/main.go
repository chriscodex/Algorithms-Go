package main

import (
	"encoding/json"
	"log"
	"net/http"
)

/* Build a server where the client sends a word and the server responds if it is a palindrome or not */

func palindrome(s string) bool {
	str := ""
	for i := len(s) - 1; i >= 0; i-- {
		str += string(s[i])
	}
	return str == s
}

type palindromeModel struct {
	Chain string `json:"chain"`
}

func palindromeHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	palMD := palindromeModel{}

	decoder.Decode(&palMD)

	res := palindrome(palMD.Chain)

	response, err := json.Marshal(res)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
