package main

import "fmt"

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
