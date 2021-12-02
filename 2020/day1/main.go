// https://adventofcode.com/2020/day/1

package main

import (
	"2020/input"
	"fmt"
)

func main() {
	input := input.GetIntInput("input.txt")

	fmt.Printf("Part1: %d\n", part1(input))
	fmt.Printf("Part2: %d\n", part2(input))
}

// returns the product of 2 entries that sum to 2020
func part1(input []int) int {
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input); j++ {
			if input[i]+input[j] == 2020 {
				return input[i] * input[j]
			}
		}
	}
	return 0
}

// returns the product of 3 entries that sum to 2020
func part2(input []int) int {
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input); j++ {
			for k := 0; k < len(input); k++ {
				if input[i]+input[j]+input[k] == 2020 {
					return input[i] * input[j] * input[k]
				}
			}
		}
	}
	return 0
}
