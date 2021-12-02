// https://adventofcode.com/2020/day/7

package main

import (
	"2020/input"
	"fmt"
)

func main() {
	input := input.GetStringInput("input.txt")

	fmt.Printf("Part1: %d\n", part1(input))
	// fmt.Printf("Part2: %d\n", part2(input))
}

type bag struct {
	contains map[string]int
	name string
}

parseInput (input []string) map[string]bag {
	//parse input into bag objects


}

func part1(input []string) int {

}
