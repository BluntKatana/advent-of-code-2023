package main

import (
	"fmt"
	"os"
	"strings"
)

var mappedStrings = map[string]int{
	"one": 1, "two": 2, "three": 3,
	"four": 4, "five": 5, "six": 6,
	"seven": 7, "eight": 8, "nine": 9,
}

// go run main.go parse.go
func main() {
	content, _ := os.ReadFile("../input.txt")
	lines := strings.Fields(string(content))

	total := 0
	for _, line := range lines {
		var values []int

		potential_num := ""
		for _, char := range line {
			potential_num += string(char)

			// Check if a characters is a number
			if char >= 48 && char <= 57 {
				values = append(values, int(char)-'0')
			}

			// If the character could be one of the last characters of the
			// value numbers we check if the potential_num up until now is mapped
			// to one of the value numbers
			if strings.Contains("eorxnt", string(char)) {
				for value, number := range mappedStrings {
					if strings.HasSuffix(potential_num, value) {
						values = append(values, number)
					}
				}
			}
		}

		total += values[0]*10 + values[len(values)-1]
		potential_num = ""
	}

	// print result
	fmt.Println("Part 2:", total)
}
