package day14

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// Tilt south meaning all rocks (O) move up south until either
// they hit the bottom (index = 0) or a fixed rock (#)
func (d *Dish) TiltSouth() {
	// For each row we iterate from top to bottom by at least row-steps
	for step := 0; step < len(*d); step++ {

		// For each row starting at the first we move each rock down south
		for row := len(*d) - 1; row >= 0; row-- {
			for col := 0; col < len((*d)[row]); col++ {
				if (*d)[row][col] == 'O' {
					// Only move down if we are
					// - not at the bottom
					// - the position below us is not a fixed rock
					// - the position below us is not a normal rock
					if !(row == len(*d)-1 || (*d)[row+1][col] == 'O' || (*d)[row+1][col] == '#') {
						// If there is no fixed rock (#) below us, we can move up
						// We move down by replacing the current position with a free space (.)
						// and the position below us with a rock (O)
						(*d)[row] = (*d)[row][:col] + "." + (*d)[row][col+1:]
						(*d)[row+1] = (*d)[row+1][:col] + "O" + (*d)[row+1][col+1:]
					}
				}
			}
		}
	}
}

func (d *Dish) TiltWest() {
	for step := 0; step < len(*d); step++ {
		for row := 0; row < len(*d); row++ {
			for col := 0; col < len((*d)[row]); col++ {
				if (*d)[row][col] == 'O' {
					if !(col == 0 || (*d)[row][col-1] == 'O' || (*d)[row][col-1] == '#') {
						(*d)[row] = (*d)[row][:col] + "." + (*d)[row][col+1:]
						(*d)[row] = (*d)[row][:col-1] + "O" + (*d)[row][col:]
					}
				}
			}
		}
	}
}

func (d *Dish) TiltEast() {
	for step := 0; step < len(*d); step++ {
		for row := 0; row < len(*d); row++ {
			for col := len((*d)[row]) - 1; col >= 0; col-- {
				if (*d)[row][col] == 'O' {
					if !(col == len((*d)[row])-1 || (*d)[row][col+1] == 'O' || (*d)[row][col+1] == '#') {
						(*d)[row] = (*d)[row][:col] + "." + (*d)[row][col+1:]
						(*d)[row] = (*d)[row][:col+1] + "O" + (*d)[row][col+2:]
					}
				}
			}
		}
	}
}

func (d Dish) String() string {
	return strings.Join(d, "")
}

func (d Day14) Part2(filename *string) string {
	var start = time.Now()

	var content, _ = os.ReadFile(*filename)
	var lines = strings.Split(string(content), "\n")

	var dish Dish

	for _, line := range lines {
		dish = append(dish, line)
	}

	cache := make(map[string]int)

	cycles := 1_000_000_000
	// Iterate through the cycles
	for curr_cycle := 0; curr_cycle < cycles; curr_cycle++ {

		dish.TiltNorth()
		dish.TiltWest()
		dish.TiltSouth()
		dish.TiltEast()

		// If we find a loop, we can calculate the remaining cycles
		if _, ok := cache[dish.String()]; ok {
			var cycle = curr_cycle - cache[dish.String()]
			var remaining = (cycles - curr_cycle) % cycle // remaining cycles

			if remaining == 0 {
				break
			}
		} else {
			cache[dish.String()] = curr_cycle
		}
	}

	dish.Print()

	fmt.Println(time.Since(start))
	return fmt.Sprint(dish.TotalLoadNorth())
}
