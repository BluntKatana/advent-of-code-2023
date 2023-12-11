package day7

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Day7 struct{}

const (
	FIVE_OF_A_KIND  = 7
	FOUR_OF_A_KIND  = 6
	FULL_HOUSE      = 5
	THREE_OF_A_KIND = 4
	TWO_PAIR        = 3
	ONE_PAIR        = 2
	HIGH_CARD       = 1
)

type CamelCard struct {
	Card      string
	HandType  int
	BidAmount int
}

type CamelCardList []CamelCard

// Implement Len for CamelCardList
func (e CamelCardList) Len() int {
	return len(e)
}

var ORDER = "AKQJT98765432"

// Implement Less for CamelCardList (sort by hand type, then by card)
func (e CamelCardList) Less(i, j int) bool {
	// If the hand type is not equal, then we compare the individual characters from the beginning
	// of the string to determine which is greater. This is done in the following
	// order: A, K, Q, J, T, 9, 8, 7, 6, 5, 4, 3, or 2, where A is the highest
	// and 2 is the lowest.
	if e[i].HandType == e[j].HandType {
		for idx := range e[i].Card {
			var i_idx int = strings.Index(ORDER, string(e[i].Card[idx]))
			var j_idx int = strings.Index(ORDER, string(e[j].Card[idx]))

			if i_idx != j_idx {
				// if the characters are not equal, then we need to compare them
				// based on the order string
				return i_idx > j_idx
			}
		}

	}
	return e[i].HandType < e[j].HandType
}

// Implement Swap for CamelCardList
func (e CamelCardList) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func (d Day7) Part1(filename *string) string {
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
		var hand_type int = get_camel_card_type(card, get_counts(card))

		cards = append(cards, CamelCard{Card: card, HandType: hand_type, BidAmount: bid_amount})
	}

	// sort cards by hand type
	sort.Sort(CamelCardList(cards))

	// calculate the total
	var total int = 0
	for rank, card := range cards {
		total += card.BidAmount * (rank + 1)
	}

	fmt.Println(time.Since(start))
	return fmt.Sprint(total)
}

func get_camel_card_type(card string, counts map[string]int) int {

	// Check for five of a kind
	if len(counts) == 1 {
		return FIVE_OF_A_KIND
	} else if len(counts) == 2 && (counts[string(card[0])] == 4 || counts[string(card[1])] == 4) {
		// Check for four of a kind (either the first or second card will have 4)
		return FOUR_OF_A_KIND
	} else if len(counts) == 2 && (counts[string(card[0])] == 3 || counts[string(card[1])] == 3 || counts[string(card[2])] == 3) {
		// Check for full house (either the first, second or third card will have 3)
		return FULL_HOUSE
	} else if len(counts) == 3 && (counts[string(card[0])] == 3 || counts[string(card[1])] == 3 || counts[string(card[2])] == 3) {
		// Check for three of a kind (either the first, second or third card will have 3)
		return THREE_OF_A_KIND
	} else if len(counts) == 3 && (counts[string(card[0])] == 2 || counts[string(card[1])] == 2 || counts[string(card[2])] == 2) {
		// Check for two pair (either the first, second or third card will have 2)
		return TWO_PAIR
	} else if len(counts) == 4 {
		// Check for one pair
		return ONE_PAIR
	} else {
		// All thats left over is high card
		return HIGH_CARD
	}
}

func get_counts(card string) map[string]int {
	var counts map[string]int = map[string]int{}
	for _, card := range card {
		counts[string(card)]++
	}
	return counts
}
