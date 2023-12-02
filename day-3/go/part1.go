package day3

// initialize day 2
type Day3 struct{}

func filename_part1(test_mode bool) string {
	if test_mode {
		return "./day-3/test_part1.txt"
	}
	return "./day-3/input.txt"
}

func (d Day3) Part1(test_mode bool) string {
	// content, _ := os.ReadFile(filename_part1(test_mode))
	// lines := strings.Split(string(content), "\n")

	return "tbd"
}
