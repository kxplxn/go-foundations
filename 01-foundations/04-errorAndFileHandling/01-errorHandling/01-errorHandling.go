package _0104_errorAndFileHandling_

import (
	"fmt"
	"os"
)

/*
1. In Go, errors are simply values — normal Go values.
2. They are represented with the built-in `error` type.
3. Just like any value, we can store errors in variables, return them from
	  functions, and pass them as parameters to functions.
4. We examine the error value that a function returns to determine if an
   error has occured.

================================================================================

"Defer, Panic, Recover" should only be resorted in situations where unexpected,
unrecoverable failures occur — truly exceptional situtations.

1. DEFER is a process that places a function call onto a stack.
2. PANIC terminates the natural flow of execution.
3. RECOVER provides us with a mechanism that allows us to recover the execution from terminating.
*/

func ErrorHandling() {
	fmt.Println("Running `ErrorHandling`...")

	// os.Open signature: `func Open(name string) (file *File, err error)`
	// -> Errors are always the last value returned from a function.
	file, err := os.Open("myFile.txt")

	// This is where the error is handled.
	if err != nil {
		fmt.Printf("ERROR: %v\n\n", err.Error())
		return
	}

	fmt.Printf("File `%s` is successfuly opened.\n\n", file.Name())
}
