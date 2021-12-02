// https://adventofcode.com/2020/day/2

package main

import (
	"2020/input"
	"fmt"
	"strings"
)

func main() {
	input := input.GetStringInput("input.txt")

	fmt.Printf("Part1: %d\n", part1(input))
	fmt.Printf("Part2: %d\n", part2(input))
}

type policy struct {
	min, max int
	char     string
	password string
}

// parsePolicy parses the formatted input into a slice of policies
func parsePolicy(input []string) []policy {
	var policies []policy
	for _, line := range input {
		var p policy
		fmt.Sscanf(line, "%d-%d %s %s", &p.min, &p.max, &p.char, &p.password)
		p.char = strings.TrimSuffix(p.char, ":") //Remove trailing colon from scan
		policies = append(policies, p)
	}
	return policies
}

// part1 returns number of passwords that meet policy (must contain [min>=char<=max] chars in pasword)
func part1(input []string) int {
	policies := parsePolicy(input)
	passingCount := 0
	for _, p := range policies {
		count := strings.Count(p.password, p.char)
		if count >= p.min && count <= p.max {
			passingCount++
		}
	}
	return passingCount
}

// part2 returns number of passwords that meet policy (must contain exclusive char in min or max position)
func part2(input []string) int {
	policies := parsePolicy(input)
	passingCount := 0
	for _, p := range policies {
		//xor that the character only appears in one of the specified locations
		if (p.password[p.min-1] == p.char[0]) != (p.password[p.max-1] == p.char[0]) {
			passingCount++
		}
	}
	return passingCount
}
