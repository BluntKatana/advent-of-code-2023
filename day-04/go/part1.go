package day4

import (
	"fmt"
	"os"
	"strings"
)

type Day4 struct{}

func (d Day4) Part1(filename *string) string {
	content, _ := os.ReadFile(*filename)
	lines := strings.Split(string(content), "\n")

	var scratch_total = 0

	for _, line := range lines {
		var all_numbers = strings.Split(line, ": ")
		var split_numbers = strings.Split(all_numbers[1], " | ")
		var winning_numbers = strings.Split(strings.Trim(split_numbers[0], " "), " ")
		var own_numbers = strings.Split(strings.Trim(split_numbers[1], " "), " ")

		// filter out the spaces for the numbers
		filter_with(&winning_numbers, "")
		filter_with(&own_numbers, "")

		var count_winners = 0

		// create a map to hold the winning numbers
		winners_num_map := map[string]bool{}
		for _, winning_number := range winning_numbers {
			winners_num_map[winning_number] = true
		}

		// loop through the own numbers and see if they match
		for _, own_number := range own_numbers {
			if winners_num_map[own_number] {
				if count_winners == 0 {
					count_winners++
				} else {
					count_winners *= 2
				}
			}
		}

		scratch_total += count_winners
	}

	return fmt.Sprint(scratch_total)
}
