package day9

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func (d Day9) Part2(filename *string) string {
	start := time.Now()

	// Parse input
	var content, _ = os.ReadFile(*filename)
	var lines = strings.Split(string(content), "\n")

	var total = 0

	fmt.Println(lines)

	fmt.Println(time.Since(start))
	return fmt.Sprint(total)
}
