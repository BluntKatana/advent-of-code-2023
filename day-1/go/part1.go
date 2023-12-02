package day1

import (
	"fmt"
	"os"
	"strings"
)

// initialize day 1
type Day1 struct{}

func filename_part1(test_mode bool) string {
	if test_mode {
		return "./day-1/test_part1.txt"
	}
	return "./day-1/input.txt"
}

func (d Day1) Part1(test_mode bool) string {
	content, _ := os.ReadFile(filename_part1(test_mode))

	lines := strings.Fields(string(content))
	total := 0

	for _, line := range lines {
		var values []int
		for _, char := range line {
			if char >= 48 && char <= 57 {
				values = append(values, int(char)-'0')
			}
		}

		total += values[0]*10 + values[len(values)-1]
	}

	return fmt.Sprint(total)
}
