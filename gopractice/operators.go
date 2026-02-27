package main

import "fmt"

func main() {
	var a = 15 + 25
	a += 5
	fmt.Println(a)

	sum1 := 23 +12
	sum2 := sum1 + 10
	sum3 := sum2 + sum1
	mod := sum3 % sum1

	fmt.Println(sum1)
	fmt.Println(sum2)
	fmt.Println(sum3)
	fmt.Println(mod)a

}
