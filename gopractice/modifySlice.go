package main

import "fmt"

func main() {
	// Access Elements of a Slice
	prices := []int{10, 20, 30}

	fmt.Println(prices[0])
	fmt.Println(prices[2])

	// Change Elements of a Slice
	prices1 := []int{10, 20, 30}
	prices1[2] = 50

	fmt.Println(prices1[0])
	fmt.Println(prices1[2])

	// Append Elements To a Slice

	// Syntax
	// slice_name = append(slice_name, element1, element2, ...)

	myslice := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("myslice = %v\n", myslice)
	fmt.Printf("length = %d\n", len(myslice))
	fmt.Printf("capacity = %d\n", cap(myslice))

	myslice = append(myslice, 20, 21)
	fmt.Printf("myslice = %v\n", myslice)
	fmt.Printf("length = %d\n", len(myslice))
	fmt.Printf("capacity = %d\n", cap(myslice))

	// Append One Slice To Another Slice

	// Syntax
	// slice3 = append(slice1, slice2...)

	myslice1 := []int{1, 2, 3}
	myslice2 := []int{4, 5, 6}
	myslice3 := append(myslice1, myslice2...)

	fmt.Printf("myslice3 = %v\n", myslice3)
	fmt.Printf("length = %d\n", len(myslice3))
	fmt.Printf("capacity = %d\n", cap(myslice3))

	// Change The Length of a Slice

	arr1 := []int{9, 10, 11, 12, 13, 14}
	myslice4 := arr1[1:3] // change length by re-slicing the array

	fmt.Printf("myslice4 = %v\n", myslice4)
	fmt.Printf("length = %d\n", len(myslice4))
	fmt.Printf("capacity = %d\n", cap(myslice4))

	// change length by appending items

	myslice4 = append(myslice4, 20, 21, 22, 23)

	fmt.Printf("myslice4 = %v\n", myslice4)
	fmt.Printf("length = %d\n", len(myslice4))
	fmt.Printf("capacity = %d\n", cap(myslice4))

	// Memory Efficiency with the use of copy function
	// The copy() function creates a new underlying array with only the required elements for the slice. This will reduce the memory used for the program.

	// Syntax
	// copy(dest, src)

	numbers := []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15}	// original slice
	fmt.Printf("numbers = %v\n", numbers)
	fmt.Printf("length = %d\n", len(numbers))
	fmt.Printf("capacity = %d\n", cap(numbers))

	// create copy with only needed numbers
	neededNums := numbers[:len(numbers)-10]
	numbersCopy := make([]int, len(neededNums))
	copy(numbersCopy, neededNums)

  fmt.Printf("numbersCopy = %v\n", numbersCopy)
  fmt.Printf("length = %d\n", len(numbersCopy))
  fmt.Printf("capacity = %d\n", cap(numbersCopy))

}
