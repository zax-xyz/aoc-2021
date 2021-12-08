package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/zaxutic/aoc-2021/lib"
)

func main() {
	fmt.Println(p1())
	fmt.Println(p2())
}

func p1() int {
	lines := lib.ReadLines("day6.txt")

	fishPlural := []int{}

	for _, fish := range strings.Split(lines[0], ",") {
		fishInt, _ := strconv.Atoi(fish)
		fishPlural = append(fishPlural, fishInt)
	}

	for i := 0; i < 80; i++ {
		for i := range fishPlural {
			fishPlural[i]--
			if fishPlural[i] < 0 {
				fishPlural[i] = 6
				fishPlural = append(fishPlural, 8)
			}
		}
	}

	return len(fishPlural)
}

func p2() int {
	lines := lib.ReadLines("day6.txt")

	fishPlural := map[int]int{}

	for _, fish := range strings.Split(lines[0], ",") {
		fishInt, _ := strconv.Atoi(fish)
		fishPlural[fishInt]++
	}

	for i := 0; i < 256; i++ {
		tmp := fishPlural[0]
		for i := 1; i <= 8; i++ {
			fishPlural[i - 1] = fishPlural[i]
		}

		fishPlural[6] += tmp
		fishPlural[8] = tmp
	}

	count := 0
	for _, fish := range fishPlural {
		count += fish
	}

	return count
}
