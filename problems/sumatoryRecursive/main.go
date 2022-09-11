package main

func sum(num int) int {
	if num == 1 {
		return 1
	}
	return num + sum(num-1)
}
