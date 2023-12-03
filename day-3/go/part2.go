package day3

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func filename_part2(test_mode bool) string {
	if test_mode {
		return "./day-3/test_part2.txt"
	}
	return "./day-3/input.txt"
}

var gears = map[Dir][]int{}

func (d Day3) Part2(test_mode bool) string {
	content, _ := os.ReadFile(filename_part2(test_mode))
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

	// iterate through engine to find valid words
	for row_num, row := range engine {
		// create a list of numbers which are later combines
		// into a single potetial part number
		var curr_num int = 0
		var gears_of_num []Dir = []Dir{}

		for col_num, char := range row {
			num, err := strconv.Atoi(char)

			if err == nil {
				// if number then add to current number
				curr_num = curr_num*10 + num
			} else {
				if curr_num > 0 {
					fmt.Println(curr_num, gears_of_num)
				}
				// if not number then add the current number to gears map
				if len(gears_of_num) > 0 && curr_num > 0 {
					for _, gear := range gears_of_num {
						gears[gear] = append(gears[gear], curr_num)
					}
				}

				// reset current number and gears_of_num
				curr_num = 0
				gears_of_num = []Dir{}
			}

			// if there is a gear in the surrounding
			// then add it to the gears_of_num
			has_gear, gears := has_gear_in_surrounding(row_num, col_num, engine)
			if has_gear && err == nil {
				// check if the gear is already in the gears_of_num
				// if not then add it
				for _, gear := range gears {
					var is_in_gears_of_num bool = false
					for _, gear_of_num := range gears_of_num {
						if gear_of_num == gear {
							is_in_gears_of_num = true
							break
						}
					}

					if !is_in_gears_of_num {
						gears_of_num = append(gears_of_num, gear)
					}
				}
			}
		}

		// check for last number
		if len(gears_of_num) > 0 && curr_num > 0 {
			for _, gear := range gears_of_num {
				gears[gear] = append(gears[gear], curr_num)
			}
		}
	}

	fmt.Println("MAP:", gears)

	// count the number of gears with 2 or more numbers
	var total_part_numbers int = 0
	for _, nums := range gears {
		if len(nums) == 2 {
			total_part_numbers += nums[0] * nums[1]
		}
	}

	return fmt.Sprint(total_part_numbers)
}

func has_gear_in_surrounding(row_num int, col_num int, array_2d [][]string) (bool, []Dir) {
	var surrounding []Dir = []Dir{
		{row_num - 1, col_num}, {row_num + 1, col_num}, // up, down
		{row_num, col_num - 1}, {row_num, col_num + 1}, // left, right
		{row_num - 1, col_num - 1}, {row_num - 1, col_num + 1}, // up-left, up-right
		{row_num + 1, col_num - 1}, {row_num + 1, col_num + 1}, // down-left, down-right
	}

	var gears []Dir = []Dir{}

	for _, dir := range surrounding {
		if dir.row >= 0 && dir.row < len(array_2d) && dir.col >= 0 && dir.col < len(array_2d[0]) {
			if array_2d[dir.row][dir.col] == "*" {
				gears = append(gears, dir)
			}
		}
	}

	return len(gears) > 0, gears
}
