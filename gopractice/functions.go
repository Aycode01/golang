package main

import "fmt"

func main () {
	// Create a Slice With []datatype{values}

	myslice1 := []int{}
	fmt.Println(myslice1)
	fmt.Println(len(myslice1))
	fmt.Println(cap(myslice1))

	myslice2 := []string{"Go", "Slices", "Are", "Powerful"}
	fmt.Println(myslice2)
	fmt.Println(len(myslice2))
	fmt.Println(cap(myslice2))

	// Create a Slice From an Array

	// Syntax
	// var myarray = [length]datatype{value}	// An Array
	// myslice := myarray[start:end]			// A slice made from the array

	arr1 := [6]int{10, 11, 12, 13, 14, 15}
	myslice := arr1[1:5]

	fmt.Printf("myslice = %v\n", myslice)
	fmt.Printf("length = %d\n", len(myslice))
	fmt.Printf("capacity = %d\n", cap(myslice))

	// Create a Slice With The make() Function

	//Syntax
	// slice_name := make([]type, length, capacity)
	myslice3 := make([]int, 5, 10)

	fmt.Printf("myslice2 = %v\n", myslice3)
	fmt.Printf("length = %d\n", len(myslice3))
	fmt.Printf("capacity = %d\n", cap(myslice3))

	// With omitted capacity

	myslice4 := make([]int, 5)

	fmt.Printf("myslice2 = %v\n", myslice4)
	fmt.Printf("length = %d\n", len(myslice4))
	fmt.Printf("capacity = %d\n", cap(myslice4))
}
