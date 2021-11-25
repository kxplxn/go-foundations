package _010401_errorHandling_

import (
	"errors"
	"fmt"
	"math"
)

func volume(r float64) (float64, error) {
	if r < 0 {
		return 0, errors.New("volume calculation failed: radius negative")
	}
	return (4.0 / 3.0) * math.Pi * r * r * r, nil
}

func CustomErrors() {
	fmt.Println("Running `CustomErrors`...")
	radius := -1.0
	vol, err := volume(radius)
	if err != nil {
		fmt.Printf("ERROR: %v\n\n", err)
		return
	}
	fmt.Printf("Volume of sphere is %.2f.\n\n", vol)
}
