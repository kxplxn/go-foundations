package _010302_objectOrientation

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

func AnonymousFields() {
	fmt.Println("Running `AnonymousFields`...")
	tire := TireB{PartB{Manufacturer: "Brodacaro"}}
	fmt.Printf("Result of `tire.Mfc()` is `%s`.\n\n", tire.Mfc())
}
