package day16

import (
	"fmt"
	"os"
	"strings"
	"sync"
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

	// Initialize parallelism
	var wg sync.WaitGroup
	wg.Add(len(grid)*2 + len(grid[0])*2)

	// Find the max energized
	var maxEnergized = 0

	// Check the top and bottom sides
	for x := 0; x < len(grid); x++ {
		go FindEnergized(x, 0, 0, 1, grid, &maxEnergized, &wg)
		go FindEnergized(x, len(grid)-1, 0, -1, grid, &maxEnergized, &wg)
	}

	// Check the left and right sides
	for y := 0; y < len(grid[0]); y++ {
		go FindEnergized(0, y, 1, 0, grid, &maxEnergized, &wg)
		go FindEnergized(len(grid)-1, y, -1, 0, grid, &maxEnergized, &wg)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	fmt.Println(time.Since(start))
	return fmt.Sprint(maxEnergized)
}

func FindEnergized(x, y, dx, dy int, grid MirrorGrid, maxEnergized *int, wg *sync.WaitGroup) {
	var positionMap = make(PositionMap)
	grid.Traverse(x, y, dx, dy, &positionMap)
	var energized = grid.GetEnergized(positionMap)
	if energized > *maxEnergized {
		*maxEnergized = energized
	}
	wg.Done()
}
