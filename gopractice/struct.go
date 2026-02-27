package main

// A struct is used to create a collection of members of different data types, into a single variable.
// A struct can be useful for grouping data together to create records.
// SYNTAX
// type struct_name struct {
//     member1 datatype
//     member2 datatype
//     member3 datatype
// }

type person struct {
	name   string
	age    int
	job    string
	salary int
}

func main() {
	// Access struct members
	var pers1 person
	var pers2 person

	pers1.name = "John"
	pers1.age = 30
	pers1.job = "Software Engineer"
	pers1.salary = 50000

	pers2.name = "Jane"
	pers2.age = 28
	pers2.job = "Data Scientist"
	pers2.salary = 60000

	// println("Person 1:")
	// println("Name:", pers1.name)
	// println("Age:", pers1.age)
	// println("Job:", pers1.job)
	// println("Salary:", pers1.salary)

	// println("\nPerson 2:")
	// println("Name:", pers2.name)
	// println("Age:", pers2.age)
	// println("Job:", pers2.job)
	// println("Salary:", pers2.salary)

	// Access struct as function arguments
	printPerson(pers1)
	printPerson(pers2)

}

func printPerson(pers person) {
	println("Name:", pers.name)
	println("Age:", pers.age)
	println("Job:", pers.job)
	println("Salary:", pers.salary)
}
