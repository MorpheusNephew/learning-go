package main

import (
	"fmt"
	"io"
	"os"
)

func fileLen(filename string) (int, error) {
	file, err := os.Open(filename)

	if err != nil {
		return 0, err
	}

	defer file.Close()

	data := make([]byte, 2048)

	var totalCount int

	for {

		count, err := file.Read(data)

		if err != nil {
			if err == io.EOF {
				break
			}

			return 0, nil
		}

		totalCount += count
	}

	return totalCount, nil
}

// Write a function called `fileLen` that has an input parameter of type `string` and returns an `int` and an `error`.
// The function takes in a filename and returns the number of bytes in the file.
// If there is an error reading the file, return the error.
// Use `defer` to make sure the file is closed properly.
func main() {
	filename := os.Args[1]
	count, err := fileLen(filename)

	if err != nil {
		fmt.Printf("An error occurred \"%s\"\n", err.Error())
		os.Exit(1)
	}

	fmt.Println("The length of the file is", count)
}
