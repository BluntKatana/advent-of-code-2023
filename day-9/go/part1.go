package day9

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type Day9 struct{}

func (d Day9) Part1(filename *string) string {
	var start = time.Now()

	// Parse input
	var content, _ = os.ReadFile(*filename)
	var lines = strings.Split(string(content), "\n")

	var total = 0

	fmt.Println(lines)

	fmt.Println(time.Since(start))
	return fmt.Sprint(total)
}
