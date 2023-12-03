// https://adventofcode.com/2021/day/7

package main

import (
	"2021/input"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

func main() {
	input := input.GetStringInput("input.txt")

	fmt.Printf("Part1: %d\n", part1(parseInput(input)))
	fmt.Printf("Part2: %d\n", part2(parseInput(input)))
}

// extract list of integers from input
func parseInput(input []string) []int {
	split := strings.Split(input[0], ",")
	var ret []int
	for _, num := range split {
		inum, err := strconv.Atoi(num)
		i64num := int(inum)
		if err != nil {
			panic(err)
		}
		ret = append(ret, i64num)
	}
	return ret
}

// calculate the position that is the shortest distance for all crabs
func part1(crabs []int) int {
	// fmt.Println("start:", crabs)
	// get min and max position
	var max int = slices.Max(crabs)
	var fuelCost []int = make([]int, max)
	// calculate total fuel cost for each position
	for position := 0; position < max; position++ {
		currFuelCost := 0
		cost := 0
		for _, crabPos := range crabs {
			cost = int(math.Abs(float64(crabPos - position)))
			currFuelCost += cost
			// fmt.Printf("%d;%d:%d \n", crabPos, position, cost)
		}
		// fmt.Println()
		fuelCost[position] = currFuelCost
	}
	// fmt.Println("fuelCost:", fuelCost)
	return slices.Min(fuelCost)
}

// ===================================================

func part2(crabs []int) int {
	// fmt.Println("start:", crabs)
	// get min and max position
	var max int = slices.Max(crabs)
	var fuelCost []int = make([]int, max)
	// calculate total fuel cost for each position
	for position := 0; position < max; position++ {
		currFuelCost := 0
		cost := 0
		for _, crabPos := range crabs {
			distance := int(math.Abs(float64(crabPos - position)))
			cost = (int(math.Pow(float64(distance), float64(2))) + distance) / 2
			currFuelCost += cost
			// fmt.Printf("%d;%d:%d \n", crabPos, position, cost)
		}
		// fmt.Println()
		fuelCost[position] = currFuelCost
	}
	// fmt.Println("fuelCost:", fuelCost)
	return slices.Min(fuelCost)
}
