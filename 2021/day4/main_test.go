package main

import (
	"2021/input"
	"testing"
)

var numbers []int = []int{
	7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1,
}

var boards []Board = []Board{
	{[5][5]int{
		{22, 13, 17, 11, 0},
		{8, 2, 23, 4, 24},
		{21, 9, 14, 16, 7},
		{6, 10, 3, 18, 5},
		{1, 12, 20, 15, 19},
	}, [5][5]bool{}},
	{[5][5]int{
		{3, 15, 0, 2, 22},
		{9, 18, 13, 17, 5},
		{19, 8, 7, 25, 23},
		{20, 11, 10, 24, 4},
		{14, 21, 16, 12, 6},
	}, [5][5]bool{}},
	{[5][5]int{
		{14, 21, 17, 24, 4},
		{10, 16, 15, 9, 19},
		{18, 8, 23, 26, 20},
		{22, 11, 13, 6, 5},
		{2, 0, 12, 3, 7},
	}, [5][5]bool{}},
}

func TestPart1(t *testing.T) {
	expected := 4512

	result := part1(numbers, boards)
	if result != expected {
		t.Errorf("got: %d, want: %d", result, expected)
	}
}

func TestPart1Full(t *testing.T) {
	input := input.GetStringInput("input.txt")
	numbers = getBingoNums(input[0])
	boards = getBingoBoards(input[1:])

	expected := 4662

	result := part1(numbers, boards)
	if result != expected {
		t.Errorf("got: %d, want: %d", result, expected)
	}
}

// func TestPart2(t *testing.T) {
// 	expected := 230

// 	result := part2(entries)
// 	if result != expected {
// 		t.Errorf("got: %d, want: %d", result, expected)
// 	}
// }

// func TestPart2Full(t *testing.T) {
// 	input := input.GetStringInput("input.txt")
// 	expected := 2105273490

// 	result := part2(input)
// 	if result != expected {
// 		t.Errorf("got: %d, want: %d", result, expected)
// 	}
// }
