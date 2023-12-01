package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	content, _ := os.ReadFile("../input.txt")
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

	fmt.Println("Part 1:", total)
}
