package day14

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type Day14 struct{}

type Dish []string

func (d Dish) Print() {
	for _, line := range d {
		fmt.Println(line)
	}
}

// Tilt north meaning all rocks (O) move up north until either
// they hit the top (index = 0) or a fixed rock (#)
func (d *Dish) TiltNorth() {
	// For each row we iterate from bottom to top by at least row-steps
	for step := 0; step < len(*d); step++ {

		// For each row starting at the first we move each rock up north
		for row := 0; row < len(*d); row++ {
			for col := 0; col < len((*d)[row]); col++ {
				if (*d)[row][col] == 'O' {
					// Only move up if we are
					// - not at the top
					// - the position above us is not a fixed rock
					// - the position above us is not a normal rock
					if !(row == 0 || (*d)[row-1][col] == 'O' || (*d)[row-1][col] == '#') {
						// If there is no fixed rock (#) above us, we can move up
						// We move up by replacing the current position with a free space (.)
						// and the position above us with a rock (O)
						(*d)[row] = (*d)[row][:col] + "." + (*d)[row][col+1:]
						(*d)[row-1] = (*d)[row-1][:col] + "O" + (*d)[row-1][col+1:]
					}
				}
			}
		}
	}
}

func (d *Dish) TotalLoadNorth() int {
	var load = 0
	for i := 0; i < len(*d); i++ {
		for _, c := range (*d)[i] {
			if c == 'O' {
				load += len(*d) - i
			}
		}
	}
	return load
}

func (d Day14) Part1(filename *string) string {
	var start = time.Now()

	var content, _ = os.ReadFile(*filename)
	var lines = strings.Split(string(content), "\n")

	var dish Dish

	for _, line := range lines {
		dish = append(dish, line)
	}

	dish.TiltNorth()

	var totalLoad = dish.TotalLoadNorth()

	fmt.Println(time.Since(start))
	return fmt.Sprint(totalLoad)
}
