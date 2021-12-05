package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/zaxutic/aoc-2021/lib"
)

type point struct {
	x, y int
}

func main() {
	p1()
	p2()
}

func minMax(x, y int) (min, max int) {
	if x < y {
		return x, y
	}

	return y, x
}

func p1() {
	lines := lib.ReadLines("day5.txt")

	grid := [1000][1000]int{}
	dangerous := 0

	for _, line := range lines {
		from, to := parsePoints(line)
		if from.x != to.x && from.y != to.y {
			// points do not create a horizontal or vertical line
			continue
		}

		if from.x != to.x {
			min, max := minMax(from.x, to.x)
			for i := min; i <= max; i++ {
				grid[i][from.y]++
				if grid[i][from.y] == 2 {
					dangerous++
				}
			}
		} else if from.y != to.y {
			min, max := minMax(from.y, to.y)
			for i := min; i <= max; i++ {
				grid[from.x][i]++
				if grid[from.x][i] == 2 {
					dangerous++
				}
			}
		}
	}

	fmt.Println(dangerous)
}

func parsePoints(line string) (from, to point) {
	parts := strings.Split(line, " -> ")

	fromParts := strings.Split(parts[0], ",")
	x1, _ := strconv.Atoi(fromParts[0])
	y1, _ := strconv.Atoi(fromParts[1])

	toParts := strings.Split(parts[1], ",")
	x2, _ := strconv.Atoi(toParts[0])
	y2, _ := strconv.Atoi(toParts[1])

	return point{x1, y1}, point{x2, y2}
}

func p2() {
}
