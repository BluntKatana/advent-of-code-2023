package day19

import (
	"fmt"
	"time"
)

func (d Day19) Part2(filename *string) string {
	var start = time.Now()

	// var content, _ = os.ReadFile(*filename)
	// var lines = strings.Split(string(content), "\n")

	fmt.Println(time.Since(start))
	return fmt.Sprint(0)
}
