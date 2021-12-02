package lib

import (
	"bufio"
	"log"
	"os"
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
