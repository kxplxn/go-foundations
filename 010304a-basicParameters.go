package main

import "fmt"

func add(a, b, c int) int {
	return a + b + c
}

func basicParameters() {
	fmt.Println("Running `valueParameters`...")
	a, b, c := 20, 10, 5
	fmt.Printf("Result of `add(a, b, c)` is `%v`.\n\n", add(a, b, c))
}
