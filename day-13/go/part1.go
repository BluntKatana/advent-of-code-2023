package day13

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type Day13 struct{}

type Pattern []string

func (s Pattern) String() {
	for _, line := range s {
		fmt.Println(line)
	}
}

func (s Pattern) GetColumnString(col int) string {
	var column = ""

	for _, line := range s {
		column += string(line[col])
	}

	return column
}

func (s Pattern) FindVerticalReflection() int {
	for pivot := 0; pivot < len(s[0])-1; pivot++ {
		// Check if the pivot points are the same
		var pivotLeft, pivotRight = pivot, pivot + 1
		var areEqual = true

		// Continue to check if the columns are the same
		for pivotLeft >= 0 && pivotRight < len(s[0]) {
			var leftCol = s.GetColumnString(pivotLeft)
			var rightCol = s.GetColumnString(pivotRight)

			if leftCol != rightCol {
				areEqual = false
				break
			} else {
				pivotLeft--
				pivotRight++
			}
		}

		if areEqual {
			return pivot + 1
		}
	}

	return 0
}

func (s Pattern) FindHorizontalReflection() int {
	for pivot := 0; pivot < len(s)-1; pivot++ {
		var pivotTop, pivotBottom = pivot, pivot + 1
		var areEqual = true

		// Continue to check if the rows are the same
		for pivotTop >= 0 && pivotBottom < len(s) {
			if s[pivotTop] != s[pivotBottom] {
				areEqual = false
				break
			} else {
				pivotTop--
				pivotBottom++
			}
		}

		if areEqual {
			return pivot + 1
		}
	}

	return 0
}

func (d Day13) Part1(filename *string) string {
	var start = time.Now()

	var content, _ = os.ReadFile(*filename)
	var lines = strings.Split(string(content), "\n")

	var pattern []string
	var sum = 0

	for lineIdx, line := range lines {
		if line == "" || lineIdx == len(lines)-1 {
			var patternObj = Pattern(pattern)

			var row = patternObj.FindHorizontalReflection()
			var col = patternObj.FindVerticalReflection()

			sum += (row * 100) + col

			pattern = []string{}

			continue
		}

		pattern = append(pattern, line)
	}

	fmt.Println(time.Since(start))
	return fmt.Sprint(sum)
}
