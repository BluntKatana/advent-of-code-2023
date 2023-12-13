package day13

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// Count the amount of different characters of two strings
func DiffChars(s1, s2 string) int {
	delta := 0
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			delta += 1
		}
	}

	return delta
}

func (s Pattern) FindVerticalReflectionWithSmudge() int {
	var pivotNotAllowed = s.FindVerticalReflection() - 1

	for pivot := 0; pivot < len(s[0])-1; pivot++ {
		// Check if the pivot points are the same
		var pivotLeft, pivotRight = pivot, pivot + 1
		var areDifferent, hasSmudge = false, false

		// Continue to check if the columns are the same
		for pivotLeft >= 0 && pivotRight <= len(s[0])-1 {
			var leftCol = s.GetColumnString(pivotLeft)
			var rightCol = s.GetColumnString(pivotRight)

			var diff = DiffChars(leftCol, rightCol)

			if leftCol != rightCol && (diff > 1 || hasSmudge) {
				areDifferent = true
				break
			}
			if diff == 1 {
				hasSmudge = true
			}
			pivotLeft--
			pivotRight++
		}

		if !areDifferent && pivot != pivotNotAllowed {
			return pivot + 1
		}
	}

	return 0
}

func (s Pattern) FindHorizontalReflectionWithSmudge() int {
	var pivotNotAllowed = s.FindHorizontalReflection() - 1

	for pivot := 0; pivot < len(s)-1; pivot++ {
		// Check if the pivot points are the same
		var pivotTop, pivotBottom = pivot, pivot + 1
		var areDifferent, hasSmudge = false, false

		// Continue to check if the rows are the same
		for pivotTop >= 0 && pivotBottom <= len(s)-1 {
			topRow := s[pivotTop]
			bottomRow := s[pivotBottom]
			var diff = DiffChars(topRow, bottomRow)

			if topRow != bottomRow && (diff > 1 || hasSmudge) {
				areDifferent = true
				break
			}
			if diff == 1 {
				hasSmudge = true
			}
			pivotTop--
			pivotBottom++
		}

		if !areDifferent && pivot != pivotNotAllowed {
			return pivot + 1
		}
	}

	return 0
}
func (d Day13) Part2(filename *string) string {
	var start = time.Now()

	var content, _ = os.ReadFile(*filename)
	var lines = strings.Split(string(content), "\n")

	var pattern []string
	var sum = 0

	for lineIdx, line := range lines {
		if line == "" || lineIdx == len(lines)-1 {
			var patternObj = Pattern(pattern)
			var row = patternObj.FindHorizontalReflectionWithSmudge()
			var col = patternObj.FindVerticalReflectionWithSmudge()

			sum += (row * 100) + col
			pattern = []string{}

			continue
		}

		pattern = append(pattern, line)
	}

	fmt.Println(time.Since(start))
	return fmt.Sprint(sum)
}
