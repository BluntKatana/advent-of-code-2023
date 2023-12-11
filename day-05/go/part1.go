package day5

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Day5 struct{}

type MappedRange struct {
	Destination int
	Source      int
	Length      int
}

func (d Day5) Part1(filename *string) string {
	content, _ := os.ReadFile(*filename)
	lines := strings.Split(string(content), "\n")

	// create a map of maps from the input
	var steps = [7]string{"seed-to-soil", "soil-to-fertilizer", "fertilizer-to-water", "water-to-light", "light-to-temperature", "temperature-to-humidity", "humidity-to-location"}
	var ranges map[string][]MappedRange = make(map[string][]MappedRange)
	var seeds []int

	for _, step := range steps {
		ranges[step] = []MappedRange{}
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
			for _, seed := range unparsed_seeds {
				var parsed_seed int
				fmt.Sscanf(seed, "%d", &parsed_seed)
				seeds = append(seeds, parsed_seed)
			}
			continue
		}

		// check if line contains a step
		if strings.Contains(line, ":") {
			processing_step = strings.Split(line, " ")[0]
			continue
		}

		// parse the line with the ranges

		// check if line is empty
		var parsed_ranges []string = strings.Split(line, " ")
		var destination, _ = strconv.Atoi(parsed_ranges[0])
		var source, _ = strconv.Atoi(parsed_ranges[1])
		var length, _ = strconv.Atoi(parsed_ranges[2])

		// add the range to the map
		ranges[processing_step] = append(ranges[processing_step], MappedRange{destination, source, length})
	}

	var lowest_location int = -1

	// process the seeds until they reach the location
	for _, seed := range seeds {
		var step_number int = seed
		for _, step := range steps {
			// for every step check if the seed is in the range
			for _, range_ := range ranges[step] {
				if step_number >= range_.Source && step_number < range_.Source+range_.Length {
					step_number = range_.Destination + step_number - range_.Source

					// break out of the current for loop and continue with the next step
					break
				}
			}
		}

		// check if the current seed is the lowest
		if lowest_location == -1 || step_number < lowest_location {
			lowest_location = step_number
		}
	}

	return fmt.Sprint(lowest_location)
}
