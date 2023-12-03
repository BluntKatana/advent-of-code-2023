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

	// initialize 2d array
	total_rows := len(lines)
	total_cols := len(lines[0])
	engine := create_2d_array(total_rows, total_cols)

	// fill the engine
	for row_num, line := range lines {
		for col_num, char := range line {
			engine[row_num][col_num] = string(char)
		}
	}

	var gears = map[Coord][]int{}

	// iterate through engine to find gears
	for row_num, row := range engine {
		// keep track of current number and gears of current number
		var curr_num int = 0
		var gears_of_curr_num []Coord = []Coord{}

		for col_num, char := range row {
			num, err := strconv.Atoi(char)

			if err == nil {
				// if parsing to number is successful then add to current number
				curr_num = curr_num*10 + num
			} else {
				// if not number then add the current number to gears map
				if len(gears_of_curr_num) > 0 && curr_num > 0 {
					for _, gear := range gears_of_curr_num {
						gears[gear] = append(gears[gear], curr_num)
					}
				}

				// reset current number and gears_of_num
				curr_num = 0
				gears_of_curr_num = []Coord{}
			}

			// if there is a gear in the surrounding
			// then add it to the gears_of_num
			has_gear, gears := has_gear_in_surrounding(row_num, col_num, engine)
			if has_gear && err == nil {
				// check if the gear is already in the gears_of_num
				// if not then add it
				for _, gear := range gears {
					var is_in_gears_of_num bool = false
					for _, gear_of_num := range gears_of_curr_num {
						if gear_of_num == gear {
							is_in_gears_of_num = true
							break
						}
					}

					if !is_in_gears_of_num {
						gears_of_curr_num = append(gears_of_curr_num, gear)
					}
				}
			}
		}

		// check for last number
		if len(gears_of_curr_num) > 0 && curr_num > 0 {
			for _, gear := range gears_of_curr_num {
				gears[gear] = append(gears[gear], curr_num)
			}
		}
	}

	// count the number of gears with exactly 2 numbers
	var total_part_numbers int = 0
	for _, nums := range gears {
		if len(nums) == 2 {
			total_part_numbers += nums[0] * nums[1]
		}
	}

	return fmt.Sprint(total_part_numbers)
}

func has_gear_in_surrounding(row_num int, col_num int, array_2d [][]string) (bool, []Coord) {
	var surrounding []Coord = []Coord{
		{row_num - 1, col_num}, {row_num + 1, col_num}, // up, down
		{row_num, col_num - 1}, {row_num, col_num + 1}, // left, right
		{row_num - 1, col_num - 1}, {row_num - 1, col_num + 1}, // up-left, up-right
		{row_num + 1, col_num - 1}, {row_num + 1, col_num + 1}, // down-left, down-right
	}

	var gears []Coord = []Coord{}

	for _, dir := range surrounding {
		if dir.row >= 0 && dir.row < len(array_2d) && dir.col >= 0 && dir.col < len(array_2d[0]) {
			if array_2d[dir.row][dir.col] == "*" {
				gears = append(gears, dir)
			}
		}
	}

	return len(gears) > 0, gears
}
