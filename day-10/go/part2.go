package day10

import (
	"fmt"
	"time"
)

func (d Day10) Part2(filename *string) string {
	var start = time.Now()

	// Parse input
	// var content, _ = os.ReadFile(*filename)
	// var lines = strings.Split(string(content), "\n")

	var total = 0

	fmt.Println(time.Since(start))
	return fmt.Sprint(total)
}
