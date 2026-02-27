package main

import "fmt"

const PI = 3.14

func main() {
	fmt.Println(PI)

	// Typed constants are declared with defined type

	const A int = 1

	fmt.Println(A)

	// Untyped Constant

	const B = 1

	fmt.Println(A)

	// Multiple Constants Declaration

	const (
		A int = 1
		B = 3.14
		C = "Hi!"
	)
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
}

// When a constant is declared, it is not possible to change the value later
