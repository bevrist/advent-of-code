package main

import (
	"2022/input"
	"testing"
)

var testInput []string = []string{
	"1000",
	"2000",
	"3000",
	"",
	"4000",
	"",
	"5000",
	"6000",
	"",
	"7000",
	"8000",
	"9000",
	"",
	"10000",
}

func TestPart1(t *testing.T) {
	expected := 24000
	parsed := parseInput(testInput)

	result := part1(parsed)
	if result != expected {
		t.Errorf("got: %d, want: %d", result, expected)
	}
}

func TestPart1Full(t *testing.T) {
	expected := 71471

	input := input.GetStringInput("input.txt")
	parsed := parseInput(input)

	result := part1(parsed)
	if result != expected {
		t.Errorf("got: %d, want: %d", result, expected)
	}
}

func TestPart2(t *testing.T) {
	expected := 45000
	parsed := parseInput(testInput)

	result := part2(parsed)
	if result != expected {
		t.Errorf("got: %d, want: %d", result, expected)
	}
}

func TestPart2Full(t *testing.T) {
	expected := 211189

	input := input.GetStringInput("input.txt")
	parsed := parseInput(input)

	result := part2(parsed)
	if result != expected {
		t.Errorf("got: %d, want: %d", result, expected)
	}
}
