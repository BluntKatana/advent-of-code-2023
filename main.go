package main

// Import packages and all the day modules
import (
	"day1"
	"day10"
	"day11"
	"day12"
	"day13"
	"day14"
	"day15"
	"day16"
	"day2"
	"day3"
	"day4"
	"day5"
	"day6"
	"day7"
	"day8"
	"day9"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

type Day interface {
	Part1(*string) string
	Part2(*string) string
}

// Create a map of days
var days = map[int]Day{
	1: day1.Day1{}, 2: day2.Day2{}, 3: day3.Day3{},
	4: day4.Day4{}, 5: day5.Day5{}, 6: day6.Day6{},
	7: day7.Day7{}, 8: day8.Day8{}, 9: day9.Day9{},
	10: day10.Day10{}, 11: day11.Day11{}, 12: day12.Day12{},
	13: day13.Day13{}, 14: day14.Day14{}, 15: day15.Day15{},
	16: day16.Day16{},
}

// Main function
func main() {
	// Retrieve flags from command line (day, test)
	var day_flag = flag.Int("day", -1, "Day to run (between 1 and 25)")
	var part_flag = flag.Int("part", -1, "Part to run (1 or 2)")
	var test_flag = flag.Bool("test", false, "Run day in test mode (input.txt -> test.txt)")
	var all_flag = flag.Bool("all", false, "Run all days")
	flag.Parse()

	// Check if day is valid (1-31) or there is an all flag
	if (*day_flag < 0 || *day_flag > 25) && !*all_flag {
		fmt.Println("Choose a day between 1 and 31, or use -all to run all days")
		return
	}

	// Check if part is valid (1-2) or there is an all flag
	if (*part_flag < -1 || *part_flag > 2) && !*all_flag {
		fmt.Println("Choose either part 1 or part 2, or use -1 to run both parts")
		return
	}

	// Print status message
	if *all_flag {
		fmt.Println("-- Running all days in test mode:", *test_flag)
	} else {
		if *part_flag == -1 {
			fmt.Println("-- Running day", *day_flag, "in test mode:", *test_flag)
		} else {
			fmt.Println("-- Running day", *day_flag, "part", *part_flag, "in test mode:", *test_flag)
		}
	}
	fmt.Println()

	// Run all days up untill current one
	if *all_flag {
		var current_day = time.Now().Day()
		var current_year = time.Now().Year()

		// If the current year is past 2023, run all days of 2023
		if current_year > 2023 {
			current_day = 31
		}

		for day := 1; day <= current_day; day++ {
			d, ok := days[day]
			if !ok {
				fmt.Println("Day", day, "not implemented")
				continue
			}
			filename := create_filename(&day, part_flag, test_flag)
			log_to_file(fmt.Sprintf("Day %d Part 1: %s\tPart 2: %s\n", day, d.Part1(&filename), d.Part2(&filename)), "all")
		}
		return
	}

	// Run a specific day
	// sort the days first
	day, ok := days[*day_flag]
	if !ok {
		fmt.Println("Day", *day_flag, "not implemented")
		return
	}

	filename := create_filename(day_flag, part_flag, test_flag)

	// Run both parts
	if *part_flag == -1 {
		log_to_file(fmt.Sprintf("Day %d Part 1: %s\tPart 2: %s\n", *day_flag, day.Part1(&filename), day.Part2(&filename)), "day"+fmt.Sprintf("%d", *day_flag))
		return
	}

	// Run a specific part
	if *part_flag == 1 {
		log_to_file(fmt.Sprintf("Day %d Part 1: %s\n", *day_flag, day.Part1(&filename)), "day"+fmt.Sprintf("%d", *day_flag))
		return
	}

	if *part_flag == 2 {
		log_to_file(fmt.Sprintf("Day %d Part 2: %s\n", *day_flag, day.Part2(&filename)), "day"+fmt.Sprintf("%d", *day_flag))
		return
	}
}

// Log the output to a file with a timestamp and print it to the console
// It keeps track of the amount of times a day is run
func log_to_file(str string, suffix string) {
	// Log to console
	fmt.Print(str)

	// Create log file or open existing one
	file_name := "go.log"

	// Add current date and suffix to log file name
	total_str := "#" + suffix + " | " + time.Now().Format("2006-01-02 15:04:05") + " | " + str

	// Add string attempt to log file with timestamp
	f, err := os.OpenFile(file_name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}

	// Write string to file
	if _, err := f.WriteString(total_str); err != nil {
		log.Println(err)
	}

	// Close file
	if err := f.Close(); err != nil {
		log.Println(err)
	}
}

// Create filename based on day, part and test mode
func create_filename(day *int, part *int, test_mode *bool) string {
	// Create filename (01 -> 25)

	var filename string

	// Make sure day is always 2 digits
	if *day < 10 {
		filename = "./day-" + fmt.Sprintf("0%d", *day) + "/input.txt"
	} else {
		filename = "./day-" + fmt.Sprintf("%d", *day) + "/input.txt"
	}

	if !*test_mode {
		return filename
	}
	// If test mode is on, change filename
	if *part == 1 {
		filename = "./day-" + fmt.Sprintf("%d", *day) + "/test_part1.txt"
	} else {
		filename = "./day-" + fmt.Sprintf("%d", *day) + "/test_part2.txt"
	}

	return filename
}
