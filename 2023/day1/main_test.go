package main

import (
	"2022/input"
	"testing"
)

var testInput []string = []string{
	"1abc2",
	"pqr3stu8vwx",
	"a1b2c3d4e5f",
	"treb7uchet",
}

func TestPart1(t *testing.T) {
	expected := 142
	parsed := parseInput(testInput)

	result := part1(parsed)
	if result != expected {
		t.Errorf("got: %d, want: %d", result, expected)
	}
}

func TestPart1Full(t *testing.T) {
	expected := 54597

	input := input.GetStringInput("input.txt")
	parsed := parseInput(input)

	result := part1(parsed)
	if result != expected {
		t.Errorf("got: %d, want: %d", result, expected)
	}
}

func TestPart2(t *testing.T) {
	expected := 281
	parsed := parseInput(testInput)

	result := part1(parsed)
	if result != expected {
		t.Errorf("got: %d, want: %d", result, expected)
	}
}

// func TestPart2Full(t *testing.T) {
// 	expected := 211189

// 	input := input.GetStringInput("input.txt")
// 	parsed := parseInput(input)

// 	result := part1(parsed)
// 	if result != expected {
// 		t.Errorf("got: %d, want: %d", result, expected)
// 	}
// }
