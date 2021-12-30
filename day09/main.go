package main

import (
	"fmt"

	"github.com/zaxutic/aoc-2021/lib"
)

func main() {
	fmt.Println(p1())
	fmt.Println(p2())
}

func p1() int {
	lines := lib.ReadInts("day9.txt")

	sum := 0
	for i, row := range lines {
		for j, height := range row {
			if isLow(lines, i, j) {
				sum += height + 1
			}
		}
	}

	return sum
}

func isLow(lines [][]int, row, col int) bool {
	height := lines[row][col]
	return ((row == len(lines)-1 || height < lines[row+1][col]) &&
		(row == 0 || height < lines[row-1][col]) &&
		(col == len(lines[0])-1 || height < lines[row][col+1]) &&
		(col == 0 || height < lines[row][col-1]))
}

func p2() int {
	lines := lib.ReadInts("day9.txt")
	biggestThree := [3]int{}

	for i, line := range lines {
		for j := range line {
			size := dfs(lines, i, j)

			min := biggestThree[0]
			minI := 0
			for i, n := range biggestThree {
				if n < min {
					min = n
					minI = i
				}
			}

			if size > min {
				biggestThree[minI] = size
			}
		}
	}

	return biggestThree[0] * biggestThree[1] * biggestThree[2]
}

func dfs(lines [][]int, row, col int) int {
	if lines[row][col] == 9 {
		return 0
	}

	lines[row][col] = 9

	count := 1
	directions := [][2]int{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}
	for _, direction := range directions {
		r := row + direction[0]
		c := col + direction[1]
		if r < 0 || r > len(lines)-1 ||
			c < 0 || c > len(lines[0])-1 {
			continue
		}

		count += dfs(lines, r, c)
	}

	return count
}
