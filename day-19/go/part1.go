package day19

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Day19 struct{}

type Step struct {
	Intruction string
	Result     string
}

type Workflow struct {
	Steps []Step
	Name  string
}

type Workflows map[string]Workflow

type Part struct {
	X, M, A, S int
}

func (p *Part) Solve(workflows Workflows) bool {
	p.Print()

	var currWorkflow Workflow = workflows["in"]

	for {
		fmt.Println("---- Workflow:", currWorkflow.Name)
		for _, step := range currWorkflow.Steps {
			fmt.Println("Step:", step)
			// Last step
			if step.Intruction == "" {
				fmt.Println("Last step", step)
				// check for a resul
				switch step.Result {
				case "A":
					return true
				case "R":
					return false
				default:
					currWorkflow = workflows[step.Result]
				}

				break
			}

			var pieceToCheck = string(step.Intruction[0])                        // "x", "m", "a", or "s"
			var typeOfCheck = string(step.Intruction[1])                         // ">" or "<"
			var numberToCompareTo, _ = strconv.Atoi(string(step.Intruction[2:])) // number

			var partToCheck int
			switch pieceToCheck {
			case "x":
				partToCheck = p.X
			case "m":
				partToCheck = p.M
			case "a":
				partToCheck = p.A
			case "s":
				partToCheck = p.S
			}

			var check bool
			switch typeOfCheck {
			case ">":
				check = partToCheck > numberToCompareTo
			case "<":
				check = partToCheck < numberToCompareTo
			}

			if check {
				// Check for a result in the step
				switch step.Result {
				case "A":
					return true
				case "R":
					return false
				}

				currWorkflow = workflows[step.Result]
				break
			}
		}
	}
}

func (d Day19) Part1(filename *string) string {
	var start = time.Now()

	var content, _ = os.ReadFile(*filename)
	var lines = strings.Split(string(content), "\n")

	// Parse the workflows
	var workflows Workflows = make(map[string]Workflow)
	var start_of_parts int
	for i, line := range lines {
		if line == "" {
			start_of_parts = i + 1
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

	// Parse the parts
	var parts []Part
	for _, line := range lines[start_of_parts:] {
		var split_line = strings.Split(line[1:len(line)-1], ",")
		var x, _ = strconv.Atoi(split_line[0][2:])
		var m, _ = strconv.Atoi(split_line[1][2:])
		var a, _ = strconv.Atoi(split_line[2][2:])
		var s, _ = strconv.Atoi(split_line[3][2:])
		parts = append(parts, Part{X: x, M: m, A: a, S: s})
	}

	// Solve the parts
	var total int = 0
	for _, part := range parts {
		if part.Solve(workflows) {
			total += part.X + part.M + part.A + part.S
		}
	}

	fmt.Println(time.Since(start))
	return fmt.Sprint(total)
}

func (w Workflow) Print() {
	for _, step := range w.Steps {
		fmt.Println("Instruction:", step.Intruction, "Result:", step.Result)
	}
}

func (p Part) Print() {
	fmt.Println("Part: x=", p.X, "m=", p.M, "a=", p.A, "s=", p.S)
}
