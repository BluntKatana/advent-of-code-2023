package day1

import (
	"fmt"
	"os"
	"strings"
)

// initialize day 1
type Day1 struct{}

func (d Day1) Part1(filename *string) string {
	content, _ := os.ReadFile(*filename)

	lines := strings.Fields(string(content))
	total := 0

	for _, line := range lines {
		var values []int
		for _, char := range line {
			if char >= 48 && char <= 57 {
				values = append(values, int(char)-'0')
			}
		}

		if len := len(values) - 1; len > 0 {
			total += values[0]*10 + values[len]
		}
	}

	return fmt.Sprint(total)
}
