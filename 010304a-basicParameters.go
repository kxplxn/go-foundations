package main

import "fmt"

func addA(a, b, c int) int {
	return a + b + c
}

func basicParameters() {
	fmt.Println("Running `basicParameters`...")
	a, b, c := 20, 10, 5
	fmt.Printf("Result of `addA(a, b, c)` is `%v`.\n\n", addA(a, b, c))
}
