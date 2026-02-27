package main

import "fmt"

func main() {
	// Maps are used to store data in key:value pairs. it does not accept duplicates.
	// SYNTAX
	//var a = map[KeyType]ValueType{Key1:Value1, Key2:Value2, ...}
	// b := map[KeyType]ValueType{Key1:Value1, Key2:Value2, ...}

	var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
	var b = map[string]int{"Oslo": 1, "Bergen": 2, "Trondheim": 3, "Stavanger": 4}

	fmt.Println(a)
	fmt.Println(b)

	// Create Maps using the make() function
	c := make(map[string]string)
	c["brand"] = "Ford"
	c["model"] = "Mustang"
	c["year"] = "1964"

	d := make(map[string]int)
	d["Oslo"] = 1
	d["Bergen"] = 2
	d["Trondheim"] = 3
	d["Stravenger"] = 4

	fmt.Println(c)
	fmt.Println(d)

	// Create an empty map
	var x = make(map[string]string)
	var y map[string]string

	fmt.Println(x == nil)
	fmt.Println(y == nil)

	// Access Map Elements
	// value = map_name[Key]
	fmt.Println(c["model"])

	// Update and Add Map Elements
	//map_name[key] = value
	c["color"] = "red"
	c["year"] = "1970"

	fmt.Println(c)

	// Remove element from map
	// delete(map_name, Key)

	delete(c, "year")
	fmt.Println(c)

	// Check For Specific Elements in a Map
	// val, ok := map_name[Key]

	var z = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964", "day": ""}

	val1, ok1 := z["brand"]
	val2, ok2 := z["skill"]
	val3, ok3 := z["day"]
	val4, _ := c["model"]

	fmt.Println(val1, ok1)
	fmt.Println(val2, ok2)
	fmt.Println(val3, ok3)
	fmt.Println(val4)

	// Maps are references

	var m = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
	n := m

	fmt.Println(m)
	fmt.Println(n)
	n["year"] = "1970"
	fmt.Println("After change to n:")

	fmt.Println(m)
	fmt.Println(n)

	// Iterate over maps

	for k, v := range b {
		fmt.Printf("%v : %v,", k, v)
	}

	// Iterate Over Maps in a Specific Order
	// Maps are unordered data structures. If you ned to interate over a map in a specific order, you must have a seperate data structure tht specifies that order.

	g := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}

  	var h []string             // defining the order
  	h = append(h, "one", "two", "three", "four")

  	for k, v := range g {        // loop with no order
    	fmt.Printf("%v : %v, ", k, v)
  	}

  	fmt.Println()

  	for _, element := range h {  // loop with the defined order
    	fmt.Printf("%v : %v, ", element, g[element])
  	}
}
	