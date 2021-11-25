package _010307_structs

import "fmt"

func StructReferences() {
	fmt.Println("Running `StructReferences`...")
	b := Box{Depth: 5, Width: 4, Height: 3}
	fmt.Printf("`b` before: `%v`.\n", b)
	b.Depth = 6
	fmt.Printf("`b`  after: `%v`.\n\n", b)
}
