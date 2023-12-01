package main

import "fmt"

var Testing bool = false

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
		for _, char := range line {
			if char >= 0 && char <= 90 {
				values = append(values, int(char)-'0')
			}
		}

		total += values[0]*10 + values[len(values)-1]
	}

	fmt.Println(total)
}

func combineFirstLast(values []int) int {
	first := values[0]
	last := values[len(values)-1]

	return first*10 + last
}

func grabValueFromString(line string) []int {
	var values []int
	for _, char := range line {
		if char >= 0 && char <= 90 {
			values = append(values, int(char)-'0')
		}
	}

	return values
}
