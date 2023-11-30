package main

import "fmt"

var Testing bool = true

// go run main.go parse.go
func main() {
	var lines []string
	if Testing {
		lines = parse("input_test.txt")
	} else {
		lines = parse("input.txt")
	}

	fmt.Println("lines: ", lines)
}
