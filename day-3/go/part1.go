package day3

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Day3 struct{}

func filename_part1(test_mode bool) string {
	if test_mode {
		return "./day-3/test_part1.txt"
	}
	return "./day-3/input.txt"
}

func create_2d_array(row int, col int) [][]string {
	engine := make([][]string, row)
	for i := range engine {
		engine[i] = make([]string, col)
	}
	return engine
}

func (d Day3) Part1(test_mode bool) string {
	content, _ := os.ReadFile(filename_part1(test_mode))
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

	var total_part_numbers int = 0

	// iterate through engine to find valid words
	for row_num, row := range engine {
		// create a list of numbers which are later combines
		// into a single potetial part number
		var curr_num int = 0
		var is_valid_part bool = false

		for col_num, char := range row {
			num, err := strconv.Atoi(char)

			if err == nil {
				// if number then add to current number
				curr_num = curr_num*10 + num
			} else {
				// if not number then reset current number
				if is_valid_part && curr_num > 0 {
					total_part_numbers += curr_num
				}

				curr_num = 0
				is_valid_part = false
			}

			// check for special character surrounding number
			if err == nil && has_special_char_in_surrounding(row_num, col_num, engine) {
				is_valid_part = true
			}
		}

		// check for last number
		if is_valid_part && curr_num > 0 {
			total_part_numbers += curr_num
		}
	}

	return fmt.Sprint(total_part_numbers)
}

type Dir struct {
	row int
	col int
}

func has_special_char_in_surrounding(row_num int, col_num int, array_2d [][]string) bool {
	var surrounding []Dir = []Dir{
		{row_num - 1, col_num}, {row_num + 1, col_num}, // up, down
		{row_num, col_num - 1}, {row_num, col_num + 1}, // left, right
		{row_num - 1, col_num - 1}, {row_num - 1, col_num + 1}, // up-left, up-right
		{row_num + 1, col_num - 1}, {row_num + 1, col_num + 1}, // down-left, down-right
	}

	var has_special_char bool = false

	for _, dir := range surrounding {
		if dir.row >= 0 && dir.row < len(array_2d) && dir.col >= 0 && dir.col < len(array_2d[0]) {
			char := array_2d[dir.row][dir.col]
			if _, err := strconv.Atoi(char); err != nil && char != "." {
				has_special_char = true
			}
		}
	}

	return has_special_char
}
