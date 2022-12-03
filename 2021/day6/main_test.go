package main

import (
	"2021/input"
	"testing"
)

var testInput = []string{
	"3,4,3,1,2",
}

func TestPart1(t *testing.T) {
	expected := 5934
	parsed := parseInput(testInput)

	result := part1(parsed, 80)
	if result != expected {
		t.Errorf("got: %d, want: %d", result, expected)
	}
}

func TestPart1Full(t *testing.T) {
	input := input.GetStringInput("input.txt")
	parsed := parseInput(input)

	expected := 356190

	result := part1(parsed, 80)
	if result != expected {
		t.Errorf("got: %d, want: %d", result, expected)
	}
}

func TestPart2(t *testing.T) {
	expected := int64(26984457539)
	parsed := parseInput(testInput)

	result := part2(parsed, 256)
	if result != expected {
		t.Errorf("got: %d, want: %d", result, expected)
	}
}

func TestPart2Full(t *testing.T) {
	input := input.GetStringInput("input.txt")
	parsed := parseInput(input)

	expected := int64(1617359101538)
	// 1538744430921 too low

	result := part2(parsed, 256)
	if result != expected {
		t.Errorf("got: %d, want: %d", result, expected)
	}
}
