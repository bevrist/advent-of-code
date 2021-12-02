package main

import "testing"

var entries []string = []string{"1-3 a: abcde", "1-3 b: cdefg", "2-9 c: ccccccccc"}

func TestPart1(t *testing.T) {
	expected := 2

	result := part1(entries)
	if result != expected {
		t.Errorf("got: %d, want: %d", result, expected)
	}
}

func TestPart2(t *testing.T) {
	expected := 1

	result := part2(entries)
	if result != expected {
		t.Errorf("got: %d, want: %d", result, expected)
	}
}
