package main

import "fmt"

/*
Write a program that defines a string variable called message with the value "Hi  and " and prints the fourth rune in it as a character, not a number.
*/
func main() {
	message := "Hi 👩‍🦳 and 🧔"

	something := message[3]
	somethingElse := rune(message[3])
	somethingElseCool := string(message[3])

	fmt.Println(something, somethingElse, somethingElseCool)
}
