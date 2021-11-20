package main

import "fmt"

// Result declaration defines the return value.
// Since this simple return value doesn't have a name it's called an anonymous result.
func addC(x, y int) int {
	return x + y
}

func returnSingleValue() {
	fmt.Println("Running `returnSingleValue`...")
	x, y := 20, 10
	res := addC(x, y)
	fmt.Printf("Result of `addC(x, y)` is `%d`.\n\n", res)
}
