package main

import (
	"2020/input"
	"testing"
)

var testInput []string = []string{
	"Time:      7  15   30",
	"Distance:  9  40  200",
}

func TestPart1(t *testing.T) {
	expected := 288

	parsed := parseInput(testInput)

	result := part1(parsed)
	if result != expected {
		t.Errorf("got: %d, want: %d", result, expected)
	}
}

func TestPart1Full(t *testing.T) {
	expected := 781200

	input := input.GetStringInput("input.txt")

	parsed := parseInput(input)

	result := part1(parsed)
	if result != expected {
		t.Errorf("got: %d, want: %d", result, expected)
	}
}

func TestPart2(t *testing.T) {
	expected := 71503

	result := part2(testInput)
	if result != expected {
		t.Errorf("got: %d, want: %d", result, expected)
	}
}

func TestPart2Full(t *testing.T) {
	expected := 49240091

	input := input.GetStringInput("input.txt")

	result := part2(input)
	if result != expected {
		t.Errorf("got: %d, want: %d", result, expected)
	}
}
