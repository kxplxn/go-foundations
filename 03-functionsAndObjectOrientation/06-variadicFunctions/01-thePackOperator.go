package _010306_variadicFunctions

import "fmt"

// The pack operator is used when we declare variadic functions, because the
// arguments are packed into a slice when they are passed to the function.
//
// A function declared this way is called a variadic function.

func ThePackOperator() {
	fmt.Println("Running `ThePackOperator`...")
	res := addF(1, 2, 3, 4)
	fmt.Printf("Result of `addF(1, 2, 3, 4)` is `%s`.\n\n", res)
}

// The pack operator (...) is used here for accepting an arbitrary number of integers.

func addF(vals ...int) string {
	total := 0
	for _, val := range vals {
		total += val
	}
	return fmt.Sprint(vals, total)
}
