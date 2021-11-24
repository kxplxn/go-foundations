package _010403_deferPanicRecover_

import "fmt"

func p2(i int) {
	x := []int{1, 2, 3}
	x[i] = 11
}

func r2() {
	if err := recover(); err != nil {
		fmt.Println(err)
		fmt.Println("recovered from panic")
	}
}

func HandlingRuntimeErrors() {
	fmt.Println("Running `HandlingRuntimeErrors`...")
	defer r2()
	p2(3)
}
