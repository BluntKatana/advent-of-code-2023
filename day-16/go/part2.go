package day16

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func (d Day16) Part2(filename *string) string {
	var start = time.Now()

	var content, _ = os.ReadFile(*filename)
	var lines = strings.Split(string(content), "\n")

	// Create the grid with mirrors
	var grid = make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}
	grid = MirrorGrid(grid)

	// Find the max energized
	var maxEnergized = 0

	// Check from top going down
	for x := 0; x < len(grid); x++ {
		var energized = FindEnergized(x, 0, 0, 1, grid)
		if energized > maxEnergized {
			maxEnergized = energized
		}
	}

	// Check from bottom going up
	for x := 0; x < len(grid); x++ {
		var energized = FindEnergized(x, len(grid)-1, 0, -1, grid)
		if energized > maxEnergized {
			maxEnergized = energized
		}
	}

	// Check from left going right
	for y := 0; y < len(grid[0]); y++ {
		var energized = FindEnergized(0, y, 1, 0, grid)
		if energized > maxEnergized {
			maxEnergized = energized
		}
	}

	// Check from right going left
	for y := 0; y < len(grid[0]); y++ {
		var energized = FindEnergized(len(grid)-1, y, -1, 0, grid)
		if energized > maxEnergized {
			maxEnergized = energized
		}
	}

	fmt.Println(time.Since(start))
	return fmt.Sprint(maxEnergized)
}

func FindEnergized(x, y, dx, dy int, lines [][]rune) int {
	var positionMap = make(PositionMap)
	var grid MirrorGrid = make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}
	grid = MirrorGrid(grid)
	grid.Traverse(x, y, dx, dy, &positionMap)
	var energized = grid.GetEnergized(positionMap)
	return energized
}
