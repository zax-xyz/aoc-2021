package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/zaxutic/aoc-2021/lib"
)

func main() {
	p1()
	p2()
}

func p1() {
	lines := lib.ReadLines("day7.txt")

	parts := strings.Split(lines[0], ",")
	positions := make([]int, len(parts))

	min := math.MaxInt
	max := 0

	for i, part := range parts {
		positions[i], _ = strconv.Atoi(part)
		if positions[i] < min {
			min = positions[i]
		} else if positions[i] > max {
			max = positions[i]
		}
	}

	minFuel := math.MaxInt
	for i := min; i <= max; i++ {
		fuel := 0
		for _, pos := range positions {
			// S_n = n/2 (1 + n)
			// where S_n = 1 + 2 + ... + n
			n := math.Abs(float64(i - pos))
			fuel += int(n / 2 * (1 + n))
		}

		if fuel < minFuel {
			minFuel = fuel
		}
	}

	fmt.Println(minFuel)
}

func p2() {
}
