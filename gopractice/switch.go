package main

import "fmt"

func main() {
	// Single-Case switch
	day := 4

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	}

	// Default Keyword
	month := 13

	switch month {
	case 1:
		fmt.Println("Jan")
	case 2:
		fmt.Println("Feb")
	case 3:
		fmt.Println("Mar")
	case 4:
		fmt.Println("April")
	case 5:
		fmt.Println("May")
	case 6:
		fmt.Println("June")
	case 7:
		fmt.Println("July")
	case 8:
		fmt.Println("Aug")
	case 9:
		fmt.Println("Sept")
	case 10:
		fmt.Println("Oct")
	case 11:
		fmt.Println("Nov")
	case 12:
		fmt.Println("Dec")
	default:
		fmt.Println("Not a Month")
	}

	// The Multi-case switch statement

	days := 5

	switch days {
	case 1, 3, 5:
		fmt.Println("Odd weekday")
	case 2, 4:
		fmt.Println("Even weekday")
	case 6, 7:
		fmt.Println("Weekend")
	default:
		fmt.Println("Invalid day of day number")
	}
}
