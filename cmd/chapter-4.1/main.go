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

	fmt.Println(randomInts)
}
