package day16

import (
	"fmt"
	"time"
)

type Day16 struct{}

func (d Day16) Part1(filename *string) string {
	var start = time.Now()

	// var content, _ = os.ReadFile(*filename)
	// var lines = strings.Split(string(content), "\n")

	fmt.Println(time.Since(start))
	return fmt.Sprint(0)
}
