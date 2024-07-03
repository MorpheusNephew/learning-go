package main

import "fmt"

/*
Create a struct named `Person` with three fields: `FirstName` and `LastName` of type `string` and `Age` of type `int`.
Write a function called `MakePerson` that takes in `firstName`, `lastName`, and `age` and returns a `Person`.
Write a second function `MakePersonPointer` that takes in `firstName`, `lastName`, and `age` and returns a `*Person`.
Call both from `main`. Compile your program with `go build -gcflags="-m"`.
This both compiles your code and prints out which values escape to the heap. Are you surprised about what escapes?
*/
func main() {
	person := MakePerson("FirstPerson", "LastPerson", 42)
	personPointer := MakePersonPointer("FirstPersonPointer", "LastPersonPointer", 42)

	fmt.Println("Here is the first person", person)
	fmt.Println("Here is the first person pointer", personPointer)
}

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func MakePerson(firstName, lastName string, age int) Person {
	return Person{firstName, lastName, age}
}

func MakePersonPointer(firstName, lastName string, age int) *Person {
	return &Person{firstName, lastName, age}
}
