package day10

import (
	"fmt"
	"time"
)

type Day10 struct{}

func (d Day10) Part1(filename *string) string {
	var start = time.Now()

	// Parse input
	// var content, _ = os.ReadFile(*filename)
	// var lines = strings.Split(string(content), "\n")

	var total = 0

	fmt.Println(time.Since(start))
	return fmt.Sprint(total)
}
