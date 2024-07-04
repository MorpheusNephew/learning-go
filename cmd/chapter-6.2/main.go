package main

import "fmt"

/*
Write two functions. The `UpdateSlice` function takes in a `[]string` and a `string`.
It sets the last position in the passed-in slice to the passed-in `string`.
At the end of `UpdateSlice`, print the slice after making the change.
The `GrowSlice` function also takes in a `[]string` and a `string`. It appends the `string` onto the slice.
At the end of `GrowSlice`, print the slice after making the change.
Call these functions from `main`.
Print out the slice before each function is called and after each function is called.
Do you understand why some changes are visible in `main` and why some changes are not?
*/
func main() {
	sliceForUpdateSlice := []string{"Hello", "slice", "for", "UpdateSlice", "how", "are", "you?"}
	sliceForGrowSlice := []string{"Hello", "slice", "for", "GrowSlice", "how", "are", "you?"}

	fmt.Println("Before UpdateSlice", sliceForUpdateSlice)
	UpdateSlice(sliceForUpdateSlice, "me")
	fmt.Println("After UpdateSlice", sliceForUpdateSlice)

	fmt.Println("Before GrowSlice", sliceForGrowSlice)
	GrowSlice(sliceForGrowSlice, "me")
	fmt.Println("After GrowSlice", sliceForGrowSlice)
}

/*
UpdateSlice updates the memory location of the last position in the slice which also updates the original slice
*/
func UpdateSlice(inputSlice []string, update string) {
	inputLen := len(inputSlice)
	inputSlice[inputLen-1] = update
	fmt.Println("UpdateSlice", inputSlice)
}

/*
GrowSlice does not update the original slice because the input parameter which is passed by value is a copy and that copy is given an entirely new memory address
*/
func GrowSlice(inputSlice []string, update string) {
	inputSlice = append(inputSlice, update)
	fmt.Println("GrowSlice", inputSlice)
}
