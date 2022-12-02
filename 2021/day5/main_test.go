package main

import (
	"2021/input"
	"testing"
)

var testInput = []string{
	"0,9 -> 5,9",
	"8,0 -> 0,8",
	"9,4 -> 3,4",
	"2,2 -> 2,1",
	"7,0 -> 7,4",
	"6,4 -> 2,0",
	"0,9 -> 2,9",
	"3,4 -> 1,4",
	"0,0 -> 8,8",
	"5,5 -> 8,2",
}

func TestPart1(t *testing.T) {
	expected := 5

	parsedInput := parseInput(testInput)

	// fmt.Printf("%+v\n\n", parsedInput)
	// b, _ := json.Marshal(a)
	// fmt.Printf("%v\n", string(b))

	result := part1(parsedInput)
	if result != expected {
		t.Errorf("got: %d, want: %d", result, expected)
	}
}

func TestPart1Full(t *testing.T) {
	input := input.GetStringInput("input.txt")
	parsed := parseInput(input)

	expected := 6461

	result := part1(parsed)
	if result != expected {
		t.Errorf("got: %d, want: %d", result, expected)
	}
}

func TestPart2(t *testing.T) {
	expected := 12

	parsedInput := parseInput(testInput)

	result := part2(parsedInput)
	if result != expected {
		t.Errorf("got: %d, want: %d", result, expected)
	}
}

func TestPart2Full(t *testing.T) {
	input := input.GetStringInput("input.txt")
	parsed := parseInput(input)

	expected := 18065

	result := part2(parsed)
	if result != expected {
		t.Errorf("got: %d, want: %d", result, expected)
	}
}
