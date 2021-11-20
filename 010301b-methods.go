package main

import "fmt"

type CubeType struct {
	d float64
	w float64
	h float64
}

func (c *CubeType) Volume() float64 {
	return c.d * c.w * c.h
}

func methods() {
	fmt.Println("Running `methods`...")
	c := &CubeType{d: 13.0, w: 13.0, h: 13.0}
	volumeOfCube := c.Volume()
	fmt.Printf("Result of `c.Volume()` is `%.2f`.\n\n", volumeOfCube)
}
