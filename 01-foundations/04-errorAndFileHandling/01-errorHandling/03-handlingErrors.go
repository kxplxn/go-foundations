package _0104_errorAndFileHandling_

import (
	"fmt"
	"log"
	"os"
)

func HandlingErrors() {
	fmt.Println("Running `HandlingErrors`...")

	f, err := os.Open("myFile.txt")
	// `defer` delays the execution of f.Close until we return from this function (HandlingErrors)
	defer f.Close()

	if err != nil {
		// printing via `log` package is safer than `fmt` and it also adds timing information
		log.Println(err)
		return
	}

	fmt.Println("file is successfuly opened:", f.Name())
}
