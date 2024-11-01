package main

import (
	"fmt"
	"math/rand"
)

// Create a function that launches two goroutines.
// Each goroutine writes 10 numbers to its own channel.
// Use a `for-select` loop to read from both channels, printing out the number and the goroutine that wrote the value.
// Make sure that your function exits after all values are read and that none of your goroutines leak.
func main() {
	ch1 := nRandomValues(10)
	ch2 := nRandomValues(10)

	for count := 0; count < 2; {
		select {
		case v, ok := <-ch1:
			if !ok {
				ch1 = nil
				count++
				continue
			}
			fmt.Println(v, "ch1")
		case v, ok := <-ch2:
			if !ok {
				ch2 = nil
				count++
				continue
			}
			fmt.Println(v, "ch2")
		}
	}
}

func nRandomValues(n int) <-chan int {
	ch := make(chan int, n)

	go func() {
		defer close(ch)

		for i := 0; i < n; i++ {
			ch <- rand.Intn(100) + 1
		}
	}()

	return ch
}
