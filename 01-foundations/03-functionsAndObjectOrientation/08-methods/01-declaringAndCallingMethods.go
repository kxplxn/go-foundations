package _010308_methods_

import "fmt"

// Box is an exportable object.
type Box struct {
	D float64
	W float64
	H float64
}

// To write a method, we need to add a "receiver" between the `func` keyword and
// the function name. In this case, the receiver's type is a pointer to a `Box`
// object.
func (b *Box) volume() float64 {
	return b.D * b.W * b.H
}

func DeclaringAndCallingMethods() {
	fmt.Println("Running `DeclaringAndCallingMethods`...")
	b := Box{D: 5, W: 4, H: 3}
	fmt.Printf("Result of `b.volume()` is `%v`.\n\n", b.volume())
}
