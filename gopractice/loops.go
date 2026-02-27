package main

import "fmt"

func main() {
	// For loop
	// The most common loop in Go is the for loop. It has three components: the initialization, the condition, and the post statement.
	// The for loop loops through a block of code a specified number of times.
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// This counts from 0 to 100 in increments of 10.
	for j := 10; j <= 100; j += 10 {
		fmt.Println(j)
	}

	// Continue statement
	// The continue statement is used to skip the current iteration of a loop and move to the next iteration.
	for k := 0; k < 5; k++ {
		if k == 3 {
			continue // This will skip the rest of the loop when k is 3 and move to the next iteration
		}
		fmt.Println(k) // This will print only odd numbers
	}

	// Break statement
	for l := 0; l < 10; l++ {
		if l == 3 {
			break
		}
		fmt.Println(l)
	}

	// Nested loops
	// it is possible to place a loop inside another loop
	//Here, the "inner loop" will be executed one time for each iteration of the outer loop
	adj := [2]string{"red", "big"}
	fruits := [3]string{"apple", "banana", "cherry"}
	for x:=0; x < len(adj); x++ {
		for y:=0; y < len(fruits); y++ {
			fmt.Println(adj[x], fruits[y])
		}
	}

	// Range Keyword
	// The range keyword is used to easily iterate through elements of sn array, slice or map. It returns two values: the index and the value of the element at that index.
	// SYNTAX
	// for index, value := range array/slice/map {
	// 	// code to be executed
	// }
	fruits2 := [3]string{"apple", "banana", "cherry"}
	for index, value := range fruits2 {
		fmt.Printf("Index: %d, Value: %s\n", index, value)
	}
}
