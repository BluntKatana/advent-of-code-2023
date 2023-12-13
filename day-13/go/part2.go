package day13

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// Count the amount of different characters of two strings
func Diff(s1, s2 string) int {
	delta := 0
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			delta += 1
		}
	}

	return delta
}

func (s Pattern) FindVerticalReflectionWithSmudge() (bool, int) {
	var _, notAllowedLeft = s.FindVerticalReflection()

	// Check for vertical reflections between the pivot points
	var checkLeft = 0
	var checkRight = 1

	for checkLeft < len(s[0])-1 && checkRight < len(s[0]) {
		// Check if the pivot points are the same
		var pivotLeft = checkLeft
		var pivotRight = checkRight
		var areEqual = true
		var diffIsOne = false
		// Continue to check if the columns are the same
		for pivotLeft >= 0 && pivotRight < len(s[0]) {
			var leftCol = s.GetColumnString(pivotLeft)
			var rightCol = s.GetColumnString(pivotRight)

			var diff = Diff(leftCol, rightCol)

			if leftCol != rightCol && (diff > 1 || diffIsOne) {
				areEqual = false
				break
			}
			if diff == 1 {
				diffIsOne = true
			}
			pivotLeft--
			pivotRight++
		}

		if areEqual && checkLeft != notAllowedLeft {
			return true, checkLeft
		}

		// Update the checkLeft and checkRight
		checkLeft++
		checkRight++
	}

	return false, -1
}

func (s Pattern) FindHorizontalReflectionWithSmudge() (bool, int) {
	var _, notAllowedTop = s.FindHorizontalReflection()
	// Check for horizontal reflections between the pivot points
	var checkTop = 0
	var checkBottom = 1

	for checkTop < len(s)-1 && checkBottom < len(s) {
		// Check if the pivot points are the same
		var pivotTop = checkTop
		var pivotBottom = checkBottom
		var areEqual = true
		var diffIsOne = false

		// Continue to check if the rows are the same
		for pivotTop >= 0 && pivotBottom < len(s) {
			topRow := s[pivotTop]
			bottomRow := s[pivotBottom]
			var diff = Diff(topRow, bottomRow)

			if topRow != bottomRow && (diff > 1 || diffIsOne) {
				areEqual = false
				break
			}
			if diff == 1 {
				diffIsOne = true
			}
			pivotTop--
			pivotBottom++
		}

		if areEqual && checkTop != notAllowedTop {
			return true, checkTop
		}

		// Update the checkLeft and checkRight
		checkTop++
		checkBottom++
	}

	return false, -1
}
func (d Day13) Part2(filename *string) string {
	var start = time.Now()

	var content, _ = os.ReadFile(*filename)
	var lines = strings.Split(string(content), "\n")

	var pattern []string
	var sum = 0

	for lineIdx, line := range lines {
		if line == "" || lineIdx == len(lines)-1 {
			// Check for reflections
			var patternObj = Pattern(pattern)
			// fmt.Println("----- PATTERN -----")
			// patternObj.String()
			// fmt.Println()
			var foundRow, rowNum = patternObj.FindHorizontalReflectionWithSmudge()
			var foundCol, colNum = patternObj.FindVerticalReflectionWithSmudge()

			fmt.Println("----- PATTERN -----")
			if foundRow {
				fmt.Println("FOUND ROW", rowNum+1)
				// patternRow.String()
				sum += (rowNum + 1) * 100
			}

			if foundCol {
				fmt.Println("FOUND COL", colNum+1)
				// patternCol.String()
				sum += colNum + 1
			}

			if !foundRow && !foundCol {
				fmt.Println("NOT FOUND")
				patternObj.String()
			}

			pattern = []string{}

			continue
		}

		pattern = append(pattern, line)
	}

	fmt.Println(time.Since(start))
	return fmt.Sprint(sum)
}
