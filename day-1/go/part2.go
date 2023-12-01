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

func main() {
	content, _ := os.ReadFile("../test_part2.txt")
	lines := strings.Fields(string(content))

	total := 0
	for _, line := range lines {
		var values []int

		potential_num := ""
		for idx, char := range line {
			potential_num += string(char)

			// Check if a characters is a number
			if char >= 48 && char <= 57 {
				values = append(values, int(char)-'0')
			}

			// Start checking suffix when
			// - the index is great than 1 (all numbers are at least 3 characters long)
			// - the character is not the first character
			if idx+1 >= 3 && strings.Contains("eorxnt", string(char)) {
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
