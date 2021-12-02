// https://adventofcode.com/2021/day/1

package main

import (
	"2021/input"
	"fmt"
)

func main() {
	input := input.GetIntInput("input.txt")

	fmt.Printf("Part1: %d\n", part1(input))
	fmt.Printf("Part2: %d\n", part2(input))
}

// count number of times input increases
func part1(input []int) int {
	last := input[0]
	increases := 0
	for i, depth := range input {
		if i == 0 {
			continue
		}
		if depth > last {
			increases++
		}
		last = depth
	}
	return increases
}

// count number of increases in a 3 value sliding window
func part2(input []int) int {
	back3, back2, back1 := input[0], input[1], input[2]
	increases := 0
	for i, depth := range input {
		if i <= 2 {
			continue
		}
		//compare sliding windows
		if (back3 + back2 + back1) < (depth + back2 + back1) {
			increases++
		}
		back3, back2, back1 = back2, back1, depth
	}
	return increases
}
