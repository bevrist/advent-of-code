package main

import "testing"

var entries []string = []string{
	"..##.......",
	"#...#...#..",
	".#....#..#.",
	"..#.#...#.#",
	".#...##..#.",
	"..#.##.....",
	".#.#.#....#",
	".#........#",
	"#.##...#...",
	"#...##....#",
	".#..#...#.#",
}

func TestPart1(t *testing.T) {
	expected := 7

	result := part1(entries)
	if result != expected {
		t.Errorf("got: %d, want: %d", result, expected)
	}
}

func TestPart2(t *testing.T) {
	expected := 336

	result := part2(entries)
	if result != expected {
		t.Errorf("got: %d, want: %d", result, expected)
	}
}
