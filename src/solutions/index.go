package solutions

import "fmt"

type Day struct {
	Part1 func(string)
	Part2 func(string)
}

type fn func(string)

func run(day string, part string) {
	days := map[string]Day{
		"01": getDay1(),
	}

	tmp, ok := days[day]
	if !ok {
		fmt.Println("Day not found")
		return
	}

	if part == "1" {
		tmp.Part1("")
	} else {
		tmp.Part2("")
	}
}
