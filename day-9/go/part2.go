package day9

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// Calculates the first value in the history
func (h *History) CalculateFirstPrediction() int {
	// Keep a list of the last number for a list of diffs
	var first_nums []int = []int{}

	// Keep a list of the differences between the numbers
	var curr_diffs []int = h.Numbers
	var depth = 0
	for !AllZeros(curr_diffs) {
		var temp_diffs []int = []int{}

		for i := 0; i < len(curr_diffs)-1; i++ {
			// calc diffs from right to left
			var diff = curr_diffs[i] - curr_diffs[i+1]
			temp_diffs = append(temp_diffs, diff)
		}

		// Set the last diffs to the current diffs
		first_nums = append(first_nums, temp_diffs[0])

		// Reset the current diffs
		curr_diffs = temp_diffs
		depth++
	}

	// When we have the last_nums, we can calculate the prediction
	// by adding 0 to the last number, then adding the last number
	// to the last number and so on
	var prediction = 0

	for i := 0; i < len(first_nums); i++ {
		prediction += first_nums[i]
	}

	// Return the prediction + the first number
	return prediction + h.Numbers[0]
}

func (d Day9) Part2(filename *string) string {
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
		total += history.CalculateFirstPrediction()
	}

	fmt.Println(time.Since(start))
	return fmt.Sprint(total)
}
