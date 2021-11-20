package main

import "fmt"

type PartB struct {
	Manufacturer string
}

func (p *PartB) Mfc() string {
	return p.Manufacturer
}

type TireB struct {
	PartB // anonymous field AKA embedded type
}

func anonymousFields() {
	fmt.Println("Running `anonymousFields`...")
	tire := TireB{PartB{Manufacturer: "Brodacaro"}}
	fmt.Printf("Result of `tire.Mfc()` is `%s`.\n\n", tire.Mfc())
}
