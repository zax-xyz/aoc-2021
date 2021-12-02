package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/zaxutic/aoc-2021/lib"
)

func main() {
	P1()
	P2()
}

func P1() {
	lines := lib.ReadLines("day2p1.txt")
	pos := [2]int{0, 0}
	for _, line := range lines {
		parts := strings.Fields(line)
		distance, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatal(err)
		}

		switch direction := parts[0]; direction {
		case "forward":
			pos[0] += distance
		case "down":
			pos[1] += distance
		case "up":
			pos[1] -= distance
		default:
			log.Fatal("invalid direction: " + direction)
		}
	}

	fmt.Println(pos[0] * pos[1])
}

func P2() {

}
