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

func (s Pattern) FindVerticalReflection() (bool, int) {
	// Check for vertical reflections between the pivot points
	var checkLeft = 0
	var checkRight = 1

	for checkLeft < len(s[0])-1 && checkRight < len(s[0]) {
		// Check if the pivot points are the same
		var pivotLeft = checkLeft
		var pivotRight = checkRight
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
			return true, checkLeft
		}

		// Update the checkLeft and checkRight
		checkLeft++
		checkRight++
	}

	return false, -1
}

func (s Pattern) FindHorizontalReflection() (bool, int) {
	// Check for horizontal reflections between the pivot points
	var checkTop = 0
	var checkBottom = 1

	for checkTop < len(s)-1 && checkBottom < len(s) {
		// Check if the pivot points are the same
		var pivotTop = checkTop
		var pivotBottom = checkBottom
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
			return true, checkTop
		}

		// Update the checkLeft and checkRight
		checkTop++
		checkBottom++
	}

	return false, -1
}

func (d Day13) Part1(filename *string) string {
	var start = time.Now()

	var content, _ = os.ReadFile(*filename)
	var lines = strings.Split(string(content), "\n")

	var pattern []string
	var sum = 0

	for lineIdx, line := range lines {
		if line == "" || lineIdx == len(lines)-1 {
			// Check for reflections
			var patternObj = Pattern(pattern)
			fmt.Println()
			patternObj.String()
			fmt.Println()
			var foundRow, rowNum = patternObj.FindHorizontalReflection()
			var foundCol, colNum = patternObj.FindVerticalReflection()

			if foundRow {
				fmt.Println("FOUND ROW", rowNum+1)
				sum += (rowNum + 1) * 100
			}

			if foundCol {
				fmt.Println("FOUND COL", colNum+1)
				sum += colNum + 1
			}

			pattern = []string{}

			continue
		}

		pattern = append(pattern, line)
	}

	fmt.Println(time.Since(start))
	return fmt.Sprint(sum)
}
