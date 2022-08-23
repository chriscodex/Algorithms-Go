package main

import "fmt"

func main() {
	s := "1234"
	sus := ""
	for i := len(s) - 1; i >= 0; i-- {
		sus += string(s[i])
	}
	fmt.Println(sus)
}
