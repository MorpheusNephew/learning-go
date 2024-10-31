package main

import "fmt"

// Create a function that launches three goroutines that communicate using a channel.
// The first two goroutines each write 10 numbers to the channel.
// The third goroutine reads all the numbers from the channel and prints them out.
// The function should exit when all values have been printed out.
// Make sure that none of the goroutines leak. You can create additional goroutines if needed.
func main() {
	fmt.Println("Hello world")
}
