package main

import "fmt"

func main() {
	// If statement specify a block of code to be executed if a condition is true
	if 20 > 18 {
		fmt.Println("20 is greater than 18")
	}
	x := 20
	y := 18
	if x > y {
		fmt.Println("x is greater than y")
	}

	// If else statement specify a block of code to be executed if the condition is false
	score := 10
	if score < 5 {
		fmt.Println("failed")
	} else {
		fmt.Println("passed")
	}

	// Else if statement specify a new condition if the first condition is false
	result := 85
	if result >= 75 {
		fmt.Println("Excellent")
	} else if result >= 50 {
		fmt.Println("tried")
	} else {
		fmt.Println("failed")
	}

	// Nested if statement
	num := 20
  	if num >= 10 {
    	fmt.Println("Num is more than 10.")
    	if num > 15 {
      	fmt.Println("Num is also more than 15.")
     	}
  	} else {
    fmt.Println("Num is less than 10.")
  }
}
