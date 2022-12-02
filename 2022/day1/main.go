// https://adventofcode.com/2022/day/1

package main

import (
	"2022/input"
	"fmt"
	"sort"
	"strconv"
)

func main() {
	input := input.GetStringInput("input.txt")

	parsed := parseInput(input)

	fmt.Printf("Part1: %d\n", part1(parsed))
	fmt.Printf("Part2: %d\n", part2(parsed))
}

// Inventory represents an elf inventory
type Inventory struct {
	calories []int
	total    int
}

func parseInput(input []string) []Inventory {
	var invs []Inventory
	var currInv *Inventory = new(Inventory)
	for _, line := range input {
		if line == "" {
			// find total calorie count per inventory
			for _, cal := range currInv.calories {
				currInv.total += cal
			}
			invs = append(invs, *currInv)
			currInv = new(Inventory)
			continue
		}
		num, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		currInv.calories = append(currInv.calories, num)
	}
	// find total calorie count per inventory
	for _, cal := range currInv.calories {
		currInv.total += cal
	}
	invs = append(invs, *currInv)
	return invs
}

// find elf with the most calories in inventory
func part1(invs []Inventory) int {
	top := 0
	for _, inv := range invs {
		if inv.total > top {
			top = inv.total
		}
	}
	return top
}

// get total calories of top 3 elves
func part2(invs []Inventory) int {
	// sort inventory list from greatest to least
	sort.Slice(invs, func(i, j int) bool {
		return invs[i].total > invs[j].total
	})
	// return total of top 3 inventories
	return invs[0].total + invs[1].total + invs[2].total
}
