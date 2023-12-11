package day7

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type CamelCardListWithJoker []CamelCard

// Implement Len for CamelCardListWithJoker
func (e CamelCardListWithJoker) Len() int {
	return len(e)
}

var ORDER_WITH_JOKER = "AKQT98765432J"

// Implement Less for CamelCardListWithJoker (sort by hand type, then by card)
func (e CamelCardListWithJoker) Less(i, j int) bool {
	// If the hand type is not equal, then we compare the individual characters from the beginning
	// of the string to determine which is greater. This is done in the following
	// order: A, K, Q, T, 9, 8, 7, 6, 5, 4, 3, 2 or J, where A is the highest
	// and J is the lowest.
	if e[i].HandType == e[j].HandType {
		for idx := range e[i].Card {
			var i_idx int = strings.Index(ORDER_WITH_JOKER, string(e[i].Card[idx]))
			var j_idx int = strings.Index(ORDER_WITH_JOKER, string(e[j].Card[idx]))

			if i_idx != j_idx {
				// if the characters are not equal, then we need to compare them
				// based on the order string
				return i_idx > j_idx
			}
		}

	}
	return e[i].HandType < e[j].HandType
}

// Implement Swap for CamelCardListWithJoker
func (e CamelCardListWithJoker) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func (d Day7) Part2(filename *string) string {
	// start clock
	start := time.Now()

	content, _ := os.ReadFile(*filename)
	lines := strings.Split(string(content), "\n")

	var cards []CamelCard = []CamelCard{}

	// parse the lines
	for _, line := range lines {
		split_line := strings.Split(line, " ")
		var card = split_line[0]
		var bid_amount, _ = strconv.Atoi(split_line[1])
		var hand_type int = get_camel_card_type_with_joker(card)

		cards = append(cards, CamelCard{Card: card, HandType: hand_type, BidAmount: bid_amount})
	}

	// sort cards by hand type
	sort.Sort(CamelCardListWithJoker(cards))

	// calculate the total
	var total int = 0
	for rank, card := range cards {
		total += card.BidAmount * (rank + 1)
	}

	fmt.Println(time.Since(start))
	return fmt.Sprint(total)
}

func get_camel_card_type_with_joker(card string) int {
	var counts map[string]int = map[string]int{}
	for _, card := range card {
		counts[string(card)]++
	}

	// If there is no joker, then we can just return the hand type
	if counts["J"] == 0 {
		// If we have a joker, then we need to remove it from the counts
		return get_camel_card_type(card, get_counts(card))
	}

	// Go through the card and replace the joker with the most
	// common card in the hand
	var best_hand string = "J"
	var best_hand_count int = 0
	for card, count := range counts {
		if card == "J" {
			continue
		}

		if count > best_hand_count {
			best_hand = card
			best_hand_count = count
		}
	}

	// Replace the joker with the best hand
	card = strings.ReplaceAll(card, "J", best_hand)

	// Now we can just return the hand type
	return get_camel_card_type(card, get_counts(card))
}
