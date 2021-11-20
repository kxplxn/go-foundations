package main

import "fmt"

func avg(x []float64) float64 {
	total := 0.0
	for _, val := range x {
		total += val
	}
	return total / float64(len(x))
}

func functions() {
	fmt.Println("Running `functions`...")
	avgRes := avg([]float64{43.0, 34.0, 21.0, 52.0})
	fmt.Printf("Result of `avg` is `%.2f`.\n\n", avgRes)
}
