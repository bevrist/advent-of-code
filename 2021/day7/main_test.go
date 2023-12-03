package main

import (
	"2021/input"
	"testing"
)

var testInput = []string{
	"16,1,2,0,4,2,7,1,2,14",
}

func TestPart1(t *testing.T) {
	expected := 37
	parsed := parseInput(testInput)

	result := part1(parsed)
	if result != expected {
		t.Errorf("got: %d, want: %d", result, expected)
	}
}

func TestPart1Full(t *testing.T) {
	input := input.GetStringInput("input.txt")
	parsed := parseInput(input)

	expected := 345197

	result := part1(parsed)
	if result != expected {
		t.Errorf("got: %d, want: %d", result, expected)
	}
}

func TestPart2(t *testing.T) {
	expected := 168
	parsed := parseInput(testInput)

	result := part2(parsed)
	if result != expected {
		t.Errorf("got: %d, want: %d", result, expected)
	}
}

func TestPart2Full(t *testing.T) {
	input := input.GetStringInput("input.txt")
	parsed := parseInput(input)

	expected := 963616060

	result := part2(parsed)
	if result != expected {
		t.Errorf("got: %d, want: %d", result, expected)
	}
}
