package day5

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/golang-collections/collections/stack"
)

// A normal range tuple
type Range struct {
	First int
	Last  int
}

func (d Day5) Part2(filename *string) string {
	// start clock
	start := time.Now()

	// read files
	content, _ := os.ReadFile(*filename)
	lines := strings.Split(string(content), "\n")

	// create a map of maps from the input
	var steps = [7]string{"seed-to-soil", "soil-to-fertilizer", "fertilizer-to-water", "water-to-light", "light-to-temperature", "temperature-to-humidity", "humidity-to-location"}
	var mapping map[string][]MappedRange = make(map[string][]MappedRange)
	var seed_pairs []Range

	for _, step := range steps {
		mapping[step] = []MappedRange{}
	}

	// pre-processing the input
	var processing_step string
	// ranges and seeds
	for line_idx, line := range lines {
		// check if line is empty
		if line == "" {
			continue
		}

		// parse the line with the seeds
		if line_idx == 0 {
			var unparsed_seeds []string = strings.Split(strings.Split(line, ": ")[1], " ")

			// conver the seeds into pairs and add them to the map
			for i := 0; i < len(unparsed_seeds); i += 2 {
				var initial_seed, _ = strconv.Atoi(unparsed_seeds[i])
				var length, _ = strconv.Atoi(unparsed_seeds[i+1])

				// add the seed to the map
				seed_pairs = append(seed_pairs, Range{initial_seed, initial_seed + length})
			}
			continue
		}

		// check if line contains a step
		if strings.Contains(line, ":") {
			processing_step = strings.Split(line, " ")[0]
			continue
		}

		// check if line is empty
		var parsed_ranges []string = strings.Split(line, " ")
		var destination, _ = strconv.Atoi(parsed_ranges[0])
		var source, _ = strconv.Atoi(parsed_ranges[1])
		var length, _ = strconv.Atoi(parsed_ranges[2])

		// add the range to the map
		mapping[processing_step] = append(mapping[processing_step], MappedRange{destination, source, length})
	}

	var lowest_location int = -1

	// process the range of seeds until they reach the location
	for _, seed_pair := range seed_pairs {
		// pairs that have to be processed in the current step
		var pairs_to_be_processed = stack.New()
		pairs_to_be_processed.Push(seed_pair)

		// pairs that have been processed and are ready to be processed by the next step
		var pairs_gathered_during_step []Range = []Range{}

		// loop through each step
		for _, step := range steps {
			// for each map in the step we have to process the pairs, if they are mapped to the step
			// they are done processing and we can add them to the pairs_gathered_during_step
			// otherwise we add them to the pairs_to_be_processed and they can be mapped during a next map in the step
			for _, mapp := range mapping[step] {
				var store_for_next_map []Range = []Range{}

				// for every pairs that needs to be processed we have to go through the mapping phase of the step
				for pairs_to_be_processed.Len() > 0 {
					// pop the pair from the stack
					var pair_to_be_processed = pairs_to_be_processed.Pop().(Range)

					// loop through each range of the step
					// compare the pair to the range
					var mapped_ranges_step, not_mapped_ranges_step = mapp.CompareAndSplit(pair_to_be_processed)

					// add the outcome to the correct array
					pairs_gathered_during_step = append(pairs_gathered_during_step, mapped_ranges_step...)
					store_for_next_map = append(store_for_next_map, not_mapped_ranges_step...)
				}

				// add the pairs that are not mapped to the to be processed pairs
				// we do this outside of the loop because that way we are stuck in an infinite loop
				for _, pair_to_be_processed := range store_for_next_map {
					pairs_to_be_processed.Push(pair_to_be_processed)
				}
			}

			// add the pairs_gathered_during_step to the pairs_to_be_processed
			for _, pair_gathered_during_step := range pairs_gathered_during_step {
				pairs_to_be_processed.Push(pair_gathered_during_step)
			}

			// reset the pairs_gathered_during_step
			pairs_gathered_during_step = []Range{}
		}

		// check what seed is the lowest
		for pairs_to_be_processed.Len() > 0 {
			var location_pair = pairs_to_be_processed.Pop().(Range)

			if lowest_location == -1 || location_pair.First < lowest_location {
				lowest_location = location_pair.First
			}
		}

	}

	// print the time it took to run
	fmt.Println(time.Since(start))

	return fmt.Sprint(lowest_location)
}

// Compare a mapped range to a range and split the range into a tuple of mapped ranges, and not mapped range
func (mappedRange MappedRange) CompareAndSplit(r Range) ([]Range, []Range) {
	// left|inside|right
	if r.First < mappedRange.Source && r.Last >= mappedRange.Source+mappedRange.Length {
		var left_part = Range{r.First, mappedRange.Source - 1}
		// inside part gets mapped to destination
		var inside_part = Range{mappedRange.Destination, mappedRange.Destination + mappedRange.Length}
		var right_part = Range{mappedRange.Source + mappedRange.Length, r.Last}
		return []Range{inside_part}, []Range{left_part, right_part}

		// left|inside
	} else if r.First < mappedRange.Source && r.Last > mappedRange.Source && r.Last < mappedRange.Source+mappedRange.Length {
		var left_part = Range{r.First, mappedRange.Source - 1}
		// inside part gets mapped to destination
		// var inside_part = Range{mappedRange.Destination, mappedRange.Destination + mappedRange.Length}
		var inside_part = Range{mappedRange.Destination, r.Last + mappedRange.Destination - mappedRange.Source}
		return []Range{inside_part}, []Range{left_part}

		// inside|right
	} else if r.First >= mappedRange.Source && r.First < mappedRange.Source+mappedRange.Length && r.Last >= mappedRange.Source+mappedRange.Length {
		// inside part gets mapped to destination
		var inside_part = Range{r.First + mappedRange.Destination - mappedRange.Source, mappedRange.Destination + mappedRange.Length}
		var right_part = Range{mappedRange.Source + mappedRange.Length, r.Last}
		return []Range{inside_part}, []Range{right_part}

		// inside
	} else if r.First >= mappedRange.Source && r.First < mappedRange.Source+mappedRange.Length &&
		r.Last >= mappedRange.Source && r.Last < mappedRange.Source+mappedRange.Length {
		// inside part gets mapped to destination
		var inside_part = Range{r.First + mappedRange.Destination - mappedRange.Source, r.Last + mappedRange.Destination - mappedRange.Source}
		return []Range{inside_part}, []Range{}

		// outside
	} else {
		return []Range{}, []Range{r}
	}
}

/*
	WRITTEN DOWN EXAMPLE OF TEST INPUT FOR FIRST SEED-PAIR

	(79, 93)
	step 1 seed-to-soil
	- map: 98-99 to 50-51 => completely outside so do not do anything
	- map: 50-97 to 52-99 => completely inside (map 79-93 and add mapped to next step)
	no map left: add to next step
	(81, 95)
	step 2 soil-to-fertilizer
	- map: 15-51 to 0-36 => completely outside so do not do anything
	- map: 52-53 to 37-38 => completely outside so do not do anything
	- map: 0-14 to 39-53 => completely outside so do not do anything
	no map left: add to next step (81, 94 to next step)
	(81, 95)
	step 3 fertilizer-to-water
	- map: 53-60 to 49-56 => completely outside so do not do anything
	- map: 11-52 to 0-41 => completely outside so do not do anything
	- map: 0-6 to 42-48 => nothing
	- map 7-10 to 57-60 => nothing
	no map left: add to next step (81-94 to next step)
	(81, 95)
	step 4 water-to-light
	- map: 18-24 to 88-94 => nothing
	- map: 25-94 to 18-87 => partly inside (map 81-95 and add mapped to next step) (95-95 left over)
	no map left: add to next step (95-95 to next step)
	(76, 89), (95-95)
	step 5 light-to-temperature
	- map: 77-99 to 45-67 => partly inside (map 77-89 and mapped to next step) (76 left over)
	- map: ... => nothing
	- map: 64-76 to 68-80 => inside (map 76-76 and mapped to )

	fmt.Println(ranges_of_curr_step)
*/
