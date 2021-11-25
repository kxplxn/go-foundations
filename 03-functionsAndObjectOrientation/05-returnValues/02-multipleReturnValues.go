package _010305_returnValues

import "fmt"

func addD(x, y int) (int, int, bool) {
	a := x + y
	b := x - y
	c := x > y
	return a, b, c
}

func MultipleReturnValues() {
	fmt.Println("Running `multipleReturnValues`...")
	x, y := 20, 10
	a, b, c := addD(x, y)
	fmt.Printf("Result of `addD(x, y)` is `%d %d %t`.\n\n", a, b, c)
}
