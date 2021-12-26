package main

import (
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/zaxutic/aoc-2021/lib"
)

var uniqueDigits = map[int]int {
	2: 1,
	3: 7,
	4: 4,
	7: 8,
}
var digitLen = map[int]int {
	1: 2,
	7: 3,
	4: 4,
	8: 7,
}
var digitWires = map[string]int {
	"abcefg": 0,
	"cf": 1,
	"acdeg": 2,
	"acdfg": 3,
	"bcdf": 4,
	"abdfg": 5,
	"abdefg": 6,
	"acf": 7,
	"abcdefg": 8,
	"abcdfg": 9,
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

type charSet = map[rune]struct{}

func p2() int {
	lines := lib.ReadLines("day8.txt")

	sum := 0

	for _, line := range lines {
		wires := map[int][]charSet{}
		wireSegments := map[rune]rune{}

		parts := strings.Split(line, " | ")
		signals := parts[0]
		output := parts[1]
		for _, part := range strings.Fields(signals) {
			digitWires := charSet{}
			for _, char := range part {
				digitWires[char] = struct{}{}
			}

			wires[len(part)] = append(wires[len(part)], digitWires)
		}

		// find 4, derive bd group
		bd := charSet{}
		for char := range wires[digitLen[4]][0] {
			if _, ok := wires[digitLen[1]][0][char]; !ok {
				bd[char] = struct{}{}
			}
		}

		// find 7, derive a
		for char := range wires[digitLen[7]][0] {
			if _, ok := wires[digitLen[1]][0][char]; !ok {
				wireSegments['a'] = char
				break
			}
		}

		// find 3, derive d and g
		for i, digit := range wires[5] {
			oneCount := 0
			for wire := range digit {
				if _, ok := wires[digitLen[1]][0][wire]; ok {
					oneCount++
				}
			}

			if oneCount != 2 {
				continue
			}

			// at this point, the digit SHOULD be a three
			for wire := range digit {
				_, inOne := wires[digitLen[1]][0][wire]
				_, inFour := wires[digitLen[4]][0][wire]
				if wire != wireSegments['a'] && !inOne {
					if inFour {
						wireSegments['d'] = wire
					} else {
						wireSegments['g'] = wire
					}
				}
			}

			wires[5] = remove(wires[5], i)
		}

		// find 2 and 5, derive e and b
		for _, digit := range wires[5] {
			bdCount := 0
			for wire := range digit {
				if _, ok := bd[wire]; ok {
					bdCount++
				}
			}

			var leftWire rune
			var rightWire rune

			for wire := range digit {
				_, inOne := wires[digitLen[1]][0][wire]
				if inOne {
					rightWire = wire
					continue
				}

				if wire != wireSegments['a'] &&
					wire != wireSegments['d'] &&
					wire != wireSegments['g'] {
					leftWire = wire
				}
			}

			if bdCount == 1 {
				// this is a 2
				wireSegments['e'] = leftWire
				wireSegments['c'] = rightWire
			} else {
				// this is a 5
				wireSegments['b'] = leftWire
				wireSegments['f'] = rightWire
			}
		}

		if len(wireSegments) != 7 {
			log.Fatal("Incorrect length of wireSegments: " + fmt.Sprint(len(wireSegments)))
		}

		out := 0
		for _, digit := range strings.Fields(output) {
			out *= 10

			s := strings.Split(digit, "")
			for i, c := range s {
				for k, v := range wireSegments {
					if v == rune(c[0]) {
						s[i] = string(k)
						break
					}
				}
			}

			sort.Strings(s)
			digit = strings.Join(s, "")
			n, ok := digitWires[digit]
			if !ok {
				log.Fatal("Invalid: " + digit)
			}
			out += n
		}

		sum += out
	}

	return sum
}


func remove(s []charSet, idx int) []charSet {
	s[idx] = s[len(s) - 1]
	return s[:len(s) - 1]
}
