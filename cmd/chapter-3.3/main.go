package main

import "fmt"

type Employee struct {
	firstName string
	lastName  string
	id        int
}

// Write a program that defines a struct called Employee with three fields: firstName, lastName, and id.
// The first two fields are of type string, and the last field (id) is of type int.
// Create three instances of this struct using whatever values you’d like.
// Initialize the first one using the struct literal style without names, the second using the struct literal style with names, and the third with a var declaration.
// Use dot notation to populate the fields in the third struct. Print out all three structs.
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
