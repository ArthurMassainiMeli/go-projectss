package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fileName := os.Args[1]
	fmt.Println(fileName)
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Error opening file %s\nError: %s\n", fileName, err.Error())
		os.Exit(1)
	}

	read, err := io.ReadAll(file)
	if err != nil {
		fmt.Printf("Error reading file %s\nError: %s\n", fileName, err.Error())
		os.Exit(1)
	}

	fmt.Println(string(read))
}
