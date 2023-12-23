package day20

import (
	"fmt"
	"time"
)

type Module struct {
	Type         int // 0 = broadcaster, 1 = flip-flop (%), 2 = conjuction (&)
	Name         string
	Destinations []string

	// flip-flop => 0 = off, 1 = on
	// conjuction => remembers the last
}

type Day20 struct{}

func (d Day20) Part1(filename *string) string {
	var start = time.Now()

	// var content, _ = os.ReadFile(*filename)
	// var lines = strings.Split(string(content), "\n")

	fmt.Println(time.Since(start))
	return fmt.Sprint(0)
}
