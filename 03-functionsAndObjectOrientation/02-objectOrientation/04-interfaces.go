package _010302_objectOrientation

import "fmt"

type CubeC struct {
	Depth  float64
	Width  float64
	Height float64
}

func (cube *CubeC) Volume() float64 {
	return cube.Depth * cube.Width * cube.Height
}

type Sphere struct {
	Radius float64
}

func (sphere *Sphere) Volume() float64 {
	return (4.0 / 3.0) * 3.14 * (sphere.Radius * sphere.Radius * sphere.Radius)
}

type Shape interface {
	Volume() float64
}

func getTotalVolume(shapes ...Shape) float64 {
	var volume float64
	for _, s := range shapes {
		volume += s.Volume()
	}
	return volume
}

func Interfaces() {
	fmt.Println("Running `interfaces`...")
	shapes := []Shape{
		&Sphere{Radius: 33.0},
		&Sphere{Radius: 42.0},
		&Sphere{Radius: 54.0},
		&CubeC{Depth: 33.0, Width: 42.0, Height: 54.0},
	}
	fmt.Printf("Result of `getTotalVolume(shapes)` is `%.2f`.\n\n", getTotalVolume(shapes...))
}
