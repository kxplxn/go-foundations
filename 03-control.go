package main

import "fmt"

func control() {
	fmt.Println("\n\n_CONTROL_")

	forLoop()
	ifStatement()
	switchStatement()
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

func switchStatement() {
	fmt.Println("\nSwitch:")

	runSwitch := func(number int) {
		switch number {
		case 0:
			fmt.Println("Zero")
		case 1:
			fmt.Println("One")
		case 2:
			fmt.Println("Two")
		case 3:
			fmt.Println("Three")
		case 4:
			fmt.Println("Four")
		case 5:
			fmt.Println("Five")
		default:
			fmt.Println("UNKNOWN NUMBER")
		}
	}

	runSwitch(3)
	runSwitch(123)
}
