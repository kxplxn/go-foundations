package main

import "fmt"

func control() {
	fmt.Println("\n\n_CONTROL_")

	forLoop()
	ifStatement()
}

func forLoop() {
	fmt.Println("\nFor Loop:")

	fmt.Print("Long: ")
	i := 1
	for i <= 10 {
		fmt.Printf("%v ", i)
		i++
	}
	fmt.Println()

	fmt.Print("Short: ")
	for i := 1; i <= 10; i++ {
		fmt.Printf("%v ", i)
	}
	fmt.Println()
}

func ifStatement() {
	fmt.Println("\nIf:")

	runIf := func(number int) {
		if number%2 == 0 {
			fmt.Printf("%d is EVEN.\n", number)
		} else {
			fmt.Printf("%d is ODD.\n", number)
		}
	}

	runIf(4)
	runIf(5)
}
