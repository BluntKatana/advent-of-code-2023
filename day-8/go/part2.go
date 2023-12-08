package day8

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func (d Day8) Part2(filename *string) string {
	// start clock
	start := time.Now()

	content, _ := os.ReadFile(*filename)
	lines := strings.Split(string(content), "\n")

	var network = map[string]Direction{}
	var instructions string
	var starting_elements []string

	fmt.Println("TESTTT")

	// create network in form of map
	for idx, line := range lines {
		if idx == 1 {
			continue
		}

		if idx == 0 {
			instructions = line
			continue
		}

		var parts = strings.Split(line, " = ")
		var direction = strings.Split(parts[1], ", ")
		var left_direction = direction[0][1:]
		var right_direction = direction[1][:3]
		network[parts[0]] = Direction{left_direction, right_direction}

		if parts[0][2] == 'A' {
			starting_elements = append(starting_elements, parts[0])
		}
	}

	fmt.Println(starting_elements)

	// Start at the first element of the network
	var steps = 0
	var curr_elements = starting_elements
	var instruction_idx = 0

	for !all_end_with_z(curr_elements) {
		var new_elements []string
		var curr_direction = instructions[instruction_idx]

		// loop through all elements and update them
		for _, element := range curr_elements {
			var directions = network[element]
			if string(curr_direction) == "L" {
				new_elements = append(new_elements, directions.Left)
			} else {
				new_elements = append(new_elements, directions.Right)
			}
		}

		// update curr_elements
		curr_elements = new_elements

		// update curr_elements and steps
		if instruction_idx == len(instructions)-1 {
			instruction_idx = 0
		} else {
			instruction_idx++
		}
		steps++

		fmt.Println(steps, curr_elements)
	}

	fmt.Println(time.Since(start))
	return fmt.Sprint(steps)
}

// check if a list of strings all end with a Z
func all_end_with_z(list []string) bool {
	for _, element := range list {
		if element[len(element)-1] != 'Z' {
			return false
		}
	}
	return true
}
