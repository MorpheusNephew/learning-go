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

func main() {
	filename := os.Args[1]
	count, err := fileLen(filename)

	if err != nil {
		fmt.Printf("An error occurred \"%s\"\n", err.Error())
		os.Exit(1)
	}

	fmt.Println("The length of the file is", count)
}
