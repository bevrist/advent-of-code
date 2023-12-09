package main

import (
	"2020/input"
	"testing"
)

var testInput []string = []string{
	"32T3K 765",
	"T55J5 684",
	"KK677 28",
	"KTJJT 220",
	"QQQJA 483",
}

func TestPart1(t *testing.T) {
	expected := 6440

	parsed := parseInput(testInput)

	result := part1(parsed)
	if result != expected {
		t.Errorf("got: %d, want: %d", result, expected)
	}
}

func TestPart1Full(t *testing.T) {
	expected := 246795406

	input := input.GetStringInput("input.txt")

	parsed := parseInput(input)

	result := part1(parsed)
	if result != expected {
		t.Errorf("got: %d, want: %d", result, expected)
	}
}

func TestPart2(t *testing.T) {
	expected := 5905

	parsed := parseInput(testInput)

	result := part2(parsed)
	if result != expected {
		t.Errorf("got: %d, want: %d", result, expected)
	}
}

func TestPart2Full(t *testing.T) {
	expected := 249356515

	input := input.GetStringInput("input.txt")

	parsed := parseInput(input)

	result := part2(parsed)
	if result != expected {
		t.Errorf("got: %d, want: %d", result, expected)
	}
}
