package main

import "fmt"

var Testing bool = true

// go run main.go parse.go
func main() {
	var lines []string
	if Testing {
		lines = Parse("input_test.txt")
	} else {
		lines = Parse("input.txt")
	}

	fmt.Println("lines: ", lines)
}
