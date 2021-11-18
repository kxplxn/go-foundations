package main

import "fmt"

func control() {
	fmt.Println("\n\n_CONTROL_")

	forLong()
	forShort()
}

func forLong() {
	fmt.Println("\nFor (Long):")

	i := 1
	for i <= 10 {
		fmt.Printf("%v ", i)
		i++
	}
	fmt.Println()
}

func forShort() {
	fmt.Println("\nFor (Short):")
	for i := 1; i <= 10; i++ {
		fmt.Printf("%v ", i)
	}
}
