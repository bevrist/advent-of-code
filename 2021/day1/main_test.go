package main

import "testing"

var entries []int = []int{
	199,
	200,
	208,
	210,
	200,
	207,
	240,
	269,
	260,
	263,
}

func TestPart1(t *testing.T) {
	expected := 7

	result := part1(entries)
	if result != expected {
		t.Errorf("got: %d, want: %d", result, expected)
	}
}

func TestPart2(t *testing.T) {
	expected := 5

	result := part2(entries)
	if result != expected {
		t.Errorf("got: %d, want: %d", result, expected)
	}
}
