package day10

import (
	"fmt"
	"os"
	"slices"
	"strings"
	"time"
)

func (d Day10) Part2(filename *string) string {
	var start = time.Now()

	var content, _ = os.ReadFile(*filename)
	var lines = strings.Split(string(content), "\n")

	var tiles [][]Tile = [][]Tile{}
	var startTile Tile

	// Parse the tiles into a 2D array
	for y, line := range lines {
		var parsedTiles []Tile = []Tile{}
		for x, char := range line {
			parsedTiles = append(parsedTiles, Tile{char, x, y})

			if char == 'S' {
				startTile = Tile{char, x, y}
			}
		}
		tiles = append(tiles, parsedTiles)
	}

	var startingDirs = []Direction{
		{startTile.x, startTile.y - 1},
		{startTile.x, startTile.y + 1},
		{startTile.x + 1, startTile.y},
		{startTile.x - 1, startTile.y},
	}

	// Keep track of the current steps
	var tilesInLoop []Tile = []Tile{startTile}
	var foundLoop = false

	// For each directions starting at the starting tile
	for _, dir := range startingDirs {
		if foundLoop {
			break
		}

		// Check if we are out of bounds
		if dir.x < 0 || dir.x >= len(tiles[0]) || dir.y < 0 || dir.y >= len(tiles) {
			continue
		}

		// Keep track of the current and previous tile
		// (in order to know what way we came from)
		var prevTile = startTile
		var currTile = tiles[dir.y][dir.x]

		// Keep looping until we find the startTile again (S)
		for currTile.char != 'S' {
			tilesInLoop = append(tilesInLoop, currTile)
			// If we find a | or - tile we need to keep going in the same direction
			if currTile.char == '|' || currTile.char == '-' {
				// If we came from the left we go right
				if prevTile.x == currTile.x-1 {
					// Make sure we do not go out of bounds, if so this direction is not valid
					if currTile.x+1 >= len(tiles[0]) {
						break
					}
					prevTile = currTile
					currTile = tiles[currTile.y][currTile.x+1]
					continue
				}

				// If we came from the right we go left
				if prevTile.x == currTile.x+1 {
					// Make sure we do not go out of bounds, if so this direction is not valid
					if currTile.x-1 < 0 {
						break
					}
					prevTile = currTile
					currTile = tiles[currTile.y][currTile.x-1]
					continue
				}

				// If we came from the top we go down
				if prevTile.y == currTile.y-1 {
					// Make sure we do not go out of bounds, if so this direction is not valid
					if currTile.y+1 >= len(tiles) {
						break
					}
					prevTile = currTile
					currTile = tiles[currTile.y+1][currTile.x]
					continue
				}

				// If we came from the bottom we go up
				if prevTile.y == currTile.y+1 {
					// Make sure we do not go out of bounds, if so this direction is not valid
					if currTile.y-1 < 0 {
						break
					}
					prevTile = currTile
					currTile = tiles[currTile.y-1][currTile.x]
					continue
				}
			}

			// If we find a L tile we need to go to the top or right
			if currTile.char == 'L' {
				// If we came from the right we go up
				if prevTile.x == currTile.x+1 {
					// Make sure we do not go out of bounds, if so this direction is not valid
					if currTile.y-1 < 0 {
						break
					}
					prevTile = currTile
					currTile = tiles[currTile.y-1][currTile.x]
					continue
				}

				// If we came from the top we go right
				if prevTile.y == currTile.y-1 {
					// Make sure we do not go out of bounds, if so this direction is not valid
					if currTile.x+1 >= len(tiles[0]) {
						break
					}
					prevTile = currTile
					currTile = tiles[currTile.y][currTile.x+1]
					continue
				}
			}

			// If we find a F tile we need to go to the bottom or right
			if currTile.char == 'F' {
				// If we came from the right we go down
				if prevTile.x == currTile.x+1 {
					// Make sure we do not go out of bounds, if so this direction is not valid
					if currTile.y+1 >= len(tiles) {
						break
					}
					prevTile = currTile
					currTile = tiles[currTile.y+1][currTile.x]
					continue
				}

				// If we came from the bottom we go right
				if prevTile.y == currTile.y+1 {
					// Make sure we do not go out of bounds, if so this direction is not valid
					if currTile.x+1 >= len(tiles[0]) {
						break
					}
					prevTile = currTile
					currTile = tiles[currTile.y][currTile.x+1]
					continue
				}
			}

			// If we find a 7 tile we need to go to the left or bottom
			if currTile.char == '7' {
				// If we came from the left we go down
				if prevTile.x == currTile.x-1 {
					// Make sure we do not go out of bounds, if so this direction is not valid
					if currTile.y+1 >= len(tiles) {
						break
					}
					prevTile = currTile
					currTile = tiles[currTile.y+1][currTile.x]
					continue
				}

				// If we came from the bottom we go left
				if prevTile.y == currTile.y+1 {
					// Make sure we do not go out of bounds, if so this direction is not valid
					if currTile.x-1 < 0 {
						break
					}
					prevTile = currTile
					currTile = tiles[currTile.y][currTile.x-1]
					continue
				}
			}

			// If we find a J tile we need to go to the left or top
			if currTile.char == 'J' {
				// If we came from the left we go up
				if prevTile.x == currTile.x-1 {
					// Make sure we do not go out of bounds, if so this direction is not valid
					if currTile.y-1 < 0 {
						break
					}
					prevTile = currTile
					currTile = tiles[currTile.y-1][currTile.x]
					continue
				}

				// If we came from the top we go left
				if prevTile.y == currTile.y-1 {
					// Make sure we do not go out of bounds, if so this direction is not valid
					if currTile.x-1 < 0 {
						break
					}
					prevTile = currTile
					currTile = tiles[currTile.y][currTile.x-1]
					continue
				}
			}

			// If we do not find any of the above tiles we have found a loop
			// and we can break out of the loop
			break
		}

		// If we found the startTile we have found a loop
		if currTile.char == 'S' {
			foundLoop = true
			break
		} else {
			tilesInLoop = []Tile{startTile}
		}
	}

	var enclosedTiles = 0
	var lastInLoopTile = Tile{}
	for y := 0; y < len(tiles); y++ {
		var walls int = 0
		for x := 0; x < len(tiles[0]); x++ {
			// Check if point is part of the loop
			if slices.ContainsFunc(tilesInLoop, func(p Tile) bool { return p.x == x && p.y == y }) {
				tile := tiles[y][x]

				if tile.char == '-' {
					continue
				}

				if tile.char == '|' {
					walls += 1
				}

				if tile.char == 'S' {
					walls += 1
				}

				if tile.char == 'J' && lastInLoopTile.char == 'F' {
					walls += 1
				}

				if tile.char == '7' && lastInLoopTile.char == 'L' {
					walls += 1
				}

				lastInLoopTile = tile

				continue
			}

			// If there are an odd number of walls, the tile is enclosed
			if walls%2 != 0 {
				enclosedTiles += 1
			}
		}
	}

	fmt.Println(time.Since(start))
	return fmt.Sprint(enclosedTiles)
}
