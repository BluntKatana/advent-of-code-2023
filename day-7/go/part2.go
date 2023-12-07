package day7

import (
	"fmt"
	"time"
)

func (d Day7) Part2(filename *string) string {
	// start clock
	start := time.Now()

	// // read files
	// content, _ := os.ReadFile(*filename)
	// lines := strings.Split(string(content), "\n")

	var result = 0

	fmt.Println(time.Since(start))
	return fmt.Sprint(result)
}
