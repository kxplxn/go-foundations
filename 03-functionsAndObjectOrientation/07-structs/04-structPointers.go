package _010307_structs

import "fmt"

func StructPointers() {
	fmt.Println("Running `StructPointers`...")
	b := Box{Depth: 5, Width: 4, Height: 3}

	fmt.Printf("`b` before: `%v`.\n", b)

	// & -> The "Address-Of" Operator
	ptr := &b

	// * -> The "De-reference" Operator
	//(*ptr).D = 7

	// Implicit De-referencing
	ptr.Depth = 7

	fmt.Printf("`b`  after: `%v`.\n\n", b)
}
