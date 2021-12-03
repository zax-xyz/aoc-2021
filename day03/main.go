package main

import (
	"fmt"

	"github.com/zaxutic/aoc-2021/lib"
)

func main() {
	p1()
	p2()
}

func p1() {
	lines := lib.ReadLines("day3.txt")
	bits := mostCommonBits(lines)

	gamma := 0
	for _, bit := range bits {
		gamma = (gamma << 1) + bit
	}

	epsilon := gamma ^ ((1 << 12) - 1)

	fmt.Println(gamma * epsilon)
}

func mostCommonBits(lines []string) []int {
	bits := make([]int, len(lines[0]))

	for _, line := range lines {
		for i, bit := range line {
			bits[i] += int(bit - '0')
		}
	}

	for i, bit := range bits {
		if bit >= (len(lines) + 1) / 2 {
			bits[i] = 1
		} else {
			bits[i] = 0
		}
	}

	return bits
}

func p2() {
	lines := lib.ReadLines("day3.txt")
	bits := mostCommonBits(lines)

	oxygen := bitstringToInt(filter(lines, bits, false))
	co2 := bitstringToInt(filter(lines, bits, true))

	fmt.Println(oxygen * co2)
}

// flip changes the behaviour from most common bit to least common
func filter(lines []string, bits []int, flip bool) string {
	linesCopy := make([]string, len(lines))
	copy(linesCopy, lines)
	var filtered []string

	for i := 0; len(filtered) != 1; i++ {
		// filter by lines that match the most common bit at i
		filtered = []string{}
		for _, line := range linesCopy {
			bit := bits[i]
			if flip {
				bit = 1 - bit
			}

			if int(line[i] - '0') == bit {
				filtered = append(filtered, line)
			}
		}

		linesCopy = make([]string, len(filtered))
		copy(linesCopy, filtered)

		// we consider the most common bit within each filtered subset of lines
		// we create, not the whole list
		bits = mostCommonBits(filtered)
	}

	return filtered[0]
}

func bitstringToInt(bits string) int {
	res := 0
	for _, bit := range bits {
		res = (res << 1) + int(bit - '0')
	}

	return res
}
