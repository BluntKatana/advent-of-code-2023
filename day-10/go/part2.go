package day10

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func (t Tile) Str() string {
	return fmt.Sprintf("(Char: %c | X: %d | Y: %d)", t.char, t.x, t.y)
}

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

		// If we find a ground tile this direction is not valid
		if currTile.char == '.' {
			continue
		}

		// Keep looping until we find the startTile again (S)
		for currTile.char != 'S' {
			// If we find a ground tile this direction is not valid
			if currTile.char == '.' {
				break
			}

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

	// Find the amount of tiles which are enclosed by the starting tiles
	PrintTiles(tiles, startTile)

	// // Find the two tiles which are adjacent to the starting tile
	// var tilesConnectedToStartTile []Tile = []Tile{}
	// for _, tile := range tilesInLoop {
	// 	if tile.x == startTile.x && tile.y == startTile.y {
	// 		continue
	// 	}

	// 	for _, dir := range startingDirs {
	// 		if tile.x == dir.x && tile.y == dir.y {
	// 			tilesConnectedToStartTile = append(tilesConnectedToStartTile, tile)
	// 		}
	// 	}
	// }

	// For every column and row we loop through,
	// if we find a tile which is in the loop we update a counter
	// for every tile which not in the loop and where the counter is odd we know it is a tile which is enclosed by the loop
	var tilesEnclosedByLoop int = 0
	var tilesInLoopMatrix [][]Tile = tiles

	for _, row := range tiles {
		fmt.Println("----")
		PrintRow(row, "Full row")
		// Grab the tiles which are in the loop in this row
		var tilesInLoopInRow []Tile = []Tile{}
		var tilesInLoopInRowNoMinus []Tile = []Tile{}
		for _, tile := range row {
			for _, tileInLoop := range tilesInLoop {
				if tile.x == tileInLoop.x && tile.y == tileInLoop.y {
					tilesInLoopInRow = append(tilesInLoopInRow, tile)
					if tile.char != '-' {
						tilesInLoopInRowNoMinus = append(tilesInLoopInRowNoMinus, tile)
					}
				}
			}
		}

		PrintRow(tilesInLoopInRow, "Tiles in loop in row", fmt.Sprint(len(tilesInLoopInRow)))

		// We basically check every odd tile in the loop like so:
		// 1: ..|..|... = between the first and second tile there are 2 tiles which are enclosed by the loop
		// 2: ..|..|..|..|.. = between the first and second tile , and third and fourth 2 tiles which are enclosed by the loop
		// 3: ..F-----7 = between the F and 7 there are no tiles which are enclosed by the loop as they are part of the loop
		// So we have to iterate from the odd tile to the even tile and count the tiles between them,
		// unless in the case of 3 where the odd tile begins at the end of the loop
		var enclosedTilesInRow int = 0

		// Loop through every tile in the loop which is in this row
		for i := 0; i < len(tilesInLoopInRowNoMinus)-1; i++ {
			var tileFrom Tile = tilesInLoopInRowNoMinus[i]
			var tileFromIdx = i

			// Check if there is a tile next to the tile to check
			if tileFromIdx+1 >= len(tilesInLoopInRowNoMinus) {
				break
			}

			var tileTo Tile = tilesInLoopInRowNoMinus[i+1]
			var tileToIdx = i + 1

			// Add the amount of tiles between the tile to check and the next tile
			var enclosedTileAmount = tileTo.x - tileFrom.x - 1

			enclosedTilesInRow += enclosedTileAmount

			for i := 1; i < enclosedTileAmount+1; i++ {
				var tileBetween = tiles[tileFrom.y][tileFrom.x+i]
				MarkAsInsideLoop(tilesInLoopMatrix, tileBetween)
			}

			fmt.Println("Tile to check from:", tileFrom.Str(), "to", tileTo.Str(), "tiles between:", enclosedTileAmount, "index:", tileFromIdx)

			// Set the next index (i) to the index of the next tile
			i = tileToIdx
		}

		// If the amount of tiles enclosed in this row is odd we add it to the total
		fmt.Println("Enclosed tiles in row", enclosedTilesInRow)
		tilesEnclosedByLoop += enclosedTilesInRow
	}

	PrintTiles(tilesInLoopMatrix, startTile)

	fmt.Println(time.Since(start))
	return fmt.Sprint(tilesEnclosedByLoop)
}

func MarkAsInsideLoop(tiles [][]Tile, tile Tile) {
	tiles[tile.y][tile.x].char = 'X'
}

func PrintRow(row []Tile, prefix ...string) {
	for _, pre := range prefix {
		fmt.Print(pre, " ")
	}

	fmt.Print(": ")

	for _, tile := range row {
		fmt.Print(string(tile.char))
	}
	fmt.Println()
}
