package day8

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type Day8 struct{}

type Direction struct {
	Left  string
	Right string
}

func (d Day8) Part1(filename *string) string {
	// start clock
	start := time.Now()

	content, _ := os.ReadFile(*filename)
	lines := strings.Split(string(content), "\n")

	var network = map[string]Direction{}
	var instructions string

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
	}

	// Start at the first element of the network
	var steps = 0
	var curr_element = "AAA"
	var instruction_idx = 0
	for curr_element != "ZZZ" {
		var directions = network[curr_element]
		var curr_direction = instructions[instruction_idx]
		if string(curr_direction) == "L" {
			curr_element = directions.Left
		} else {
			curr_element = directions.Right
		}

		if instruction_idx == len(instructions)-1 {
			instruction_idx = 0
		} else {
			instruction_idx++
		}

		steps++
	}

	fmt.Println(time.Since(start))
	return fmt.Sprint(steps)
}
