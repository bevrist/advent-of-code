package main

import "testing"

var entries []string = []string{
	"forward 5",
	"down 5",
	"forward 8",
	"up 3",
	"down 8",
	"forward 2",
}

func TestPart1(t *testing.T) {
	expected := 7

	result := part1(entries)
	if result != expected {
		t.Errorf("got: %d, want: %d", result, expected)
	}
}

// func TestPart2(t *testing.T) {
// 	expected := 5

// 	result := part2(entries)
// 	if result != expected {
// 		t.Errorf("got: %d, want: %d", result, expected)
// 	}
// }
