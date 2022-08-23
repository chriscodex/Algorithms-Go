package main

// Recursive Algorithm - Calculate the sum of the elements of an array

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	s := sum(arr, 0)
	fmt.Println(s)
}

func sum(arr []int, pos int) int {
	if pos == len(arr)-1 {
		return arr[pos]
	}
	return arr[pos] + sum(arr, pos+1)
}
