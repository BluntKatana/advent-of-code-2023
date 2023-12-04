package day4

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// filter out a string from an array of strings
func filter_with(arr *[]string, str string) {
	for i := 0; i < len(*arr)-1; i++ {
		if (*arr)[i] == str {
			*arr = append((*arr)[:i], (*arr)[i+1:]...)
		}
	}
}

func (d Day4) Part2(filename *string) string {
	content, _ := os.ReadFile(*filename)
	lines := strings.Split(string(content), "\n")

	// clock the time
	start := time.Now()

	var cards map[int]int = map[int]int{}

	for card_num, line := range lines {
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
				count_winners++
			}
		}

		// check if we won!
		if count_winners > 0 {
			for i := 1; i <= count_winners; i++ {
				cards[card_num+i] += 1 + cards[card_num]
			}
		}

		// add the card to self
		cards[card_num] += 1
	}

	// count the total number of scratch cards
	var total_scratch_cards = 0

	for _, scratch_count := range cards {
		total_scratch_cards += scratch_count
	}

	// print the time it took to run
	fmt.Println(time.Since(start))

	return fmt.Sprint(total_scratch_cards)
}
