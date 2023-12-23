package day19

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type PartRange map[string][2]int

func CloneMap(original map[string][2]int) map[string][2]int {
	c := make(map[string][2]int, 4)
	for k, v := range original {
		c[k] = v
	}
	return c
}

func count(workflows map[string]Workflow, workflow string, values map[string][2]int) int {
	total := 0
	for _, step := range workflows[workflow].Steps {
		// Last step
		if len(step.Intruction) == 0 {
			fmt.Println("Last step", step)
			// check for a resul
			switch step.Result {
			case "A":
				return 0
			case "R":
				product := 1
				for _, v := range values {
					product *= (v[1] - v[0] + 1)
				}
				return product
			}

			total += count(workflows, step.Result, values)
			return total
		}

		var category = string(step.Intruction[0])                // "x", "m", "a", or "s"
		var operator = string(step.Intruction[1])                // ">" or "<"
		var right, _ = strconv.Atoi(string(step.Intruction[2:])) // number

		// low high for this category
		v := values[category]

		// for every rule, the category this rule applies to splits
		// in a true and a false range
		// one below and one above the condition
		var tv [2]int
		var fv [2]int
		if operator == "<" {
			tv = [2]int{v[0], right - 1}
			fv = [2]int{right, v[1]}
		} else if operator == ">" {
			tv = [2]int{right + 1, v[1]}
			fv = [2]int{v[0], right}
		} else {
			total += count(workflows, step.Result, values)
			continue
		}

		// for true range, create a clone and keep counting
		// but start at consequence instead
		if tv[0] <= tv[1] {
			v2 := CloneMap(values)
			v2[category] = tv
			total += count(workflows, step.Result, v2)
		}

		// for false range, keep processing rest of rules
		// unless low is already higher than high
		// then we can just break
		if fv[0] > fv[1] {
			break
		}

		values[category] = fv
	}

	return total
}

func SolveRanges(workflows Workflows, workflow string, partRange PartRange) int {
	var currWorkflow Workflow = workflows[workflow]
	var total int = 0

	for _, step := range currWorkflow.Steps {
		// Last step of the workflow
		if step.Intruction == "" {
			switch step.Result {
			case "A":
				// return the product of all the ranges
				var total = 1
				for _, r := range partRange {
					total *= (r[1] - r[0] + 1)
				}
				return total
			case "R":
				return 0
			default:
				total += SolveRanges(workflows, step.Result, partRange)
				continue
			}
		}

		var pieceToCheck = string(step.Intruction[0])                        // "x", "m", "a", or "s"
		var typeOfCheck = string(step.Intruction[1])                         // ">" or "<"
		var numberToCompareTo, _ = strconv.Atoi(string(step.Intruction[2:])) // number

		var trueValues [2]int
		var falseValues [2]int

		switch typeOfCheck {
		case ">":
			trueValues = [2]int{numberToCompareTo + 1, partRange[pieceToCheck][1]}
			falseValues = [2]int{partRange[pieceToCheck][0], numberToCompareTo}
		case "<":
			trueValues = [2]int{partRange[pieceToCheck][0], numberToCompareTo - 1}
			falseValues = [2]int{numberToCompareTo, partRange[pieceToCheck][1]}
		}

		if trueValues[0] <= trueValues[1] {
			var trueRange = CloneMap(partRange)
			trueRange[pieceToCheck] = trueValues
			total += SolveRanges(workflows, step.Result, trueRange)
		}

		if falseValues[0] > falseValues[1] {
			break
		}

		partRange[pieceToCheck] = falseValues
	}

	return total
}

func (d Day19) Part2(filename *string) string {
	var start = time.Now()

	var content, _ = os.ReadFile(*filename)
	var lines = strings.Split(string(content), "\n")

	// Parse the workflows
	var workflows Workflows = make(map[string]Workflow)
	// var start_of_parts int
	for _, line := range lines {
		if line == "" {
			// start_of_parts = i + 1
			break
		}
		var split_line = strings.Split(line, "{")
		var name = split_line[0]

		var unparsedSteps = strings.Split(split_line[1], ",")
		var steps []Step

		for i, step := range unparsedSteps {
			if i == len(unparsedSteps)-1 {
				steps = append(steps, Step{Intruction: "", Result: step[0 : len(step)-1]})
				continue
			}

			var split_step = strings.Split(step, ":")
			var instruction = split_step[0]
			var result = split_step[1]

			steps = append(steps, Step{Intruction: instruction, Result: result})
		}
		workflows[name] = Workflow{Steps: steps, Name: name}
	}

	// Create all combinations of parts for each property 1-4000
	var partRange = PartRange{
		"x": [2]int{1, 4000},
		"m": [2]int{1, 4000},
		"a": [2]int{1, 4000},
		"s": [2]int{1, 4000},
	}

	fmt.Println("Solving")

	// check how many different parts are valid
	var total int = SolveRanges(workflows, "in", partRange)
	var total2 int = count(workflows, "in", partRange)

	fmt.Println(time.Since(start))
	return fmt.Sprint(total, total2)
}
