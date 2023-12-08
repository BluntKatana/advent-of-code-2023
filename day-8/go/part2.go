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

	var steps_to_z []int

	// For each starting element find the number of steps to a Z
	// and add it to the list
	for _, starting_element := range starting_elements {
		// Start at the first element of the network
		var steps = 0
		var curr_element = fmt.Sprint(starting_element)
		var instruction_idx = 0

		// Keep going until we reach a Z
		for curr_element[len(curr_element)-1] != 'Z' {
			var directions = network[curr_element]
			var curr_direction = instructions[instruction_idx]

			// Choose a direction
			if string(curr_direction) == "L" {
				curr_element = directions.Left
			} else {
				curr_element = directions.Right
			}

			// Make sure the instruction index doesn't go out of bounds
			if instruction_idx == len(instructions)-1 {
				instruction_idx = 0
			} else {
				instruction_idx++
			}

			// Increment the number of steps
			steps++
		}

		steps_to_z = append(steps_to_z, steps)
	}

	// To find the number of steps to Z where all the starting numbers
	// have the same number of steps to Z, we need to find the LCM of
	// all the steps to Z
	var steps = LCM(steps_to_z[0], steps_to_z[1], steps_to_z[2:]...)

	fmt.Println(time.Since(start))
	return fmt.Sprint(steps)
}

/**
 * The following two functions are taken from the GO.dev website
 * https://go.dev/play/p/SmzvkDjYlb
 */
// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
