package main

import (
	"fmt"
	"math/rand"
)

/*
Write a for loop that puts 100 random numbers between 0 and 100 into an int slice.
*/
func main() {
	randomInts := []int{}

	for i := 0; i < 100; i++ {
		randomInts = append(randomInts, rand.Intn(101))
	}

	fmt.Println(randomInts)
}
