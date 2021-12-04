package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/zaxutic/aoc-2021/lib"
)

func main() {
	p1()
	p2()
}

func p1() {
	lines := lib.ReadLines("day4.txt")

	drawn := []int{}
	for _, field := range strings.Split(lines[0], ",") {
		num, _ := strconv.Atoi(field)
		drawn = append(drawn, num)
	}

	boards, positions := readBoards(lines)

	for _, num := range drawn {
		for i, board := range boards {
			pos, ok := positions[i][num]
			if !ok || board[pos[0]][pos[1]] < 0 {
				continue
			}

			board[pos[0]][pos[1]] = ^board[pos[0]][pos[1]]
			if checkBoard(board, pos[0], pos[1]) {
				fmt.Println(getScore(board, num))
				return
			}
		}
	}
}

func readBoards(lines []string) (boards [][][]int, positions []map[int][2]int) {
	boards = [][][]int{}
	// O(1) lookup for positions within board
	positions = []map[int][2]int{}

	for i := 2; i < len(lines); i += 6 {
		board := [][]int{}
		boardMap := make(map[int][2]int)

		for r := 0; r < 5; r++ {
			row := []int{}
			for c, field := range strings.Fields(lines[i + r]) {
				num, _ := strconv.Atoi(field)
				row = append(row, num)
				boardMap[num] = [2]int{r, c}
			}
			board = append(board, row)
		}

		boards = append(boards, board)
		positions = append(positions, boardMap)
	}

	return boards, positions
}

// instead of checking the whole board (O(n^2)), we only check the relevant row and column (O(n))
func checkBoard(board [][]int, r int, c int) bool {
	win := true
	for _, num := range board[r] {
		if num > 0 {
			win = false
			break
		}
	}

	if win {
		return true
	}

	win = true
	for _, row := range board {
		if row[c] > 0 {
			win = false
			break
		}
	}

	return win
}

func getScore(board[][]int, num int) int {
	score := 0
	for _, row := range board {
		for _, num := range row {
			if num > 0 {
				score += num
			}
		}
	}

	return score * num
}

func p2() {
}
