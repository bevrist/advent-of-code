package main

import (
	"2022/input"
	"testing"
)

var testInput []string = []string{
	"mjqjpqmgbljsphdztnvjfqwrcgsmlb",
}
var testInput1 []string = []string{
	"bvwbjplbgvbhsrlpgdmjqwftvncz",
}
var testInput2 []string = []string{
	"nppdvjthqldpwncqszvftbrmjlhg",
}
var testInput3 []string = []string{
	"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
}
var testInput4 []string = []string{
	"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
}

func TestPart1(t *testing.T) {
	expected := 7

	result := part1(testInput)
	if result != expected {
		t.Errorf("got: %d, want: %d", result, expected)
	}
}

func TestPart1_1(t *testing.T) {
	expected := 5

	result := part1(testInput1)
	if result != expected {
		t.Errorf("got: %d, want: %d", result, expected)
	}
}
func TestPart1_2(t *testing.T) {
	expected := 6

	result := part1(testInput2)
	if result != expected {
		t.Errorf("got: %d, want: %d", result, expected)
	}
}
func TestPart1_3(t *testing.T) {
	expected := 10

	result := part1(testInput3)
	if result != expected {
		t.Errorf("got: %d, want: %d", result, expected)
	}
}
func TestPart1_4(t *testing.T) {
	expected := 11

	result := part1(testInput4)
	if result != expected {
		t.Errorf("got: %d, want: %d", result, expected)
	}
}

func TestPart1Full(t *testing.T) {
	expected := 1833
	input := input.GetStringInput("input.txt")

	result := part1(input)
	if result != expected {
		t.Errorf("got: %d, want: %d", result, expected)
	}
}

// ####################################

func TestPart2(t *testing.T) {
	expected := 19

	result := part2(testInput)
	if result != expected {
		t.Errorf("got: %d, want: %d", result, expected)
	}
}

func TestPart2_1(t *testing.T) {
	expected := 23

	result := part2(testInput1)
	if result != expected {
		t.Errorf("got: %d, want: %d", result, expected)
	}
}
func TestPart2_2(t *testing.T) {
	expected := 23

	result := part2(testInput2)
	if result != expected {
		t.Errorf("got: %d, want: %d", result, expected)
	}
}
func TestPart2_3(t *testing.T) {
	expected := 29

	result := part2(testInput3)
	if result != expected {
		t.Errorf("got: %d, want: %d", result, expected)
	}
}
func TestPart2_4(t *testing.T) {
	expected := 26

	result := part2(testInput4)
	if result != expected {
		t.Errorf("got: %d, want: %d", result, expected)
	}
}

func TestPart2Full(t *testing.T) {
	expected := 3425
	input := input.GetStringInput("input.txt")

	result := part2(input)
	if result != expected {
		t.Errorf("got: %d, want: %d", result, expected)
	}
}
