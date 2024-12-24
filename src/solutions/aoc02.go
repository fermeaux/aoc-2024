package solutions

import (
	"fmt"
	"strconv"
	"strings"
)

func getDay2() Day {
	return Day{
		Part1: func(input string) {
			lines := strings.Split(input, "\n")
			total := 0
			replacer := strings.NewReplacer("\r", "")
			for _, line := range lines {
				if line == "" {
					continue
				}

				rawLevels := strings.Split(line, " ")
				levels := []int{}
				for _, rawLevel := range rawLevels {
					level, err := strconv.Atoi(replacer.Replace(rawLevel))
					if err != nil {
						continue
					}
					levels = append(levels, level)
				}

				if len(levels) >= 2 && levels[0] > levels[1] {
					for i := 0; i < len(levels)/2; i++ {
						tmp := levels[i]
						levels[i] = levels[len(levels)-i-1]
						levels[len(levels)-i-1] = tmp
					}
				}

				last := 0
				safe := true
				for i, level := range levels {
					if i == 0 {
						last = level
						continue
					}

					if level-last >= 1 && level-last <= 3 {
						last = level
					} else {
						safe = false
						break
					}
				}

				if safe {
					total++
				}
			}

			fmt.Println(total)
		},
		Part2: func(input string) {
			lines := strings.Split(input, "\n")
			total := 0
			replacer := strings.NewReplacer("\r", "")
			for _, line := range lines {
				if line == "" {
					continue
				}

				rawLevels := strings.Split(line, " ")
				levels := []int{}
				for _, rawLevel := range rawLevels {
					level, err := strconv.Atoi(replacer.Replace(rawLevel))
					if err != nil {
						continue
					}
					levels = append(levels, level)
				}

				isSafe := false
				for i := range levels {
					tmp := RemoveAt(levels, i)
					if ReportIsValid(tmp) {
						isSafe = true
						break
					}
				}

				if isSafe {
					total++
				}
			}

			fmt.Println(total)
		},
	}
}

func ReportIsValid(levels []int) bool {
	if len(levels) >= 2 && levels[0] > levels[1] {
		for i := 0; i < len(levels)/2; i++ {
			tmp := levels[i]
			levels[i] = levels[len(levels)-i-1]
			levels[len(levels)-i-1] = tmp
		}
	}

	var last int
	for i, level := range levels {
		if i == 0 {
			last = level
			continue
		}

		if level-last >= 1 && level-last <= 3 {
			last = level
		} else {
			return false
		}
	}

	return true
}

func RemoveAt(slice []int, i int) []int {
	tmp := []int{}
	for j := 0; j < len(slice); j++ {
		if j == i {
			continue
		}
		tmp = append(tmp, slice[j])
	}
	return tmp
}
