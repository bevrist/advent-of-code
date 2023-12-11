// https://adventofcode.com/2023/day/8

package main

import (
	"2023/input"
	"fmt"
	"strings"
)

func main() {
	input := input.FileStringInput("input.txt")
	// input := []string{
	// 	"RL",
	// 	"",
	// 	"AAA = (BBB, CCC)",
	// 	"BBB = (DDD, EEE)",
	// 	"CCC = (ZZZ, GGG)",
	// 	"DDD = (DDD, DDD)",
	// 	"EEE = (EEE, EEE)",
	// 	"GGG = (GGG, GGG)",
	// 	"ZZZ = (ZZZ, ZZZ)",
	// }
	// input := []string{
	// 	"LLR",
	// 	"",
	// 	"AAA = (BBB, BBB)",
	// 	"BBB = (AAA, ZZZ)",
	// 	"ZZZ = (ZZZ, ZZZ)",
	// }

	// input for part 2
	// input := []string{
	// 	"LR",
	// 	"",
	// 	"11A = (11B, XXX)",
	// 	"11B = (XXX, 11Z)",
	// 	"11Z = (11B, XXX)",
	// 	"22A = (22B, XXX)",
	// 	"22B = (22C, 22C)",
	// 	"22C = (22Z, 22Z)",
	// 	"22Z = (22B, 22B)",
	// 	"XXX = (XXX, XXX)",
	// }

	parsed := parseInput(input)

	fmt.Printf("Part1: %d\n", part1(parsed))
	// fmt.Printf("Part2: %d\n", part2(parsed))
}

type instructions struct {
	directions string
	currentDir int
	paths      map[string]path
}

// returns next direction in instructions.directions list
func (i *instructions) next() string {
	ret := i.directions[i.currentDir]
	i.currentDir++
	// loop back to beginning
	if i.currentDir >= len(i.directions) {
		i.currentDir = 0
	}
	return string(ret)
}

type path struct {
	start string
	left  string
	right string
}

func parseInput(input []string) instructions {
	var newInstruction instructions
	newInstruction.paths = make(map[string]path)
	for i, line := range input {
		if i == 0 {
			newInstruction.directions = line
			continue
		}
		if i == 1 {
			continue
		}
		split := strings.Split(line, " ")
		start := split[0]
		leftPath := split[2][1 : len(split[2])-1]
		rightPath := split[3][:len(split[3])-1]
		newInstruction.paths[start] = path{start, leftPath, rightPath}
		// fmt.Println(start, leftPath, rightPath)
	}
	return newInstruction
}

func part1(instructions instructions) int {
	// fmt.Println(instructions)
	currPath := "AAA"
	steps := 1
	for {
		nextDir := instructions.next()
		if nextDir == "L" {
			currPath = instructions.paths[currPath].left
		} else { // "R"
			currPath = instructions.paths[currPath].right
		}
		if currPath == "ZZZ" {
			break
		}
		steps++
	}
	return steps
}

// endsWith returns true if string ends with same letter
func endsWith(input string, letter string) bool {
	if string(input[len(input)-1]) == letter {
		return true
	} else {
		return false
	}
}

// func part2(instructions instructions) int {
// 	// fmt.Println(instructions)
// 	var paths []string
// 	// get starting paths that end with A
// 	for path := range instructions.paths {
// 		if endsWith(path, "A") {
// 			paths = append(paths, path)
// 		}
// 	}
// 	// fmt.Println(paths)

// 	// walk all paths
// 	steps := 0
// 	for {
// 		if (steps % 10000000) == 0 {
// 			fmt.Print(steps, ",")
// 		}
// 		nextDir := instructions.next()
// 		isDone := true
// 		for i := range paths {
// 			if nextDir == "L" {
// 				paths[i] = instructions.paths[paths[i]].left
// 			} else { // "R"
// 				paths[i] = instructions.paths[paths[i]].right
// 			}
// 			if !endsWith(paths[i], "Z") {
// 				isDone = false
// 			}
// 		}
// 		// fmt.Println(paths)
// 		steps++
// 		if isDone {
// 			return steps
// 		}
// 	}
// }
