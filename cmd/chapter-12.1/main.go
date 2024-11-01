package main

import (
	"fmt"
	"math/rand"
	"sync"
)

// Create a function that launches three goroutines that communicate using a channel.
// The first two goroutines each write 10 numbers to the channel.
// The third goroutine reads all the numbers from the channel and prints them out.
// The function should exit when all values have been printed out.
// Make sure that none of the goroutines leak. You can create additional goroutines if needed.
func main() {
	firstWg := sync.WaitGroup{}
	firstWg.Add(2)

	buf := make(chan int)

	for i := 0; i < 2; i++ {
		go func() {
			defer firstWg.Done()

			for j := 0; j < 10; j++ {
				buf <- rand.Intn(100) + 1
			}
		}()
	}

	go func() {
		defer close(buf)

		firstWg.Wait()
	}()

	printWg := sync.WaitGroup{}
	printWg.Add(1)

	go func() {
		defer printWg.Done()

		for val := range buf {
			fmt.Println(val)
		}
	}()

	printWg.Wait()
}
