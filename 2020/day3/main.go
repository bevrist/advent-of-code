// https://adventofcode.com/2020/day/3

package main

import (
	"2020/input"
	"fmt"
)

func main() {
	input := input.GetStringInput("input.txt")

	fmt.Printf("Part1: %d\n", part1(input))
	fmt.Printf("Part2: %d\n", part2(input))
}

// mapArray creates a 2d array representing the "map"
func mapArray(input []string) [][]int {
	var array [][]int
	for _, line := range input {
		var row []int
		for _, char := range line {
			if char == '#' {
				row = append(row, 1)
			} else {
				row = append(row, 0)
			}
		}
		array = append(array, row)
	}
	return array
}

// findTrees returns the number of trees that are hit by the slope
func findTrees(input []string, slope []int) int {
	mapArr := mapArray(input)
	loc := []int{0, 0}
	treeCount := 0

	for loc[1] < len(mapArr) && loc[0] < len(mapArr[0]) {
		if mapArr[loc[1]][loc[0]] == 1 { //if the current location is a tree, increment tree counter
			treeCount++
			// mapArr[loc[1]][loc[0]] = 2
		} else {
			// mapArr[loc[1]][loc[0]] = 8
		}
		//move slope
		loc[0] += slope[0]
		loc[1] += slope[1]
		//if off right side of map, loop back to start
		if loc[0] >= len(mapArr[0]) {
			loc[0] -= len(mapArr[0])
		}
	}

	// for row, _ := range mapArr {
	// 	fmt.Println(row)
	// }

	return treeCount
}

// part1 finds the number of trees hit with a slope of 3, 1
func part1(input []string) int {
	slope := []int{3, 1}
	return findTrees(input, slope)
}

// part2 finds the product of 5 different slopes
func part2(input []string) int {
	slope1 := []int{1, 1}
	slope2 := []int{3, 1}
	slope3 := []int{5, 1}
	slope4 := []int{7, 1}
	slope5 := []int{1, 2}
	treeCount1 := findTrees(input, slope1)
	treeCount2 := findTrees(input, slope2)
	treeCount3 := findTrees(input, slope3)
	treeCount4 := findTrees(input, slope4)
	treeCount5 := findTrees(input, slope5)
	return treeCount1 * treeCount2 * treeCount3 * treeCount4 * treeCount5
}
