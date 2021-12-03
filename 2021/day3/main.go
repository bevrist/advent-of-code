// https://adventofcode.com/2021/day/3

package main

import (
	"2021/input"
	"fmt"
	"strconv"
)

func main() {
	input := input.GetStringInput("input.txt")

	fmt.Printf("Part1: %d\n", part1(input))
	// fmt.Printf("Part2: %d\n", part2(input))
}

// parse input into list of rune arrays
func parseInput(input []string) [][]int {
	var parsed [][]int

	for _, line := range input {
		lineRunes := []rune(line)

		var lineInts []int = make([]int, len(line))
		for i, char := range lineRunes {
			charInt, _ := strconv.Atoi(string(char))
			lineInts[i] = charInt
		}
		parsed = append(parsed, lineInts)
	}

	return parsed
}

// count the number of bits in each position
func countBits(input [][]int) []int {
	var counts []int = make([]int, len(input[0]))

	for _, item := range input {
		for i, num := range item {
			counts[i] += num
		}
	}

	return counts
}

// calculate most common bits in each position
func part1(input []string) int {
	gamma := make([]int, len(input[0]))
	epsilon := make([]int, len(input[0]))
	length := len(input)
	// calculate the gamma and epsilon values
	for i, bit := range countBits(parseInput(input)) {
		if bit >= (length / 2) {
			gamma[i] = 1
		} else {
			epsilon[i] = 1
		}
	}

	//convert from binary to decimal
	retGamma := ""
	for _, i := range gamma {
		retGamma += fmt.Sprintf("%d", i)
	}
	retEpsilon := ""
	for _, i := range epsilon {
		retEpsilon += fmt.Sprintf("%d", i)
	}
	out, _ := strconv.ParseInt(retGamma, 2, 64)
	out2, _ := strconv.ParseInt(retEpsilon, 2, 64)

	return int(out * out2)
}

// func part2(input []string) int {
// 	gamma := make([]int, len(input[0]))
// 	epsilon := make([]int, len(input[0]))
// 	length := len(input)
// 	parsed := parseInput(input)
// 	// calculate the gamma and epsilon values
// 	for i, bit := range countBits(parsed) {
// 		if bit >= (length / 2) {
// 			gamma[i] = 1
// 		} else {
// 			epsilon[i] = 1
// 		}
// 	}

// 	// find gamma and epsilon numbers
// 	gammaList := []int{}
// 	epsilonList := []int{}
// 	for _, line := range parsed {
// 		isGamma := true
// 		isEpsilon := true
// 		for i, bit := range line {
// 			if gamma[i] != bit {
// 				isGamma = false
// 			}
// 			if epsilon[i] != bit {
// 				isEpsilon = false
// 			}
// 		}
// 		if isGamma {
// 			gammaList = line
// 		}
// 		if isEpsilon {
// 			epsilonList = line
// 		}
// 	}

// 	fmt.Println(gammaList)
// 	fmt.Println(epsilonList)

// 	return 1
// }
