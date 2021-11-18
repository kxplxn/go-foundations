package main

import "fmt"

func numbers() {
	fmt.Println("\nNumbers:")
	fmt.Println("1 + 1 =", 1+1)
	fmt.Println("1.0 + 1.0 =", 1.0+1.0)
}

func strings() {
	fmt.Println("\nStrings:")
	fmt.Println("len(\"Hello World\") =", len("Hello World"))
	fmt.Println("\"HelloWorld\"[1] =", "HelloWorld"[1])
}
