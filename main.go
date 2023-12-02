// Description: Main file for Advent of Code 2020
package main

// Import packages
import (
	"day1"
	"day2"
	"day3"
	"flag"
	"fmt"
)

type Day interface {
	Part1(bool) string
	Part2(bool) string
}

// Create a map of days
var days = map[int]Day{
	1: day1.Day1{}, 2: day2.Day2{}, 3: day3.Day3{},
}

// Main function
func main() {
	// Retrieve flags from command line (day, test)
	var day_flag = flag.Int("day", -1, "Day to run (between 1 and 31)")
	var part_flag = flag.Int("part", -1, "Part to run (1 or 2)")
	var test_flag = flag.Bool("test", false, "Run day in test mode (input.txt -> test.txt)")
	var all_flag = flag.Bool("all", false, "Run all days")
	flag.Parse()

	// Check if day is valid (1-31) or there is an all flag
	if (*day_flag < -1 || *day_flag > 31) && !*all_flag {
		fmt.Println("Choose a day between 1 and 31, or use -all to run all days")
		return
	}

	// Check if part is valid (1-2) or there is an all flag
	if (*part_flag < -1 || *part_flag > 2) && !*all_flag {
		fmt.Println("Choose either part 1 or part 2, or use -1 to run both parts")
		return
	}

	// Print status message
	fmt.Println("Running day", *day_flag, "part", *part_flag, "in test mode:", *test_flag)
	fmt.Println()

	// Run all days
	if *all_flag {
		for day := 1; day <= 31; day++ {
			d, ok := days[day]
			if !ok {
				continue
			}
			fmt.Println("Day", day, "\tPart 1:", d.Part1(*test_flag), "\tPart 2:", d.Part2(*test_flag))
		}
		return
	}

	// Run a specific day
	// sort the days first
	day, ok := days[*day_flag]
	if !ok {
		fmt.Println("Day", *day_flag, "not implemented yet")
		return
	}

	// Run both parts
	if *part_flag == -1 {
		fmt.Println("Day", *day_flag, "Part 1:", day.Part1(*test_flag), "\tPart 2:", day.Part2(*test_flag))
		return
	}

	// Run a specific part
	if *part_flag == 1 {
		fmt.Println("Day", *day_flag, "Part 1")
		fmt.Println("Result:", day.Part1(*test_flag))
		return
	}

	if *part_flag == 2 {
		fmt.Println("Day", *day_flag, "Part 2")
		fmt.Println("Result:", day.Part2(*test_flag))
		return
	}
}
