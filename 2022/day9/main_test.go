package main

import (
	"2022/input"
	"testing"
)

var testInput []string = []string{
	"R 4",
	"U 4",
	"L 3",
	"D 1",
	"R 4",
	"D 1",
	"L 5",
	"R 2",
}

func TestPart1(t *testing.T) {
	expected := 13

	result := part1(testInput)
	if result != expected {
		t.Errorf("got: %d, want: %d", result, expected)
	}
}

func TestPart1Full(t *testing.T) {
	expected := 6337
	input := input.GetStringInput("input.txt")

	result := part1(input)
	if result != expected {
		t.Errorf("got: %d, want: %d", result, expected)
	}
}

// func TestPart2(t *testing.T) {
// 	expected := 8

// 	result := part2(testInput)
// 	if result != expected {
// 		t.Errorf("got: %d, want: %d", result, expected)
// 	}
// }

// func TestPart2Full(t *testing.T) {
// 	expected := 672280
// 	input := input.GetStringInput("input.txt")

// 	result := part2(input)
// 	if result != expected {
// 		t.Errorf("got: %d, want: %d", result, expected)
// 	}
// }
