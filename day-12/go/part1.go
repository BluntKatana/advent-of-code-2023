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
	fmt.Println(recordIdx, groupIdx, record, groups)

	// If we reach the end of the record
	if recordIdx >= len(record) {
		// ...and we do not have a group to check left
		if groupIdx == len(groups) {
			// ...it is a correct arrangement
			return 1
		}
		return 0
	}

	total := 0

	// If we encounter an OPERATIONAL spring we move to the next record and return the total
	if record[recordIdx] == '.' {
		total += Arrangements(recordIdx+1, groupIdx, record, groups)
		return total
	}

	// If we encounter a BROKEN or UNKNOWN spring
	if groupIdx < len(groups) {
		// check if the current group is correct starting from now
		count := 0
		for nextRecordIdx := recordIdx; nextRecordIdx < len(record); nextRecordIdx++ {
			// we stop count when
			// - we encounter an OPERATIONAL spring
			// - the count exceeds the current group
			// - we encounter an UNKNOWN spring but are at the groups count
			if record[nextRecordIdx] == '.' ||
				count > groups[groupIdx] ||
				record[nextRecordIdx] == '?' && count == groups[groupIdx] {
				break
			}

			// for all cases else
			count++
		}

		// if the count is equal to the current group
		if count == groups[groupIdx] {
			// ... and the next record after the count is not a BROKEN spring
			if recordIdx+count < len(record) && record[recordIdx+count] != '#' {
				// ... we jump to the next spring and leave a gap (hence the +1)
				total += Arrangements(recordIdx+count+1, groupIdx+1, record, groups)
			} else {
				// ... we jump to the next record and group
				total += Arrangements(recordIdx+count, groupIdx+1, record, groups)
			}
		}
	}

	// In addition, we also have to account for skipping an UNKNOWN spring
	if record[recordIdx] == '?' {
		total += Arrangements(recordIdx+1, groupIdx, record, groups)
	}

	return total
}
