package solutions

import (
	"fmt"
	"os"
	"path"
	"runtime"
)

type Day struct {
	Part1 func(string)
	Part2 func(string)
}

func Run(day string, part string) {
	days := map[string]Day{
		// "01": getDay1(),
		"02": getDay2(),
	}

	tmp, ok := days[day]
	if !ok {
		fmt.Println("Day not found")
		return
	}

	_, filename, _, _ := runtime.Caller(1)
	content, err := os.ReadFile(path.Join(path.Dir(filename), "../inputs/aoc-"+day+"/data.txt"))
	if err != nil {
		fmt.Println(err)
		return
	}

	// test, err := os.ReadFile(path.Join(path.Dir(filename), "../inputs/aoc-"+day+"/test.txt"))
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	if part == "1" {
		tmp.Part1(string(content))
	} else if part == "2" {
		tmp.Part2(string(content))
	}
}
