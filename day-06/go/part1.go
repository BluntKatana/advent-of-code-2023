package day6

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Day6 struct{}

func (d Day6) Part1(filename *string) string {
	content, _ := os.ReadFile(*filename)
	lines := strings.Split(string(content), "\n")

	var times []string = strings.Fields(strings.Trim(strings.Split(lines[0], ":")[1], " "))
	var distances []string = strings.Fields(strings.Trim(strings.Split(lines[1], ":")[1], " "))

	fmt.Println(times)
	fmt.Println(distances)

	var total_combos int = 1

	for time_idx, time_str := range times {
		var distance, _ = strconv.Atoi(distances[time_idx])
		var time, _ = strconv.Atoi(time_str)

		var button_time = time
		var speed = 0

		var win_times int = 0

		for button_time >= 0 {
			if button_time*speed > distance {
				// fmt.Println(button_time, speed, distance)
				win_times++
			}

			speed += 1
			button_time -= 1
		}

		fmt.Println(time, distance, win_times)

		total_combos *= win_times
	}

	return fmt.Sprint(total_combos)
}
