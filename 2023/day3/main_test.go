package main

import (
	"2020/input"
	"testing"
)

var testInput []string = []string{
	"467..114..",
	"...*......",
	"..35..633.",
	"......#...",
	"617*......",
	".....+.58.",
	"..592.....",
	"......755.",
	"...$.*....",
	".664.598..",
}

func TestPart1(t *testing.T) {
	expected := 4361

	parsed := parseInput(testInput)

	result := part1(parsed)
	if result != expected {
		t.Errorf("got: %d, want: %d", result, expected)
	}
}

func TestPart1Full(t *testing.T) {
	expected := 539713

	input := input.GetStringInput("input.txt")

	parsed := parseInput(input)

	result := part1(parsed)
	if result != expected {
		t.Errorf("got: %d, want: %d", result, expected)
	}
}

func TestPart2(t *testing.T) {
	expected := 467835

	parsed := parseInput(testInput)

	result := part2(parsed)
	if result != expected {
		t.Errorf("got: %d, want: %d", result, expected)
	}
}

// func TestPart2Full(t *testing.T) {
// 	expected := 63307

// 	input := input.GetStringInput("input.txt")

// 	parsed := parseInput(input)

// 	result := part2(parsed)
// 	if result != expected {
// 		t.Errorf("got: %d, want: %d", result, expected)
// 	}
// }
