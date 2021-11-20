package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func returnSingleValue() {
	fmt.Println("Running `returnSingleValue`...")
	x, y := 20, 10
	res := add(x, y)
	fmt.Printf("Result of `add(x, y)` is `%d`.\n\n", res)
}
