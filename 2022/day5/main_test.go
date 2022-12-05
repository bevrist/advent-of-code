package main

import (
	"2022/input"
	"testing"
)

var testInput []string = []string{
	"    [D]    ",
	"[N] [C]    ",
	"[Z] [M] [P]",
	" 1   2   3",
	"",
	"move 1 from 2 to 1",
	"move 3 from 1 to 3",
	"move 2 from 2 to 1",
	"move 1 from 1 to 2",
}

func TestPart1(t *testing.T) {
	expected := "CMZ"

	result := part1(testInput)
	if result != expected {
		t.Errorf("got: %s, want: %s", result, expected)
	}
}

func TestPart1Full(t *testing.T) {
	expected := "TLNGFGMFN"
	input := input.GetStringInput("input.txt")

	result := part1(input)
	if result != expected {
		t.Errorf("got: %s, want: %s", result, expected)
	}
}

func TestPart2(t *testing.T) {
	expected := "MCD"

	result := part2(testInput)
	if result != expected {
		t.Errorf("got: %s, want: %s", result, expected)
	}
}

func TestPart2Full(t *testing.T) {
	expected := "FGLQJCMBD"
	input := input.GetStringInput("input.txt")

	result := part2(input)
	if result != expected {
		t.Errorf("got: %s, want: %s", result, expected)
	}
}
