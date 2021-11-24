package _010402_errorsPackage_

//`errors` package provides is with functiosn we can call to manipulate and
//manage errors.

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func IsFunction() {
	defer fmt.Println()
	fmt.Println("Running `IsFunction`...")
	// this is a composite if statement that both initialises a value and uses
	// it in the test condition
	if _, err := os.Open("doesnt-exist.txt"); err != nil {
		// Using the `Is` function is the recommended way to check for specific error
		// conditions as opposed to performing basic equality checks, etc.
		if errors.Is(err, os.ErrNotExist) {
			log.Println("file does not exist")
		} else {
			log.Println(err)
		}
		return
	}
	fmt.Println("file successfuly opened")
}
