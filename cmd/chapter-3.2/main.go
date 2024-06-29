package main

import "fmt"

func main() {
	message := "Hi 👩‍🦳 and 🧔"

	something := message[3]
	somethingElse := rune(message[3])
	somethingElseCool := string(message[3])

	fmt.Println(something, somethingElse, somethingElseCool)
}
