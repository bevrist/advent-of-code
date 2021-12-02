package main

import (
	"2020/input"
	"testing"
)

var entries []string = []string{
	"abc",
	"",
	"a",
	"b",
	"c",
	"",
	"ab",
	"ac",
	"",
	"a",
	"a",
	"a",
	"a",
	"",
	"b",
}

func TestPart1(t *testing.T) {
	expected := 11

	result := part1(entries)
	if result != expected {
		t.Errorf("got: %d, want: %d", result, expected)
	}
}

func TestPart1Full(t *testing.T) {
	input := input.GetStringInput("input.txt")
	expected := 6437

	result := part1(input)
	if result != expected {
		t.Errorf("got: %d, want: %d", result, expected)
	}
}

func TestPart2(t *testing.T) {
	expected := 6

	result := part2(entries)
	if result != expected {
		t.Errorf("got: %d, want: %d", result, expected)
	}
}

func TestPart2Full(t *testing.T) {
	input := input.GetStringInput("input.txt")
	expected := 3229

	result := part2(input)
	if result != expected {
		t.Errorf("got: %d, want: %d", result, expected)
	}
}
