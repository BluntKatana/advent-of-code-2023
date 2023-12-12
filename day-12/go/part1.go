package day12

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Day12 struct{}

func (d Day12) Part1(filename *string) string {
	var start = time.Now()

	var content, _ = os.ReadFile(*filename)
	var lines = strings.Split(string(content), "\n")

	var totalArrangements = 0

	for _, line := range lines {
		split_line := strings.Split(line, " ")
		records := split_line[0]
		unparsedGroups := strings.Split(split_line[1], ",")
		groups := make([]int, len(unparsedGroups))

		for i, group := range unparsedGroups {
			group, _ := strconv.Atoi(group)
			groups[i] = group
		}

		arrangements := Arrangements(0, 0, records, groups)

		fmt.Println(records, groups, arrangements)
		totalArrangements += arrangements
	}

	fmt.Println(time.Since(start))
	return fmt.Sprint(totalArrangements)
}

func Arrangements(recordIdx, groupIdx int, record string, groups []int) int {
	// If we reach the end of the record
	if recordIdx >= len(record) {
		// ...and we are not at the end of the groups
		if groupIdx < len(groups) {
			return 0
		}
		// ...and we are at the end of the groups
		return 1
	}

	total := 0

	// If we encounter an OPERATIONAL spring we move to the next record and return
	if record[recordIdx] == '.' {
		total += Arrangements(recordIdx+1, groupIdx, record, groups)
		return total
	}

	// If we encounter an UNKNOWN spring we try both
	// - skipping this record, as if it were not broken
	// - continuing with this record, as if it were BROKEN
	if record[recordIdx] == '?' {
		total += Arrangements(recordIdx+1, groupIdx, record, groups)
	}

	// Check boundary for groups
	if groupIdx < len(groups) {
		// check if the current group fits starting from now
		count := 0
		for nextRecordIdx := recordIdx; nextRecordIdx < len(record); nextRecordIdx++ {
			// stop counting when we encounter an OPERATIONAL spring
			if record[nextRecordIdx] == '.' {
				break
			}

			// stop counting when the count exceeds the current group
			if count > groups[groupIdx] {
				break
			}

			// stop counting when we encounter an UNKOWN spring and we are at the count of the current group
			if record[nextRecordIdx] == '?' && count == groups[groupIdx] {
				break
			}

			// for all cases else
			count++
		}

		// if the count is equal to the length of the current group
		if count == groups[groupIdx] {
			// ... and the next record is not a BROKEN spring
			if recordIdx+count < len(record) && record[recordIdx+count] != '#' {
				// ... we jump to the next NOT BROKEN record and group
				total += Arrangements(recordIdx+count+1, groupIdx+1, record, groups)
			} else {
				// ... we jump to the next record and group
				total += Arrangements(recordIdx+count, groupIdx+1, record, groups)
			}
		}
	}

	return total
}
