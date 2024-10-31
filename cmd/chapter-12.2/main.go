package main

import "fmt"

// Create a function that launches two goroutines.
// Each goroutine writes 10 numbers to its own channel.
// Use a `for-select` loop to read from both channels, printing out the number and the goroutine that wrote the value.
// Make sure that your function exits after all values are read and that none of your goroutines leak.
func main() {
	fmt.Println("Hello world")
}
