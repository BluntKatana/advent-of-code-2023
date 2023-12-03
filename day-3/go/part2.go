package day3

import (
	"fmt"
	"os"
	"strings"
)

func filename_part2(test_mode bool) string {
	if test_mode {
		return "./day-3/test_part2.txt"
	}
	return "./day-3/input.txt"
}

func (d Day3) Part2(test_mode bool) string {
	content, _ := os.ReadFile(filename_part2(test_mode))
	lines := strings.Split(string(content), "\n")

	for _, line := range lines {
		fmt.Println(line)
	}

	return "tbd"
}
