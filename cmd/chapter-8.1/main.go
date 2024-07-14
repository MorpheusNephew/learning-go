package main

import "fmt"

// Write a generic function that doubles the value of any integer or float that’s passed in to it. Define any needed generic interfaces.
func main() {

	fmt.Println(Double(2))
	fmt.Println(Double(4.8))
}

type Number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | uintptr |
		float32 | float64
}

func Double[T Number](n T) T {
	return n * 2
}
