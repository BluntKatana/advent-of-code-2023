package day3

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func (d Day3) Part2(filename *string) string {

	content, _ := os.ReadFile(*filename)
	lines := strings.Split(string(content), "\n")
	// initialize engine
	var total_y int = len(lines)
	var total_x int = len(lines[0])
	var engine [][]int = make([][]int, total_y)

	for i := 0; i < total_y; i++ {
		engine[i] = make([]int, total_x)
	}

	var gears []Coord

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
			if char == '*' {
				gears = append(gears, Coord{x, y})
			}
		}

		// check for last number
		if num > 0 {
			for i := num_start; i < total_x; i++ {
				engine[y][i] = num
			}
		}
	}

	var total_part_numbers int = 0

	// print gear
	for _, gear := range gears {
		fmt.Println(gear)
	}

	for _, gear := range gears {
		var directions = []Coord{
			{-1, 0}, {1, 0}, // down, up
			{0, -1}, {0, 1}, // right, left
			{-1, -1}, {1, 1}, // down-right, up-left
			{-1, 1}, {1, -1}, // down-left, up-right
		}

		var nums []int = []int{}

		for _, dir := range directions {
			var y_new int = gear.y + dir.y
			var x_new int = gear.x + dir.x

			// check if out of bounds and if number is 0
			if y_new < 0 || y_new >= total_y || x_new < 0 || x_new >= total_x || engine[y_new][x_new] == 0 {
				continue
			}

			// make sure that the number is not already in nums
			var already_in_nums bool = false

			for _, num := range nums {
				if num == engine[y_new][x_new] {
					already_in_nums = true
					break
				}
			}

			if already_in_nums {
				continue
			}

			nums = append(nums, engine[y_new][x_new])
		}

		// only add to total_part_numbers if there are 2 numbers next to the gear
		if len(nums) == 2 {
			total_part_numbers += nums[0] * nums[1]
		}
	}

	return fmt.Sprint(total_part_numbers)
}
