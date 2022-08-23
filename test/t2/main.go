package main

import (
	"fmt"
)

// var contF = 0
// var nums []string

// func clasificador(str string) string {
// 	newStr := ""
// 	for string(str[contF]) != "+" {
// 		newStr += str[contF]
// 	}
// }
var lcl []string

func main() {
	str := "12+3+45+8+445"
	s := ""
	cont := len(str) - 1
	ctr := 0
	for string(str[cont]) != "+" {
		s += string(str[cont])
		ctr++
		cont--
	}
	cls(str, 0, len(str)-cont)
	sus := ""
	for i := len(s) - 1; i >= 0; i-- {
		sus += string(s[i])
	}
	lcl = append(lcl, sus)
	fmt.Println(lcl)
}

func cls(str string, ind int, c int) {
	if ind >= (len(str) - c) {
		return
	}
	s := ""
	for string(str[ind]) != "+" {
		s += string(str[ind])
		ind++
	}
	lcl = append(lcl, s)
	cls(str, ind+1, c)
}
