package _010310_objects_

import "fmt"

type Shape interface {
	volume() float64
}

type Cube struct {
	depth  float64
	width  float64
	height float64
}

func (cu *Cube) volume() float64 {
	return cu.depth * cu.width * cu.height
}

type Sphere struct {
	radius float64
}

func (sp *Sphere) volume() float64 {
	return (4.0 / 3.0) * 3.14 * (sp.radius * sp.radius * sp.radius)
}

func totalVolume(shapes ...Shape) float64 {
	total := 0.0
	for _, sh := range shapes {
		total += sh.volume()
	}
	return total
}

func WorkingWithObjects() {
	fmt.Println("Running `WorkingWithObjects`...")
	cu := Cube{depth: 4, width: 4, height: 4}
	sp := Sphere{radius: 2}
	fmt.Printf("Total volume is `%.2f`.\n\n", totalVolume(&cu, &sp))
}
