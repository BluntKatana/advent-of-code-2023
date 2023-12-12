package day13

import (
	"fmt"
	"time"
)

type Day13 struct{}

func (d Day13) Part1(filename *string) string {
	var start = time.Now()

	// var content, _ = os.ReadFile(*filename)
	// var lines = strings.Split(string(content), "\n")

	fmt.Println(time.Since(start))
	return fmt.Sprint(0)
}
