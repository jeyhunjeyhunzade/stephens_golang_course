package main

import (
	"fmt"
	"io"
	"os"
)


func main() {
	// brute force solution:
	// fileName := os.Args[1]
	// readFile(fileName)

	f, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, f)
}

func readFile(fileName string) {
	content, err := os.ReadFile(fileName)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println(string(content))
}