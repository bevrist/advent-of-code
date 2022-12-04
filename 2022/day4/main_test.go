package main

import (
	"2022/input"
	"testing"
)

var testInput []string = []string{
	"2-4,6-8",
	"2-3,4-5",
	"5-7,7-9",
	"2-8,3-7",
	"6-6,4-6",
	"2-6,4-8",
}

func TestPart1(t *testing.T) {
	expected := 2

	result := part1(testInput)
	if result != expected {
		t.Errorf("got: %d, want: %d", result, expected)
	}
}

func TestPart1Full(t *testing.T) {
	expected := 485
	input := input.GetStringInput("input.txt")

	result := part1(input)
	if result != expected {
		t.Errorf("got: %d, want: %d", result, expected)
	}
}

func TestPart2(t *testing.T) {
	expected := 4

	result := part2(testInput)
	if result != expected {
		t.Errorf("got: %d, want: %d", result, expected)
	}
}

func TestPart2Full(t *testing.T) {
	expected := 857
	input := input.GetStringInput("input.txt")

	result := part2(input)
	if result != expected {
		t.Errorf("got: %d, want: %d", result, expected)
	}
}
