package main

import (
	"fmt"
	"math/rand"
)

func main() {
	randomInts := []int{}

	for i := 0; i < 100; i++ {
		randomInts = append(randomInts, rand.Intn(101))
	}

	for _, num := range randomInts {
		switch {
		case num%2 == 0 && num%3 == 0:
			fmt.Println("Six!")
		case num%2 == 0:
			fmt.Println("Two!")
		case num%3 == 0:
			fmt.Println("Three!")
		default:
			fmt.Println("Never mind")
		}
	}
}
