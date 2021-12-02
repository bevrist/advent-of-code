package main

import (
	"testing"
)

var example []rune = []rune("FBFBBFFRLR")

var entries []string = []string{
	"BFFFBBFRRR",
	"FFFBBBFRRR",
	"BBFFBBFRLL",
}

func TestParser(t *testing.T) {
	expected := 357

	result := parseInput(example)
	if result != expected {
		t.Errorf("got: %d, want: %d", result, expected)
	}
}

func TestPart1(t *testing.T) {
	expected := 820

	result := part1(entries)
	if result != expected {
		t.Errorf("got: %d, want: %d", result, expected)
	}
}
