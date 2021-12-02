// https://adventofcode.com/2021/day/2

package main

import (
	"2021/input"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := input.GetStringInput("input.txt")

	fmt.Printf("Part1: %d\n", part1(input))
	fmt.Printf("Part2: %d\n", part2(input))
}

// find ending position after following input directions
func part1(input []string) int {
	//x and y position
	var pos []int = []int{0, 0}
	for _, dir := range input {
		instruction := strings.Fields(dir)
		distance, err := strconv.Atoi(instruction[1])
		if err != nil {
			panic(err)
		}
		if instruction[0] == "forward" {
			pos[0] += distance
		} else if instruction[0] == "down" {
			pos[1] += distance
		} else if instruction[0] == "up" {
			pos[1] -= distance
		}
	}

	return pos[0] * pos[1]
}

func part2(input []string) int {
	//x and y position, and aim
	var pos []int = []int{0, 0, 0}
	for _, dir := range input {
		instruction := strings.Fields(dir)
		distance, err := strconv.Atoi(instruction[1])
		aim := pos[2]
		if err != nil {
			panic(err)
		}
		if instruction[0] == "forward" {
			pos[0] += distance
			pos[1] += distance * aim
		} else if instruction[0] == "down" {
			pos[2] += distance
		} else if instruction[0] == "up" {
			pos[2] -= distance
		}
	}

	return pos[0] * pos[1]
}
