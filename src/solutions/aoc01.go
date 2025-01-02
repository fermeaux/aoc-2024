package solutions

import (
	"fmt"
	"math"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/fermeaux/aoc-2024/src/utils"
)

func getDay1() Day {
	return Day{
		Part1: func(input string) {
			lines := ParseLines(input)
			regex := regexp.MustCompile(`(\d+)\s+(\d+)`)
			leftArray := []int{}
			rightArray := []int{}
			for _, line := range lines {
				match := regex.FindStringSubmatch(line)
				if len(match) == 0 || match == nil {
					continue
				}

				left, err := strconv.Atoi(match[1])
				if err != nil {
					fmt.Println(err)
					return
				}
				leftArray = append(leftArray, left)

				right, err := strconv.Atoi(match[2])
				if err != nil {
					fmt.Println(err)
					return
				}
				rightArray = append(rightArray, right)
			}

			sort.Ints(leftArray)
			sort.Ints(rightArray)

			total := 0
			for i := range leftArray {
				left := leftArray[i]
				right := rightArray[i]
				distance := int(math.Abs(float64(left - right)))
				total += distance
			}

			fmt.Println(total)
		},
		Part2: func(input string) {
			lines := strings.Split(input, "\n")
			regex := regexp.MustCompile(`(\d+)\s+(\d+)`)
			leftMap := make(map[string]int)
			rightMap := make(map[string]int)
			for _, line := range lines {
				match := regex.FindStringSubmatch(line)
				if len(match) == 0 || match == nil {
					continue
				}

				left := match[1]
				_, exists := leftMap[left]
				if !exists {
					leftMap[left] = 0
				}
				leftMap[left]++

				right := match[2]
				_, exists = rightMap[right]
				if !exists {
					rightMap[right] = 0
				}
				rightMap[right]++
			}

			total := 0
			for key, left := range leftMap {
				right, exists := rightMap[key]
				if !exists {
					right = 0
				}
				nb, _ := strconv.Atoi(key)
				total += left * right * nb
			}

			fmt.Println(total)
		},
	}
}
