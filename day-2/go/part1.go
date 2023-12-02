package day2

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// initialize day 2
type Day2 struct{}

func filename_part1(test_mode bool) string {
	if test_mode {
		return "./day-2/test_part1.txt"
	}
	return "./day-2/input.txt"
}

var MAP_MAX_PER_COLOR = map[string]int{
	"red": 12, "green": 13, "blue": 14,
}

func (d Day2) Part1(test_mode bool) string {
	content, _ := os.ReadFile(filename_part1(test_mode))
	lines := strings.Split(string(content), "\n")

	game_id_total := 0

	for _, line := range lines {
		// Grab the game_id as a number and grab the sets of cubes
		game_id_and_cubes := strings.Split(line, ": ")
		game_id, _ := strconv.Atoi(strings.Split(game_id_and_cubes[0], " ")[1])
		sets_of_cubes := strings.Split(game_id_and_cubes[1], "; ")

		is_valid_set := true

		for _, set := range sets_of_cubes {
			// Parse a set of cubes into single cubes
			// "1 blue, 2 green, 3 red" => ["1 blue", "2 green", "3 red"]
			cubes_in_set := strings.Split(set, ", ")
			for _, cube := range cubes_in_set {
				// Parse the cube "1 blue" to number and color
				splitted_cube := strings.Split(cube, " ")
				num, _ := strconv.Atoi(splitted_cube[0])
				color := splitted_cube[1]

				// Check if cube amount is larger than
				// the maximum allowed
				if num > MAP_MAX_PER_COLOR[color] {
					is_valid_set = false
				}
			}
		}

		if is_valid_set {
			game_id_total += game_id
		}
	}

	return fmt.Sprint(game_id_total)
}
