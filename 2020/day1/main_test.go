package main

import "testing"

var entries []int = []int{1721, 979, 366, 299, 675, 1456}

func TestPart1(t *testing.T) {
	expected := 514579

	result := part1(entries)
	if result != expected {
		t.Errorf("got: %d, want: %d", result, expected)
	}
}

func TestPart2(t *testing.T) {
	expected := 241861950

	result := part2(entries)
	if result != expected {
		t.Errorf("got: %d, want: %d", result, expected)
	}
}
