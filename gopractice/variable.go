package main

import "fmt"

func main() {
	var a string
	var b string
	var c string

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	var student1 string
	student1 = "Solomon"
	fmt.Println(student1)

	var a, b, c, d int = 1, 3, 5, 7

	fmt.Println(a, b, c, d)

	var a, b = 6, "Hello"
	c, d := 7, "World"

	fmt.Println(a, b, c, d)

	var (
		a int
		b int = 1
		c string = "hello"
	)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	// Naming a variable
	// Camel Case
	var (
		myVariableName = "John"

	// Pascal Case
		MyVariableName = "John"

	//Snake Case
		my_variable_name = "John"
	)

	fmt.Printf("my name is %v, %v, %v \n", myVariableName, MyVariableName, my_variable_name )
}
