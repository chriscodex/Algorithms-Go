package main

import (
	"fmt"
)

func main() {
	arr1 := []string{"++-++++123++++++", "+321+++- +"}
	fmt.Println(arr1)
	arr1 = mI(arr1)
	fmt.Println(arr1)
}

func mI(arr []string) []string {
	flag := false
	for _, e := range arr {
		if string(e[len(e)-1]) == "+" || string(e[len(e)-1]) == "-" || string(e[len(e)-1]) == " " {
			flag = true
		}
		if string(e[0]) == "+" || string(e[0]) == "-" || string(e[0]) == " " {
			flag = true
		}
	}
	if !flag {
		return arr
	}
	for i, e := range arr {
		if string(e[len(e)-1]) == "+" || string(e[len(e)-1]) == "-" || string(e[len(e)-1]) == " " {
			arr[i] = e[:(len(e) - 1)]
		}
		if string(e[0]) == "+" || string(e[0]) == "-" || string(e[0]) == " " {
			arr[i] = e[1:]
		}
	}
	return mI(arr)
}

func delm(arr []string) []string {
	cont := false
	for _, e := range arr {
		if string(e[0]) == "+" || string(e[0]) == "-" || string(e[0]) == " " {
			cont = true
		}
	}
	if !cont {
		return arr
	}
	for i, e := range arr {
		if string(e[0]) == "+" || string(e[0]) == "-" || string(e[0]) == " " {
			arr[i] = e[1:]
		}
	}
	return delm(arr)
}
