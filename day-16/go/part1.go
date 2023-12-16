package day16

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type Day16 struct{}

type MirrorGrid [][]rune

type Position struct {
	x, y   int
	dx, dy int
}

type PositionMap map[Position]bool

func (m *MirrorGrid) Print(visited []Position) {
	for y, row := range *m {
		for x := range row {
			var found = false
			for _, pos := range visited {
				if pos.x == x && pos.y == y {
					found = true
					break
				}
			}
			if found {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func (m *MirrorGrid) GetEnergized(visited PositionMap) int {
	var energized = 0

	for y, row := range *m {
		for x := range row {
			var found = false
			for pos := range visited {
				if pos.x == x && pos.y == y {
					found = true
					break
				}
			}
			if found {
				energized++
			}
		}
	}

	return energized
}

func (m *MirrorGrid) Traverse(x, y, dx, dy int, visited *PositionMap) {
	var position = Position{x, y, dx, dy}

	// Check if we have already traversed this path
	if (*visited)[position] {
		return
	}

	// Check if we are out of bounds
	if y < 0 || y > len(*m)-1 || x < 0 || len((*m)[y])-1 < x {
		return
	}

	// Add current position to visited
	(*visited)[position] = true

	// Check if we hit a mirror
	switch (*m)[y][x] {
	case '.':
		// Continue in the same direction
		m.Traverse(x+dx, y+dy, dx, dy, visited)
		return
	case '|':
		if dx == 0 {
			m.Traverse(x+dx, y+dy, dx, dy, visited)
			return
		}
		m.Traverse(x, y+1, 0, 1, visited)
		m.Traverse(x, y-1, 0, -1, visited)
		return
	case '-':
		if dy == 0 {
			m.Traverse(x+dx, y+dy, dx, dy, visited)
			return
		}
		m.Traverse(x+1, y, 1, 0, visited)
		m.Traverse(x-1, y, -1, 0, visited)
		return
	case '/':
		// going up
		if dx == 0 && dy == -1 {
			m.Traverse(x+1, y, 1, 0, visited)
			return
		}
		// going down
		if dx == 0 && dy == 1 {
			m.Traverse(x-1, y, -1, 0, visited)
			return
		}
		// going right
		if dx == 1 && dy == 0 {
			m.Traverse(x, y-1, 0, -1, visited)
			return
		}
		// going left
		if dx == -1 && dy == 0 {
			m.Traverse(x, y+1, 0, 1, visited)
			return
		}
	case '\\':
		// going up
		if dx == 0 && dy == -1 {
			m.Traverse(x-1, y, -1, 0, visited)
			return
		}
		// going down
		if dx == 0 && dy == 1 {
			m.Traverse(x+1, y, 1, 0, visited)
			return
		}
		// going left
		if dx == 1 && dy == 0 {
			m.Traverse(x, y+1, 0, 1, visited)
			return
		}
		// going right
		if dx == -1 && dy == 0 {
			m.Traverse(x, y-1, 0, -1, visited)
			return
		}
	}
}

func (d Day16) Part1(filename *string) string {
	var start = time.Now()

	var content, _ = os.ReadFile(*filename)
	var lines = strings.Split(string(content), "\n")

	// Create the grid with mirrors
	var grid MirrorGrid = make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}
	grid = MirrorGrid(grid)

	var x, y = 0, 0
	var dx, dy = 1, 0
	var positionMap = make(PositionMap)

	grid.Traverse(x, y, dx, dy, &positionMap)
	var energized = grid.GetEnergized(positionMap)

	fmt.Println(time.Since(start))
	return fmt.Sprint(energized)
}
