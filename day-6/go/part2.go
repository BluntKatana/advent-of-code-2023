package day6

import (
	"fmt"
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

	var total_combos int = 1

	var button_time = game_time/2 - 1
	var speed = game_time / 2

	for button_time*speed > distance {
		total_combos++
		speed += 1
		button_time -= 1
	}

	total_combos *= 2
	total_combos -= 1

	// print the time it took to run
	fmt.Println(time.Since(start))

	return fmt.Sprint(total_combos)
}

// 37286485
