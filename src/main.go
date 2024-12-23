package main

import (
	"fmt"
	"os"

	"github.com/fermeaux/aoc-2024/solutions"
)

func main() {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println("Usage: go run main.go <day> <part>")
	}

	day := os.Args[1]
	part := "1"
	if len(os.Args) == 3 {
		if os.Args[2] != "1" && os.Args[2] != "2" {
			fmt.Println("Part should be 1 or 2")
			return
		}
		part = os.Args[2]
	}

	fmt.Println("Day: " + day + " | Part: " + part)

	solutions.Run(day, part)
}
