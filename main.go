package main

import (
	"fmt"
)

func main(){
	arr := []string{"++-++++123+++","+321"}
	fmt.Println(arr)
	arr = delm(arr)
	fmt.Println(arr)
}

func delm(arr []string) []string {
	cont := false
	for _,e := range arr {
                if string(e[0])=="+" || string(e[0])=="-" ||
		string(e[0])==" " {
			cont = true
                }
        }
	if !cont {
		return arr
	}
	for i,e := range arr {
                if string(e[0])=="+" || string(e[0])=="-" || 
		string(e[0])==" " {
			arr[i]=e[1:]
               	}
       	}
	return delm(arr)	
}
