package _010302_objectOrientation

import "fmt"

type CubeB struct {
	depth  float64
	width  float64
	height float64
}

func (c *CubeB) Volume() float64 {
	return c.depth * c.width * c.height
}

func BasicObjects() {
	fmt.Println("Running `basicObjects`...")
	cube := &CubeB{depth: 4, width: 4, height: 4}
	fmt.Printf("Type of `cube` is `%T`.\n", cube)
	fmt.Printf("Result of `cube.Volume()` is `%.2f`.\n\n", cube.Volume())
}
