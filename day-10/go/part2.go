package day10

import (
	"fmt"
	"os"
	"slices"
	"strings"
	"time"
)

func (t Tile) Str() string {
	return fmt.Sprintf("(Char: %c | X: %d | Y: %d)", t.char, t.x, t.y)
}

type Point struct {
	x, y int
}

type TileGrid [][]rune

func (t TileGrid) FindStart() Point {
	for y := 0; y < len(t); y++ {
		for x := 0; x < len(t[0]); x++ {
			if t[y][x] == 'S' {
				return Point{x, y}
			}
		}
	}

	return Point{-1, -1}
}

func (t TileGrid) Print(tiles ...Point) {
	for y, line := range t {
		for x, tile := range line {
			var found = false
			for _, p := range tiles {
				if p.x == x && p.y == y {
					found = true
					break
				}
			}

			if found {
				fmt.Printf("\033[1;31m%c\033[0m", tile)
				continue
			} else {
				fmt.Printf("%c", tile)
			}
		}
		fmt.Println()
	}
}

const (
	UP    = 0
	RIGHT = 1
	DOWN  = 2
	LEFT  = 3
)

func (t TileGrid) NextDir(p Point, dir int) int {
	var char = t[p.y][p.x]

	// Check which direction we're going
	switch char {
	case '|':
		return dir
	case '-':
		return dir
	case 'S':
		return dir
	// F: Coming from the right => go down, coming from down => go right
	case 'F':
		if dir == LEFT {
			return DOWN
		}
		if dir == UP {
			return RIGHT
		}
	// 7: Coming from the left => go down, coming from down => go left
	case '7':
		if dir == RIGHT {
			return DOWN
		}
		if dir == UP {
			return LEFT
		}
	// J: Coming from the left => go up, coming from up => go left
	case 'J':
		if dir == RIGHT {
			return UP
		}
		if dir == DOWN {
			return LEFT
		}
	case 'L':
		// L: Coming from the top => go right, coming from right => go up
		if dir == DOWN {
			return RIGHT
		}
		if dir == LEFT {
			return UP
		}
	}
	// fmt.Println("UNKNOWN CHAR", string(char))
	return -1
}

func (d Day10) Part2(filename *string) string {
	var start = time.Now()

	var content, _ = os.ReadFile(*filename)
	var lines = strings.Split(string(content), "\n")

	var tileGrid = TileGrid{}

	// Parse the tiles into a 2D array
	for y, line := range lines {
		for _, char := range line {
			if len(tileGrid) <= y {
				tileGrid = append(tileGrid, []rune{})
			}

			tileGrid[y] = append(tileGrid[y], char)
		}
	}

	var startPoint = tileGrid.FindStart()
	var loopTiles []Point = []Point{}

	var currPoint = startPoint
	var currDir int = LEFT

	for {
		// fmt.Println("BEFORE", currPoint, currDir)
		loopTiles = append(loopTiles, currPoint)

		currDir = tileGrid.NextDir(currPoint, currDir)

		switch currDir {
		case UP:
			currPoint.y -= 1
		case DOWN:
			currPoint.y += 1
		case RIGHT:
			currPoint.x += 1
		case LEFT:
			currPoint.x -= 1
		}

		// fmt.Println("AFTER", currPoint, currDir)

		if tileGrid[currPoint.y][currPoint.x] == 'S' {
			break
		}

	}

	var enclosedPoints []Point = []Point{}
	var lastSeenRune rune = ' '
	for y := 0; y < len(tileGrid); y++ {
		var walls int = 0
		for x := 0; x < len(tileGrid[0]); x++ {
			// Check if point is part of the loop
			if slices.ContainsFunc(loopTiles, func(p Point) bool { return p.x == x && p.y == y }) {
				char := tileGrid[y][x]

				if char == '-' {
					continue
				}

				if char == '|' {
					walls += 1
				}

				if char == 'S' {
					walls += 1
				}

				if char == 'J' && lastSeenRune == 'F' {
					walls += 1
				}

				if char == '7' && lastSeenRune == 'L' {
					walls += 1
				}

				lastSeenRune = char

				continue
			}

			// If there are an odd number of walls, the tile is enclosed
			if walls%2 != 0 {
				enclosedPoints = append(enclosedPoints, Point{x, y})
			}
		}
	}

	fmt.Println(time.Since(start))
	return fmt.Sprint(len(enclosedPoints))
}
