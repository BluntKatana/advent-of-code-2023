package day6

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func (d Day6) Part2(filename *string) string {
	// start clock
	start := time.Now()

	content, _ := os.ReadFile(*filename)
	lines := strings.Split(string(content), "\n")

	re := regexp.MustCompile("[0-9]+")

	var game_time, _ = strconv.Atoi(strings.Join(re.FindAllString(lines[0], -1), ""))
	var distance, _ = strconv.Atoi(strings.Join(re.FindAllString(lines[1], -1), ""))

	// solve the equation
	var ds = math.Sqrt(float64(game_time*game_time - 4*distance))

	var from = math.Floor((float64(game_time) - ds) / 2)
	var to = math.Ceil((float64(game_time) + ds) / 2)

	var total_combos = int(to) - int(from) - 1

	// print the time it took to run
	fmt.Println(time.Since(start))

	return fmt.Sprint(total_combos)
}

// 37286485
