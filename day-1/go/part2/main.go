package main

import (
	"fmt"
	"strings"
)

type ValueInt struct {
	Value string
	Int   int
}

var Testing bool = false

var mappedStrings = []ValueInt{
	{"one", 1},
	{"two", 2},
	{"three", 3},
	{"four", 4},
	{"five", 5},
	{"six", 6},
	{"seven", 7},
	{"eight", 8},
	{"nine", 9},
}

// go run main.go parse.go
func main() {
	var lines []string
	if Testing {
		lines = Parse("input_test.txt")
	} else {
		lines = Parse("input.txt")
	}

	total := 0
	for _, line := range lines {
		var values []int

		potential_num := ""
		for _, char := range line {
			potential_num += string(char)
			if char >= 0 && char <= 90 {
				values = append(values, int(char)-'0')
			}
			// Check if potential_num mapped to one of the mappedString values
			for _, valueints := range mappedStrings {
				if strings.HasSuffix(potential_num, valueints.Value) {
					values = append(values, valueints.Int)
				}
			}
		}

		potential_num = ""
		total += values[0]*10 + values[len(values)-1]
	}

	fmt.Println(total)
}
