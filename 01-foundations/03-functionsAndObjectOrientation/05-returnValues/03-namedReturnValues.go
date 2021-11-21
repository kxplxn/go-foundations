package _010305_returnValues

import "fmt"

func addE(x, y int) (a, b int, c bool) {
	a = x + y
	b = x - y
	c = x > y
	return
	// or `return a, b, c`
	//or `return a, b, false`
}

func NamedReturnValues() {
	fmt.Println("Running `namedReturnValues`...")
	a, b, c := addE(20, 10)
	fmt.Printf("Result of `addE(20, 10)` is `%d %d %t`.\n\n", a, b, c)
}
