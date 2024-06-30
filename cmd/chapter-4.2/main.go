package main

import (
	"fmt"
	"math/rand"
)

/*
Loop over the slice you created in exercise

1. For each value in the slice, apply the following rules:

a. If the value is divisible by 2, print “Two!”

b. If the value is divisible by 3, print “Three!”

c. If the value is divisible by 2 and 3, print “Six!”. Don’t print anything else.

d. Otherwise, print “Never mind”.
*/
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
