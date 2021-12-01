package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	p1()
	p2()
}

func p1() {
	f, err := os.Open("p1p1.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	prev := math.MaxInt
	increases := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		i, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}

		if i > prev {
			increases++
		}

		prev = i
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(increases)
}

func p2() {
	nums := readNums("p1p2.txt")

	increases := 0

	WINDOW_SIZE := 3
	windowSum := 0
	for i := 0; i < WINDOW_SIZE; i++ {
		windowSum += nums[i]
	}

	for i := WINDOW_SIZE; i < len(nums); i++ {
		newWindowSum := windowSum + nums[i] - nums[i - WINDOW_SIZE]
		if newWindowSum > windowSum {
			increases++
		}

		windowSum = newWindowSum
	}

	fmt.Println(increases)
}

func readNums(filename string) []int {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	var nums []int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		i, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}

		nums = append(nums, i)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return nums
}
