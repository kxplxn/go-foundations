package _010306_variadicFunctions

import "fmt"

/*
The unpack operator simply unpacks the contents of a slice when we call a
varidic function with a slice.
*/
func TheUnpackOperator() {
	fmt.Println("Running `TheUnpackOperator`...")
	numbers := []int{1, 2, 3, 4, 5, 6}
	res := addF(numbers...)
	fmt.Printf("Result of `addF(numbers...)` is `%s`.\n\n", res)
}
