package solutions

import "fmt"

func getDay1() Day {
	return Day{
		Part1: func(input string) {
			fmt.Println("Day 1 Part 1")
			fmt.Println(input)
		},
		Part2: func(input string) {
			fmt.Println("Day 1 Part 2")
		},
	}
}
