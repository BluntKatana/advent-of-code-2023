package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	content, _ := os.ReadFile("../input.txt")
	lines := strings.Split(string(content), "\n")

	total := 0

	for _, line := range lines {
		// Grab the sets of cubes
		game_id_and_cubes := strings.Split(line, ": ")
		sets_of_cubes := strings.Split(game_id_and_cubes[1], "; ")

		largest_blue, largest_red, largest_green := 1, 1, 1

		for _, set := range sets_of_cubes {
			// Parse a set of cubes into single cubes
			// "1 blue, 2 green, 3 red" => ["1 blue", "2 green", "3 red"]
			cubes_in_set := strings.Split(set, ", ")
			for _, cube := range cubes_in_set {
				// Parse the cube "1 blue" to number and color
				splitted_cube := strings.Split(cube, " ")
				num, _ := strconv.Atoi(splitted_cube[0])
				color := splitted_cube[1]

				// Check for each color if the current num is largest
				switch color {
				case "blue":
					if num > largest_blue {
						largest_blue = num
					}
				case "red":
					if num > largest_red {
						largest_red = num
					}
				case "green":
					if num > largest_green {
						largest_green = num
					}
				}
			}
		}

		total += largest_blue * largest_green * largest_red
	}

	fmt.Println("Result: ", total)
}
