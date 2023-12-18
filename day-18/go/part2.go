package day18

import (
	"fmt"
	"math/big"
	"os"
	"strings"
	"time"
)

func (d Day18) Part2(filename *string) string {
	var start = time.Now()

	var content, _ = os.ReadFile(*filename)
	var lines = strings.Split(string(content), "\n")

	var positions []Position = []Position{{0, 0}}
	var total_length = 0

	for _, line := range lines {
		var split_line = strings.Split(line, " ")
		// var dir = split_line[0]
		// var length, _ = strconv.Atoi(split_line[1])

		// convert hex to int
		var unparsed_dir = string(split_line[2][len(split_line[2])-2])
		var unparsed_length = split_line[2][2 : len(split_line[2])-2]

		var big_dir = new(big.Int)
		big_dir.SetString(unparsed_dir, 16)
		var dir = int(big_dir.Int64())

		var big_length = new(big.Int)
		big_length.SetString(unparsed_length, 16)
		var length = int(big_length.Int64())

		var prev_pos = positions[len(positions)-1]

		switch dir {
		// up
		case 3:
			newPos := Position{prev_pos.X, prev_pos.Y - length}
			positions = append(positions, newPos)
		// down
		case 1:
			newPos := Position{prev_pos.X, prev_pos.Y + length}
			positions = append(positions, newPos)
		// left
		case 2:
			newPos := Position{prev_pos.X - length, prev_pos.Y}
			positions = append(positions, newPos)
		// right
		case 0:
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
