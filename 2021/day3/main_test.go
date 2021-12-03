package main

import (
	"2021/input"
	"testing"
)

var entries []string = []string{
	"00100",
	"11110",
	"10110",
	"10111",
	"10101",
	"01111",
	"00111",
	"11100",
	"10000",
	"11001",
	"00010",
	"01010",
}

func TestPart1(t *testing.T) {
	expected := 198

	result := part1(entries)
	if result != expected {
		t.Errorf("got: %d, want: %d", result, expected)
	}
}

func TestPart1Full(t *testing.T) {
	input := input.GetStringInput("input.txt")
	expected := 749376

	result := part1(input)
	if result != expected {
		t.Errorf("got: %d, want: %d", result, expected)
	}
}

// func TestPart2(t *testing.T) {
// 	expected := 230

// 	result := part2(entries)
// 	if result != expected {
// 		t.Errorf("got: %d, want: %d", result, expected)
// 	}
// }

// func TestPart2Full(t *testing.T) {
// 	input := input.GetStringInput("input.txt")
// 	expected := 2105273490

// 	result := part2(input)
// 	if result != expected {
// 		t.Errorf("got: %d, want: %d", result, expected)
// 	}
// }
