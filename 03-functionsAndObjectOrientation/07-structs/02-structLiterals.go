package _010307_structs

import "fmt"

/*
With struct literals (eg. `D: 20.0`), the order of the named fields is
irrelevant; and ommited fields are set to their zero values for their tyype.
*/

func StructLiterals() {
	fmt.Println("Running `StructLiterals`...")
	b := Box{Depth: 20.0, Width: 15.0, Height: 25.0}
	fmt.Printf("Value of `b` is %v.\n\n", b)
}
