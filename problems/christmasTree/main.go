package main

import "fmt"

func addSpaces(str string, cant int) string {
	var chain string
	for i := 0; i < cant; i++ {
		chain += " "
	}
	return chain + str
}

func asteriskGenerator(cant int) string {
	var ast string
	for i := 0; i < cant; i++ {
		ast += "*"
	}
	fmt.Println(ast)
	return ast
}

func buildTree(numb int) {
	var tree string
	alt := 2 * numb
	cant := numb
	for j := 1; j <= alt; j += 2 {
		tree = asteriskGenerator(j)
		tree = addSpaces(tree, cant-1)
		fmt.Println(tree)
		cant--
	}
}

func main() {
	buildTree(6)
}
