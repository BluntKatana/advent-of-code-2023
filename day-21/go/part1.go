package day21

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type Day21 struct{}

type Plots map[[2]int]bool
type Rocks map[[2]int]bool

func (d Day21) Part1(filename *string) string {
	var start = time.Now()

	var content, _ = os.ReadFile(*filename)
	var lines = strings.Split(string(content), "\n")

	var plots Plots = make(map[[2]int]bool)
	var rocks Rocks = make(map[[2]int]bool)
	for y, line := range lines {
		for x, char := range line {
			if char == '#' {
				rocks[[2]int{x, y}] = true
			} else if char == 'S' {
				plots[[2]int{x, y}] = true
			}
		}
	}

	// PrintGarden(rocks, plots)

	steps := 64

	for i := 0; i < steps; i++ {
		var newPlots Plots = make(map[[2]int]bool)
		for plot := range plots {
			for _, dir := range [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
				// create new position
				var newPos = [2]int{plot[0] + dir[0], plot[1] + dir[1]}

				// check if new position does not exist yet
				if _, ok := plots[newPos]; ok {
					continue
				}

				// check if new position is allowed (not on rock)
				if _, isOnRock := rocks[newPos]; isOnRock {
					continue
				}

				// if not on rock, add to new plots
				newPlots[newPos] = true
			}
		}

		plots = newPlots
		if i%1000 == 0 {
			fmt.Println(i, len(plots))
		}
	}

	fmt.Println(time.Since(start))
	return fmt.Sprint(len(plots))
}

// func PrintGarden(rocks Rocks, plots Plots) {
// 	var maxX, maxY int
// 	for _, rock := range rocks {
// 		if rock.X > maxX {
// 			maxX = rock.X
// 		}
// 		if rock.Y > maxY {
// 			maxY = rock.Y
// 		}
// 	}

// 	for _, plot := range plots {
// 		if plot.X > maxX {
// 			maxX = plot.X
// 		}
// 		if plot.Y > maxY {
// 			maxY = plot.Y
// 		}
// 	}

// 	for y := 0; y <= maxY; y++ {
// 		for x := 0; x <= maxX; x++ {
// 			var isRock bool
// 			for _, rock := range rocks {
// 				if rock.X == x && rock.Y == y {
// 					isRock = true
// 					break
// 				}
// 			}

// 			if isRock {
// 				fmt.Print("#")
// 				continue
// 			}

// 			var isPlot bool
// 			for _, plot := range plots {
// 				if plot.X == x && plot.Y == y {
// 					isPlot = true
// 					break
// 				}
// 			}

// 			if isPlot {
// 				fmt.Print("O")
// 				continue
// 			}

// 			fmt.Print(".")
// 		}
// 		fmt.Println()
// 	}
// }
