package day15

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type Day15 struct{}

func (d Day15) Part1(filename *string) string {
	var start = time.Now()

	var content, _ = os.ReadFile(*filename)
	var lines = strings.Split(string(content), "\n")
	var steps = strings.Split(lines[0], ",")

	var hash_sum = 0

	for i := 0; i < len(steps); i++ {
		hash_sum += Hash(steps[i])
	}

	fmt.Println(time.Since(start))
	return fmt.Sprint(hash_sum)
}

func Hash(s string) int {
	var current_value = 0
	for _, c := range s {
		current_value += int(c)
		current_value *= 17
		current_value = current_value % 256
	}

	return current_value
}
