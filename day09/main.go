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

func isLow(lines [][]int, row int, col int) bool {
	height := lines[row][col]
	return (
		(row == len(lines) - 1 || height < lines[row + 1][col]) &&
		(row == 0 || height < lines[row - 1][col]) &&
		(col == len(lines[0]) - 1 || height < lines[row][col + 1]) &&
		(col == 0 || height < lines[row][col - 1]))
}

func p2() int {
	return 0
}
