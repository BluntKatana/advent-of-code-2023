package day11

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// For every empty row or column, add n rows or columns to that side
func (u Universe) ExpandN(n int) {
	// find max X and Y
	var maxX, maxY int
	for _, asteroid := range u {
		if asteroid.X > maxX {
			maxX = asteroid.X
		}
		if asteroid.Y > maxY {
			maxY = asteroid.Y
		}
	}

	// Check if there are empty rows or columns
	var emptyRows, emptyColumns []int
	for y := 0; y <= maxY; y++ {
		var found = false
		for _, asteroid := range u {
			if asteroid.Y == y {
				found = true
				break
			}
		}
		if !found {
			emptyRows = append(emptyRows, y)
		}
	}
	for x := 0; x <= maxX; x++ {
		var found = false
		for _, asteroid := range u {
			if asteroid.X == x {
				found = true
				break
			}
		}
		if !found {
			emptyColumns = append(emptyColumns, x)
		}
	}

	// For each asteroid in the universe, update their X and Y
	for i, asteroid := range u {
		for _, emptyRow := range emptyRows {
			if asteroid.Y > emptyRow {
				u[i].Y += n - 1
			}
		}
		for _, emptyColumn := range emptyColumns {
			if asteroid.X > emptyColumn {
				u[i].X += n - 1
			}
		}
	}
}

func (d Day11) Part2(filename *string) string {
	var start = time.Now()

	var content, _ = os.ReadFile(*filename)
	var lines = strings.Split(string(content), "\n")

	universe := Universe{}

	var galaxyCount = 1
	for y, line := range lines {
		for x, char := range line {
			if string(char) == "#" {
				universe.Add(Asteroid{X: x, Y: y, Number: galaxyCount})
				galaxyCount++
			}
		}
	}

	// Expand the universe
	universe.ExpandN(1_000_000)

	// Find the distances between all asteroids (Manhattan distance)
	totalSum := 0
	for i, asteroid := range universe {
		for j := i + 1; j < len(universe); j++ {
			var nextAsteroid = universe[j]
			totalSum += int(universe.ManhattanDistance(asteroid, nextAsteroid))
		}
	}
	fmt.Println(time.Since(start))
	return fmt.Sprint(totalSum)
}
