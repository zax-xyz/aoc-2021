package main

import (
	"fmt"
	"strings"

	"github.com/zaxutic/aoc-2021/lib"
)

var uniqueDigits = map[int]int {
	2: 1,
	3: 7,
	4: 4,
	7: 8,
}

func main() {
	fmt.Println(p1())
	fmt.Println(p2())
}

func p1() int {
	lines := lib.ReadLines("day8.txt")

	count := 0

	for _, line := range lines {
		parts := strings.Split(line, " | ")
		output := parts[1]
		for _, digit := range strings.Fields(output) {
			if _, ok := uniqueDigits[len(digit)]; ok {
				count++
			}
		}
	}

	return count
}

func p2() int {
	return 0
}
