package _010301_functionsAndMethods

import "fmt"

type CubeA struct {
	d float64
	w float64
	h float64
}

func (c *CubeA) Volume() float64 {
	return c.d * c.w * c.h
}

func Methods() {
	fmt.Println("Running `methods`...")
	c := &CubeA{d: 13.0, w: 13.0, h: 13.0}
	volumeOfCube := c.Volume()
	fmt.Printf("Result of `c.Volume()` is `%.2f`.\n\n", volumeOfCube)
}
