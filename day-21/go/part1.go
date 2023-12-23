package day21

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type Day21 struct{}

type Garden [][]rune

func (g Garden) String() string {
	var result string
	for _, row := range g {
		result += string(row) + "\n"
	}
	return result
}

func (g *Garden) Step() {
	var newGarden Garden
	for _, row := range *g {
		var newRow []rune
		for _, char := range row {
			if char == 'S' || char == 'O' {
				newRow = append(newRow, '.')
			} else {
				newRow = append(newRow, char)
			}
		}
		newGarden = append(newGarden, newRow)
	}

	// Loop through garden and for every S or O place the O in the new garden
	// in each direction
	for y, row := range *g {
		for x, char := range row {
			if char == 'S' || char == 'O' {
				// place O in new garden in each direction
				for _, dir := range [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
					var newX = x + dir[0]
					var newY = y + dir[1]
					if newX >= 0 && newX < len(row) && newY >= 0 && newY < len(*g) && newGarden[newY][newX] == '.' {
						newGarden[newY][newX] = 'O'
					}
				}
			}
		}
	}

	// update garden
	*g = newGarden
}

func (g Garden) Count() int {
	var count int
	for _, row := range g {
		for _, char := range row {
			if char == 'O' {
				count++
			}
		}
	}
	return count
}

func (d Day21) Part1(filename *string) string {
	var start = time.Now()

	var content, _ = os.ReadFile(*filename)
	var lines = strings.Split(string(content), "\n")

	var garden Garden
	for _, line := range lines {
		var row []rune
		for _, char := range line {
			row = append(row, char)
		}
		garden = append(garden, row)
	}

	steps := 64

	for i := 0; i < steps; i++ {
		garden.Step()
		fmt.Printf("Step %d\n", i+1)
	}

	fmt.Println(garden)

	fmt.Println(time.Since(start))
	return fmt.Sprint(garden.Count())
}
