package main

import "fmt"

type Employee struct {
	firstName string
	lastName  string
	id        int
}

func main() {
	personOne := Employee{"First", "Last", 353}
	personTwo := Employee{firstName: "My First", lastName: "My Last", id: 3551}
	var personThree Employee

	personThree.firstName = "First Three"
	personThree.lastName = "Last Three"
	personThree.id = 128571

	fmt.Println(personOne)
	fmt.Println(personTwo)
	fmt.Println(personThree)
}
