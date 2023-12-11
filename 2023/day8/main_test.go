package main

import (
	"2020/input"
	"testing"
)

var testInput1 []string = []string{
	"RL",
	"",
	"AAA = (BBB, CCC)",
	"BBB = (DDD, EEE)",
	"CCC = (ZZZ, GGG)",
	"DDD = (DDD, DDD)",
	"EEE = (EEE, EEE)",
	"GGG = (GGG, GGG)",
	"ZZZ = (ZZZ, ZZZ)",
}
var testInput2 []string = []string{
	"LLR",
	"",
	"AAA = (BBB, BBB)",
	"BBB = (AAA, ZZZ)",
	"ZZZ = (ZZZ, ZZZ)",
}

func TestPart1_1(t *testing.T) {
	expected := 2

	parsed := parseInput(testInput1)

	result := part1(parsed)
	if result != expected {
		t.Errorf("got: %d, want: %d", result, expected)
	}
}

func TestPart1_2(t *testing.T) {
	expected := 6

	parsed := parseInput(testInput2)

	result := part1(parsed)
	if result != expected {
		t.Errorf("got: %d, want: %d", result, expected)
	}
}

func TestPart1Full(t *testing.T) {
	expected := 18827

	input := input.GetStringInput("input.txt")

	parsed := parseInput(input)

	result := part1(parsed)
	if result != expected {
		t.Errorf("got: %d, want: %d", result, expected)
	}
}

// func TestPart2(t *testing.T) {
// 	expected := 5905

// 	parsed := parseInput(testInput)

// 	result := part2(parsed)
// 	if result != expected {
// 		t.Errorf("got: %d, want: %d", result, expected)
// 	}
// }

// func TestPart2Full(t *testing.T) {
// 	expected := 249356515

// 	input := input.GetStringInput("input.txt")

// 	parsed := parseInput(input)

// 	result := part2(parsed)
// 	if result != expected {
// 		t.Errorf("got: %d, want: %d", result, expected)
// 	}
// }
