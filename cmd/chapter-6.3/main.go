package main

import "fmt"

// Write a program that builds a `[]Person` with 10,000,000 entries (they could all be the same names and ages).
// See how long it takes to run. Change the value of `GOGC` and see how that affects the time it takes for the program to complete.
// Set the environment variable `GODEBUG=gctrace=1` to see when garbage collections happen and see how changing `GOGC` changes the number of garbage collections.
// What happens if you create the slice with a capacity of 10,000,000?
func main() {
	people := []Person{}

	for i := 0; i < 10000000; i++ {
		firstName := fmt.Sprintf("FirstName%d", i)
		lastName := fmt.Sprintf("LastName%d", i)
		age := i + 1
		people = append(people, MakePerson(firstName, lastName, age))
	}

	numOfPeople := len(people)

	fmt.Println("Here's the number of people we have", numOfPeople)
}

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func MakePerson(firstName, lastName string, age int) Person {
	return Person{firstName, lastName, age}
}
