package main

import "fmt"

/* Escribe un programa que muestre los números del 1 al 100, sustituyendo los múltiplos de 3 por la última palabra "fizz",
los múltiplos de 5 por "buzz" y los múltiplos de ambos por la palabra "fizzbuzz" */

func main() {
	arr := []any{}
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 != 0 {
			arr = append(arr, "fizz")
		} else if i%5 == 0 && i%5 != 0 {
			arr = append(arr, "buzz")
		} else if i%3 == 0 && i%5 == 0 {
			arr = append(arr, "fizzbuzz")
		} else {
			arr = append(arr, i)
		}
	}
	fmt.Println(arr)
}
