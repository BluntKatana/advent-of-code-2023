package day12

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func (d Day12) Part2(filename *string) string {
	var start = time.Now()

	var content, _ = os.ReadFile(*filename)
	var lines = strings.Split(string(content), "\n")

	var totalArrangements = 0

	for _, line := range lines {
		split_line := strings.Split(line, " ")
		records := split_line[0]
		unfoldedRecords := strings.Join([]string{records, records, records, records, records}, "?")

		unfoldedUnparsedGroups := strings.Split(strings.Join([]string{split_line[1], split_line[1], split_line[1], split_line[1], split_line[1]}, ","), ",")
		unfoldedGroups := make([]int, len(unfoldedUnparsedGroups))

		for i, group := range unfoldedUnparsedGroups {
			group, _ := strconv.Atoi(group)
			unfoldedGroups[i] = group
		}

		// initialize cache with size (records x groups + 1) and fill it with -1 values
		var cache [][]int

		for recordIdx := 0; recordIdx < len(unfoldedRecords); recordIdx++ {
			cache = append(cache, make([]int, len(unfoldedGroups)+1))

			for groupIdx := 0; groupIdx < len(unfoldedGroups)+1; groupIdx++ {
				cache[recordIdx][groupIdx] = -1
			}
		}

		arrangements := ArrangementsWithCache(0, 0, unfoldedRecords, unfoldedGroups, cache)

		totalArrangements += arrangements
	}

	fmt.Println(time.Since(start))
	return fmt.Sprint(totalArrangements)
}

func ArrangementsWithCache(recordIdx, groupIdx int, record string, groups []int, cache [][]int) int {
	// If we reach the end of the record
	if recordIdx >= len(record) {
		// ...and we do not have a group to check left
		if groupIdx == len(groups) {
			// ...it is a correct arrangement
			return 1
		}
		return 0
	}

	// Check if the (recordIdx, groupIdx) has already been computed if so use it
	if cache[recordIdx][groupIdx] != -1 {
		return cache[recordIdx][groupIdx]
	}

	total := 0

	// If we encounter an OPERATIONAL spring we move to the next record and return the total
	if record[recordIdx] == '.' {
		total += ArrangementsWithCache(recordIdx+1, groupIdx, record, groups, cache)

		// add the computed total to the cache
		cache[recordIdx][groupIdx] = total
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
				total += ArrangementsWithCache(recordIdx+count+1, groupIdx+1, record, groups, cache)
			} else {
				// ... we jump to the next record and group
				total += ArrangementsWithCache(recordIdx+count, groupIdx+1, record, groups, cache)
			}
		}
	}

	// In addition, we also have to account for skipping an UNKNOWN spring
	if record[recordIdx] == '?' {
		total += ArrangementsWithCache(recordIdx+1, groupIdx, record, groups, cache)
	}

	// add the computed total to the cache
	cache[recordIdx][groupIdx] = total
	return total
}
