package main

import "fmt"

// Write a program that defines a variable named greetings of type slice of strings with the following values:
// "Hello", "Hola", "नमस्कार", "こんにちは", and "Привіт".
// Create a subslice containing the first two values;
// a second subslice with the second, third, and fourth values;
// and a third subslice with the fourth and fifth values. Print out all four slices.
func main() {
	greetings := []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}

	firstTwo := greetings[:2]

	nextThree := greetings[2:]

	lastTwo := greetings[3:]

	fmt.Println(greetings)
	fmt.Println(firstTwo)
	fmt.Println(nextThree)
	fmt.Println(lastTwo)
}
