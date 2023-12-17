package day17

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type Day17 struct{}

type Blocks [][]int
type Direction struct {
	X, Y int
}

var DIRECTIONS = []Direction{
	{0, -1}, {0, 1}, {-1, 0}, {1, 0},
}

type Node struct {
	X, Y, Steps int
	Dir         Direction
}

func (b Blocks) ShortestPath() int {
	// initialize distance matrix with max int
	var dist = make([][]int, len(b))
	for i := range dist {
		dist[i] = make([]int, len(b[i]))
		for j := range dist[i] {
			dist[i][j] = -1
		}
	}

	// initialize queue
	var queue = make([]Node, 0)
	queue = append(queue, Node{0, 0, 0, Direction{0, 0}})
	dist[0][0] = 0

	for len(queue) > 0 {
		var current = queue[0]
		queue = queue[1:]

		for _, direction := range DIRECTIONS {
			var newX = current.X + direction.X
			var newY = current.Y + direction.Y

			// Check if the new position is out of bounds
			if newX < 0 || newY < 0 || newX >= len(b) || newY >= len(b[0]) {
				continue
			}

			// Check if the new position is already visited
			if dist[newX][newY] != -1 {
				continue
			}

			// Check if we can move in the same direction
			if current.Dir == direction && current.Steps < 3 {
				queue = append(queue, Node{newX, newY, current.Steps + 1, direction})
				dist[newX][newY] = dist[current.X][current.Y] + b[newX][newY]
			} else if current.Dir != direction {
				queue = append(queue, Node{newX, newY, 0, direction})
				dist[newX][newY] = dist[current.X][current.Y] + b[newX][newY]
			}

			// Update the distance
		}
	}

	// print the distance matrix
	for _, row := range dist {
		for _, col := range row {
			fmt.Print(col, "\t")
		}
		fmt.Println()
	}

	return dist[len(b)-1][len(b[0])-1]
}

func (d Day17) Part1(filename *string) string {
	var start = time.Now()

	var content, _ = os.ReadFile(*filename)
	var lines = strings.Split(string(content), "\n")

	// initialize the blocks
	var blocks = make(Blocks, len(lines))
	for i, line := range lines {
		blocks[i] = make([]int, len(line))
		for j, char := range line {
			blocks[i][j] = int(char - '0')
		}
	}

	// traverse the blocks
	var stepsTaken = blocks.ShortestPath()
	fmt.Println(stepsTaken)

	// blocks.Print()

	fmt.Println(time.Since(start))
	return fmt.Sprint(0)
}

func (b Blocks) Print() {
	for _, row := range b {
		for _, col := range row {
			fmt.Print(col)
		}
		fmt.Println()
	}
}
