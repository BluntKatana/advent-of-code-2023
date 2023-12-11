package day9

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

type Day9 struct{}

type History struct {
	Numbers []int
}

// Calculate the last value in the history
func (h *History) CalculateLastPrediction() int {
	// Keep a list of the last number for a list of diffs
	var last_nums []int = []int{}

	// Keep a list of the differences between the numbers
	var curr_diffs []int = h.Numbers
	var depth = 0
	for !slices.ContainsFunc(curr_diffs, func(num int) bool { return num != 0 }) {
		var temp_diffs []int = []int{}

		for i := 0; i < len(curr_diffs)-1; i++ {
			// calc diffs from left to righ
			var diff = curr_diffs[i+1] - curr_diffs[i]
			temp_diffs = append(temp_diffs, diff)
		}

		// Set the last diffs to the current diffs
		last_nums = append(last_nums, temp_diffs[len(temp_diffs)-1])

		// Reset the current diffs
		curr_diffs = temp_diffs
		depth++
	}

	// When we have the last_nums, we can calculate the prediction
	// by adding 0 to the last number, then adding the last number
	// to the last number and so on
	var prediction = 0

	for i := 0; i < len(last_nums); i++ {
		prediction += last_nums[i]
	}

	// Return the prediction + the last number
	return prediction + h.Numbers[len(h.Numbers)-1]
}

func (d Day9) Part1(filename *string) string {
	var start = time.Now()

	// Parse input
	var content, _ = os.ReadFile(*filename)
	var lines = strings.Split(string(content), "\n")

	var total = 0

	// Go over each line to calculate the total
	for _, line := range lines {
		var numbers = strings.Split(line, " ")
		var history = History{Numbers: []int{}}

		// Go over each number in the line
		for _, number := range numbers {
			var num, _ = strconv.Atoi(number)

			// If the history is not full, add the number to the history
			history.Numbers = append(history.Numbers, num)

		}

		// Calculate the prediction
		total += history.CalculateLastPrediction()
	}

	fmt.Println(time.Since(start))
	return fmt.Sprint(total)
}
