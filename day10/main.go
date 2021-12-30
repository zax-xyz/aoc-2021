package main

import (
	"fmt"
	"log"
	"sort"

	"github.com/zaxutic/aoc-2021/lib"
)

var match = map[rune]rune{
	')': '(',
	']': '[',
	'}': '{',
	'>': '<',
}

var errorPoints = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

func main() {
	fmt.Println(p1())
	fmt.Println(p2())
}

func p1() int {
	lines := lib.ReadLines("day10.txt")
	points := 0

	for _, line := range lines {
		brackets := []rune{}

		for _, char := range line {
			if char == '(' || char == '[' || char == '{' || char == '<' {
				brackets = append(brackets, char)
			} else if opening, ok := match[char]; ok {
				if last := len(brackets) - 1; opening == brackets[last] {
					brackets = brackets[:last]
				} else {
					points += errorPoints[char]
					break
				}
			} else {
				log.Fatal("Unexpected character: " + string(char))
			}
		}
	}

	return points
}

var autocompletePoints = map[rune]int{
	'(': 1,
	'[': 2,
	'{': 3,
	'<': 4,
}

func p2() int {
	lines := lib.ReadLines("day10.txt")
	scores := []int{}

	for _, line := range lines {
		brackets := []rune{}
		points := 0
		corrupt := false

		for _, char := range line {
			if char == '(' || char == '[' || char == '{' || char == '<' {
				brackets = append(brackets, char)
			} else if opening, ok := match[char]; ok {
				if last := len(brackets) - 1; opening == brackets[last] {
					brackets = brackets[:last]
				} else {
					corrupt = true
					break
				}
			} else {
				log.Fatal("Unexpected character: " + string(char))
			}
		}

		if corrupt {
			continue
		}

		for i := len(brackets) - 1; i >= 0; i-- {
			points *= 5
			points += autocompletePoints[brackets[i]]
		}

		scores = append(scores, points)
	}

	sort.Ints(scores)
	return scores[len(scores)/2]
}
