package day3

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Day3 struct{}

type Coord struct {
	x int
	y int
}

func (d Day3) Part1(filename *string) string {
	start := time.Now()
	content, _ := os.ReadFile(*filename)
	lines := strings.Split(string(content), "\n")

	// initialize engine
	var total_y int = len(lines)
	var total_x int = len(lines[0])
	var engine [][]int = make([][]int, total_y)

	for i := 0; i < total_y; i++ {
		engine[i] = make([]int, total_x)
	}

	var special_chars []Coord

	// iterate through engine to find valid words
	for y, row := range lines {
		var num = 0
		var num_start = 0

		for x, char := range row {
			var parsed_num, _ = strconv.Atoi(string(char))
			var is_number bool = char >= '0' && char <= '9'

			if is_number {
				// add to current number
				num = num*10 + parsed_num
			} else {
				// if not number then add curren numbet
				// to engine at all positions starting from num_start
				if num > 0 {
					for i := num_start; i < x; i++ {
						engine[y][i] = num
					}
				}

				// reset current number and num_start
				num = 0
				num_start = x + 1
			}

			// if there is a special char in the surrounding
			if !is_number && char != '.' {
				special_chars = append(special_chars, Coord{x, y})
			}
		}

		// check for last number
		if num > 0 {
			for i := num_start; i < total_x; i++ {
				engine[y][i] = num
			}
		}
	}

	// iterate through special chars to find the total part numbers
	var total_part_numbers int = 1

	for _, coord := range special_chars {
		var directions = []Coord{
			{-1, 0}, {1, 0}, // down, up
			{0, -1}, {0, 1}, // right, left
			{-1, -1}, {1, 1}, // down-right, up-left
			{-1, 1}, {1, -1}, // down-left, up-right
		}

		// iterate through directions to find numbers that surround
		// the special char
		for _, direction := range directions {
			var new_y = coord.y + direction.y
			var new_x = coord.x + direction.x

			// make sure that y and x are within the bounds of the engine
			if new_y < 0 || new_y >= total_y || new_x < 0 || new_x >= total_x || engine[new_y][new_x] == 0 {
				continue
			}

			// if the number is found then add it to the total part numbers
			total_part_numbers += engine[new_y][new_x]

			// remove the number from the engine so that it won't be counted again
			// also remove the places where the number is found
			engine[new_y][new_x] = 0

			// remove the lef and right part of the new_x and new_y
			// so that it won't be counted again
			for x := new_x + 1; x < total_x && engine[new_y][x] > 0; x++ {
				engine[new_y][x] = 0
				break
			}

			for x := new_x - 1; x >= 0 && engine[new_y][x] > 0; x-- {
				engine[new_y][x] = 0
				break
			}

		}
	}

	t := time.Now()
	elapsed := t.Sub(start)

	fmt.Println("Part 1 took", elapsed)

	return fmt.Sprint(total_part_numbers)
}
