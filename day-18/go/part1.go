package day18

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

type Day18 struct{}

type Position struct {
	X, Y int
}

func shoelace(pts []Position) int {
	sum := 0
	p0 := pts[len(pts)-1]
	for _, p1 := range pts {
		sum += p0.Y*p1.X - p0.X*p1.Y
		p0 = p1
	}
	return int(math.Abs(float64(sum / 2)))
}

func (d Day18) Part1(filename *string) string {
	var start = time.Now()

	var content, _ = os.ReadFile(*filename)
	var lines = strings.Split(string(content), "\n")

	var positions []Position = []Position{{0, 0}}
	var total_length = 0

	for _, line := range lines {
		var split_line = strings.Split(line, " ")
		var dir = split_line[0]
		var length, _ = strconv.Atoi(split_line[1])
		// var hex = split_line[2][1 : len(split_line[2])-1]

		var prev_pos = positions[len(positions)-1]

		switch dir {
		case "U":
			newPos := Position{prev_pos.X, prev_pos.Y - length}
			positions = append(positions, newPos)
		case "D":
			newPos := Position{prev_pos.X, prev_pos.Y + length}
			positions = append(positions, newPos)
		case "L":
			newPos := Position{prev_pos.X - length, prev_pos.Y}
			positions = append(positions, newPos)
		case "R":
			newPos := Position{prev_pos.X + length, prev_pos.Y}
			positions = append(positions, newPos)
		}

		total_length += length
	}

	// add up
	// - area of the polygon
	// - perimeter of the polygon
	var area = shoelace(positions) + total_length/2 + 1

	fmt.Println(time.Since(start))
	return fmt.Sprint(area)
}
