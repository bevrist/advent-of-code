// https://adventofcode.com/2021/day/4

package main

import (
	"2021/input"
	"fmt"
	"strconv"
	"strings"
)

type Board struct {
	pos [5][5]int
	chk [5][5]bool
}

// extract list of bingo numbers from input text
func getBingoNums(input string) []int {
	split := strings.Split(input, `,`)
	var numbers []int
	for _, num := range split {
		inum, err := strconv.Atoi(num)
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, inum)
	}
	return numbers
}

// extract bingoBoards from input text
func getBingoBoards(input []string) []Board {
	var boards []Board
	y := 0
	var currBoard *Board = new(Board)
	for _, row := range input {
		row = strings.Join(strings.Fields(row), " ") //sanitize extra spaces in input
		split := strings.Split(row, " ")
		// skip empty rows
		if len(split) <= 1 {
			if y != 0 {
				// export board that has been populated
				boards = append(boards, *currBoard)
				currBoard = new(Board)
				y = 0
			}
			continue
		}
		// extract items and add to board
		for x := 0; x < 5; x++ {
			num, err := strconv.Atoi(split[x])
			if err != nil {
				panic(err)
			}
			currBoard.pos[x][y] = num
		}
		y++
	}
	return boards
}

func main() {
	var boards []Board
	var numbers []int

	input := input.GetStringInput("input.txt")

	// extract bingo numbers
	numbers = getBingoNums(input[0])

	// extract bingo boards
	boards = getBingoBoards(input[1:])

	fmt.Printf("Part1: %d\n", part1(numbers, boards))
	// fmt.Printf("Part2: %d\n", part2(input))
}

// find the winning board from the provided set of numbers
func part1(numbers []int, boards []Board) int {
	// attempt to find a winning streak from the numbers list, starting with ptr position 4 (first 5 numbers)
	for ptr := 4; ptr < len(numbers); ptr++ {
		// fmt.Printf("ptr: %d, num: %d\n", ptr, numbers[ptr])
		// iterate through boards to find a winning board
		for brdNum, board := range boards {
			_ = brdNum
			// for current board, check current subset of numbers list and mark winning board numbers on the `chk` board
			for _, inNum := range numbers[:ptr+1] {
				// compare current board against current number and mark winning numbers on `chk` board
				for x := 0; x < 5; x++ {
					for y := 0; y < 5; y++ {
						if board.pos[x][y] == inNum {
							board.chk[x][y] = true
						}
					}
				}
			}
			// check the `chk` board for a win
			if isWinner(board) {
				fmt.Printf("\nWinning Board Number: %d \n", brdNum)
				printBoard(board)
				// calculate winning number by adding up unselected numbers and multiplying by the final number
				return getWinTotal(board, numbers[ptr])
			}
			// fmt.Printf("\nBoard Number: %d \n", brdNum)
			// printBoard(board)
		}
	}
	// if no boards match, return -999
	return -999
}

func getWinTotal(board Board, finalNum int) int {
	// add up all NON-selected numbers
	total := 0
	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			if !board.chk[x][y] {
				total += board.pos[x][y]
			}
		}
	}
	return total * finalNum
}

// isWinner checks if a board has a winning streak (vertical or horizontal)
func isWinner(board Board) bool {
	//horizontal
	horzWin := 0 //count horizontal winning numbers, win if = 5
	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			// fmt.Printf(" %d ", board.pos[x][y])
			if board.chk[x][y] {
				horzWin++
			}
		}
		// fmt.Println("")
		if horzWin == 5 {
			return true
		} else {
			horzWin = 0
		}
	}
	//vertical
	vertWin := 0 //count vertical winning numbers, win if = 5
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			// fmt.Printf(" %d ", board.pos[x][y])
			if board.chk[x][y] {
				vertWin++
			}
		}
		// fmt.Println("")
		if vertWin == 5 {
			return true
		} else {
			vertWin = 0
		}
	}
	return false
}

func printBoard(board Board) {
	// compare current board against current number and mark winning numbers on `chk` board
	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			selected := ""
			if board.chk[x][y] {
				selected = "*"
			}
			fmt.Printf(" %s%d ", selected, board.pos[x][y])
		}
		fmt.Printf(" \n ")
	}
}
