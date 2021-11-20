package main

import "fmt"

// `*int` means `Integer Pointer`s.
func multiply(x, y *int) int {
	return *x * *y // `*` suffixes are `Dereference Operator`s.
}

func pointerArguments() {
	fmt.Println("Running `pointerArguments`...")
	x, y := 5, 2
	// `&` here are `Address-Of Operator`s.
	res := multiply(&x, &y)
	fmt.Printf("Result of `multiply(&x, &y)` is `%d`.\n\n", res)
}
