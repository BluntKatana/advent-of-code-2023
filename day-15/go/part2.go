package day15

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Lens struct {
	Label       string
	FocalLength int
}

type LensMap map[int][]Lens

// Add the lends to the end of the box
func (lm LensMap) Add(l Lens, box int) {
	// Check if entry for box exists
	if _, ok := lm[box]; !ok {
		lm[box] = []Lens{l}
		return
	}

	// If the label is present in the box, replace it
	for i, lens := range lm[box] {
		if lens.Label == l.Label {
			lm[box][i] = l
			return
		}
	}

	// Otherwise, append it
	lm[box] = append(lm[box], l)
}

func (lm LensMap) Remove(l Lens, box int) {
	// Check if entry for box exists
	if _, ok := lm[box]; !ok {
		return
	}

	// If the label is present in the box, remove it
	for i, lens := range lm[box] {
		if lens.Label == l.Label && len(lm[box]) == 1 {
			delete(lm, box)
			return
		} else if lens.Label == l.Label {
			lm[box] = append(lm[box][:i], lm[box][i+1:]...)
			return
		}
	}
}

func (lm LensMap) TotalFocusingPower() int {
	var total_power int = 0
	for box, lenses := range lm {
		for slot, lens := range lenses {
			total_power += (box + 1) * (slot + 1) * lens.FocalLength
		}
	}

	return total_power
}

func (lm LensMap) String() string {
	var s string
	for box, lenses := range lm {
		s += fmt.Sprintf("%d: ", box)
		for _, lens := range lenses {
			s += fmt.Sprintf("%s=%d ", lens.Label, lens.FocalLength)
		}
		s += "\n"
	}

	return s
}

func (d Day15) Part2(filename *string) string {
	var start = time.Now()

	var content, _ = os.ReadFile(*filename)
	var lines = strings.Split(string(content), "\n")
	var steps = strings.Split(lines[0], ",")

	var lens_map LensMap = make(map[int][]Lens)

	for i := 0; i < len(steps); i++ {
		if strings.Contains(steps[i], "-") {
			label_to_remove := strings.Split(steps[i], "-")[0]
			lens_map.Remove(Lens{label_to_remove, 0}, Hash(label_to_remove))
		} else {
			var lens = strings.Split(steps[i], "=")
			var label = lens[0]
			var focal_length, _ = strconv.Atoi(lens[1])
			lens_map.Add(Lens{label, focal_length}, Hash(label))
		}
	}

	fmt.Println(time.Since(start))
	return fmt.Sprint(lens_map.TotalFocusingPower())
}
