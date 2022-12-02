package main

import (
	"2022/input"
	"testing"
)

var testInput []string = []string{
	"A Y",
	"B X",
	"C Z",
}

func TestPart1(t *testing.T) {
	expected := 15

	result := part1(testInput)
	if result != expected {
		t.Errorf("got: %d, want: %d", result, expected)
	}
}

func TestPart1Full(t *testing.T) {
	expected := 12645
	input := input.GetStringInput("input.txt")

	result := part1(input)
	if result != expected {
		t.Errorf("got: %d, want: %d", result, expected)
	}
}

// func TestPart2(t *testing.T) {
// 	expected := 45000
// 	parsed := parseInput(testInput)

// 	result := part2(parsed)
// 	if result != expected {
// 		t.Errorf("got: %d, want: %d", result, expected)
// 	}
// }

// func TestPart2Full(t *testing.T) {
// 	expected := 211189

// 	input := input.GetStringInput("input.txt")
// 	parsed := parseInput(input)

// 	result := part2(parsed)
// 	if result != expected {
// 		t.Errorf("got: %d, want: %d", result, expected)
// 	}
// }
