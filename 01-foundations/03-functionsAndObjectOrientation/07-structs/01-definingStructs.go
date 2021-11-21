package _010307_structs

import "fmt"

// Box is an exportable object.
// We should always comment exported structs.
type Box struct {
	Depth  float64
	Width  float64
	Height float64
}

func DefiningStructs() {
	fmt.Println("Running `definingStructs`...")
	b := Box{5, 4, 4}
	fmt.Printf("Value of `b` is `%v`.\n", b)
}
