package main

import "fmt"

func variables() {
	fmt.Println("\n\n_VARIABLES_")

	basics()
	naming()
	scopes()
}

func basics() {
	fmt.Println("\nBasics:")

	var a string = "Hello World"
	fmt.Println(a)

	var b string
	b = "Hello World"
	fmt.Println(b)

	var c string
	c = "first"
	fmt.Println(c)
	c = "second"
	fmt.Println(c)

	var d string
	d = "first "
	fmt.Println(d)
	d = d + "second"
	fmt.Println(d)

	var e string = "hello"
	var f string = "world"
	fmt.Println(e == f)

	var g string = "hello"
	var h string = "hello"
	fmt.Println(g == h)

	i := "Hello World"
	fmt.Println(i)

	var j = "Hello World"
	fmt.Println(j)

	k := 5
	fmt.Println(k)
}

func naming() {
	fmt.Println("\nNaming:")

	dogsName := "Max"
	fmt.Println("My dog's name is", dogsName)
}

func scopes() {
	outerX := "I'm outer X!"

	innerScope1 := func() {
		innerX := "I'm inner X!"

		fmt.Println(outerX)
		fmt.Println(innerX)
	}

	innerScope2 := func() {
		fmt.Println(outerX)
		//fmt.Println(innerX) //Doesn't work!
		fmt.Println("Inner X cannot be printed.")
	}

	innerScope1()
	innerScope2()
}
