package day14

import (
	"fmt"
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

func (d Day14) Part2(filename *string) string {
	var start = time.Now()

	// var content, _ = os.ReadFile(*filename)
	// var lines = strings.Split(string(content), "\n")

	var sum = 0

	fmt.Println(time.Since(start))
	return fmt.Sprint(sum)
}
