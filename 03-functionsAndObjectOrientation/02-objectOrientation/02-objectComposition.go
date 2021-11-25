package _010302_objectOrientation

import "fmt"

type PartA struct {
	Manufacturer string
}

func (p *PartA) Mfc() string {
	return p.Manufacturer
}

type Tire struct {
	Part PartA
}

func ObjectComposition() {
	fmt.Println("Running `objectComposition`...")
	tire := Tire{PartA{Manufacturer: "Baban"}}
	fmt.Printf("Result of `tire.PartA.Mfc()` is `%s`.\n\n", tire.Part.Mfc())
}
