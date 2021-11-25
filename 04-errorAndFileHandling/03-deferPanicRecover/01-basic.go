package _010403_deferPanicRecover_

import "fmt"

/*
1.  The `panic` function is a built-in that causes the current function or
    method to begin panicking.
2.  You can think of "panicking" as a way of returning from a function.
3a. `panic` and `recover` are declared with the blank interface type:
     ```
     func panic(v interface{})
	    func recover() interface{}
     ```
3b. The blank interface type is somewhat like an `any` type. We can pass an
    argument value of any type when making a call to `panic`.
3c. Also, the value that is returned by a `recover` function is the value
    passed when the `panic` function call was made or if a real runtime
    error, the error condition
*/

func p(s string, i int) {
	// Here, if the condition is met, we enter a panic state. This should only
	// be used in exceptional cases where it's really not safe for the code to
	// continue its current flow.
	if i > 2 {
		panic(s)
	}
}

func r() {
	if err := recover(); err != nil {
		fmt.Println(err)
		fmt.Print("recovered from panic\n\n")
	}
}

// Defer is a process that places a function call onto a stack. Each deferred
// function is executed in reverse order when the calling function completes
// whether a panic is called or not.

func BasicDeferPanicRecover() {
	fmt.Println("Running `HandlingRuntimeErrors`...")
	defer r()
	p("runtime error: enter panic state", 3)
}
