package utils

import "strings"

func ParseLines(input string) []string {
	replacer := strings.NewReplacer("\r", "")
	lines := strings.Split(input, "\n")
	tmp := []string{}

	for _, line := range lines {
		if line == "" {
			continue
		}

		tmp = append(tmp, replacer.Replace(line))
	}

	return tmp
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
