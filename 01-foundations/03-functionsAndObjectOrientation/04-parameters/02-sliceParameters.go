package _010304_parameters

import "fmt"

func addB(numbers []int) int {
	total := 0
	for _, n := range numbers {
		total += n
	}
	return total
}

func SliceParameters() {
	fmt.Println("Running `sliceParameters`...")
	numbers := []int{20, 10, 5}
	fmt.Printf("Result of `addB(numbers)` is `%v`.\n\n", addB(numbers))
}

// Slice parameters are passed as a value rather than a reference...
