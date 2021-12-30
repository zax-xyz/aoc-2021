package lib

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func ReadLines(filename string) []string {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	var nums []string

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		nums = append(nums, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return nums
}

func ReadInts(filename string) [][]int {
	lines := ReadLines(filename)

	ints := make([][]int, len(lines))

	for i, line := range lines {
		lineInts := make([]int, len(line))
		for j, char := range line {
			num, err := strconv.Atoi(string(char))
			if err != nil {
				log.Fatal("Unexpected non-integer in input file: " + string(char))
			}

			lineInts[j] = num
		}

		ints[i] = lineInts
	}

	return ints
}
