package main

import "fmt"

func main() {
	// With var or with := sign to a defined lengths
	var arr1 = [3]int{1, 2, 3}
	arr2 := [5]int{4, 5, 6, 7, 8}

	fmt.Println(arr1)
	fmt.Println(arr2)

	// with inferred lengths

	var arr3 = [...]int{1, 2, 3}
	arr4 := [...]int{4, 5, 6, 7, 8}

	fmt.Println(arr3)
	fmt.Println(arr4)

	// For strings

	var cars = [4]string{"volvo", "BMW", "Ford", "Mazda"}
	names := [...]string{"Ayo", "Bayo", "Dayo", "Sayo", "Tayo"}

	fmt.Println(cars)
	fmt.Println(names)

	// Access elements of an array

	fmt.Println(cars[0])
	fmt.Println(names[4])

	// Change element of an array

	names[4] = "Bami"
	fmt.Println(names)

	// Initialize Only Specific Elements
	//  The array above has 5 elements.
	arr5 := [5]int{1: 10, 2: 40}
	fmt.Println(arr5)
	// 1:10 means: assign 10 to array index 1 (second element).
	// 2:40 means: assign 40 to array index 2 (third element).

	// Length of an array

	fmt.Println(len(arr1))
	fmt.Println(len(arr2))
	fmt.Println(len(arr3))
	fmt.Println(len(arr4))
	fmt.Println(len(arr5))

}
