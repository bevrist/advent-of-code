// https://adventofcode.com/2022/day/4

package main

import (
	"2022/input"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := input.GetStringInput("input.txt")

	fmt.Printf("Part1: %d\n", part1(input))
	fmt.Printf("Part2: %d\n", part2(input))
}

type Pair struct {
	task1 []int
	task2 []int
}

// create Pair list with arrays of numbers contained in pairs
func parseInput(in []string) []Pair {
	var ret []Pair
	for _, line := range in {
		currPair := Pair{}
		pair := strings.Split(line, ",")
		left := strings.Split(pair[0], "-")
		right := strings.Split(pair[1], "-")
		left0, _ := strconv.Atoi(left[0])
		left1, _ := strconv.Atoi(left[1])
		right0, _ := strconv.Atoi(right[0])
		right1, _ := strconv.Atoi(right[1])
		// add task numbers to Pair struct
		for i := left0; i <= left1; i++ {
			currPair.task1 = append(currPair.task1, i)
		}
		for i := right0; i <= right1; i++ {
			currPair.task2 = append(currPair.task2, i)
		}
		ret = append(ret, currPair)
	}
	return ret
}

// get count of pairs that fully overlap
func part1(in []string) int {
	parsed := parseInput(in)
	// find pairs that fully overlap
	count := 0
	for _, pair := range parsed {
		// for each pair, compare all of one side with the other and save if all matched
		allLeftMatch := true
		for _, i := range pair.task1 {
			matchExists := false
			for _, j := range pair.task2 {
				if i == j {
					matchExists = true
					break
				}
			}
			if !matchExists {
				allLeftMatch = false
			}
		}
		allRightMatch := true
		for _, i := range pair.task2 {
			matchExists := false
			for _, j := range pair.task1 {
				if i == j {
					matchExists = true
					break
				}
			}
			if !matchExists {
				allRightMatch = false
			}
		}
		// if either first or second pairs' entire list overlaps
		if allLeftMatch || allRightMatch {
			count++
		}
	}
	return count
}

// count number of pairs that overlap at all
func part2(in []string) int {
	parsed := parseInput(in)
	count := 0
	for _, pair := range parsed {
		overlaped := false
		// if any numbers in pair1 overlap with pair2, increment count
		for _, p1 := range pair.task1 {
			for _, p2 := range pair.task2 {
				if p1 == p2 {
					overlaped = true
					count++
					break
				}
			}
			if overlaped {
				break
			}
		}
	}
	return count
}
