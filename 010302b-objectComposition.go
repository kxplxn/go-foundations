package main

import "fmt"

type Part struct {
	Manufacturer string
}

func (p *Part) Mfc() string {
	return p.Manufacturer
}

type Tire struct {
	Part Part
}

func objectComposition() {
	fmt.Println("Running `objectComposition`...")
	tire := Tire{Part{Manufacturer: "Baban"}}
	fmt.Printf("Result of `tire.Part.Mfc()` is `%s`.\n\n", tire.Part.Mfc())
}
