package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var base = 2
var arr []int
var bin []int

func div(c int) int {
	if c <= 1 {
		arr = append(arr, c%base)
		return 1
	}
	if c >= base {
		arr = append(arr, c%base)
		return div(c / base)
	}
	return div(c / base)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	num, _ := strconv.Atoi(string(scanner.Text()))
	if num > 255 {
		fmt.Println("Error, ingrese un numero menor a 256")
		return
	}
	div(8)
	if len(arr) < 8 {
		tam := 8 - len(arr)
		for i := 0; i < tam; i++ {
			arr = append(arr, 0)
		}
	}
	for i := len(arr) - 1; i >= 0; i-- {
		bin = append(bin, arr[i])
	}
	fmt.Println(bin)
}
