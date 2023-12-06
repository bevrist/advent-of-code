// https://adventofcode.com/2023/day/6

package main

import (
	"2023/input"
	"fmt"
	"log"
	"regexp"
	"strconv"
)

func main() {
	input := input.FileStringInput("input.txt")

	// input := []string{
	// 	"Time:      7  15   30",
	// 	"Distance:  9  40  200",
	// }

	parsed := parseInput(input)

	fmt.Printf("Part1: %d\n", part1(parsed))
	fmt.Printf("Part2: %d\n", part2(input))
}

type result struct {
	time     int
	distance int
}

// strToInt convert string representation of number to integer, panic on failure
func strToInt(in string) int {
	num, err := strconv.Atoi(in)
	if err != nil {
		log.Fatalln(err)
	}
	return num
}

func parseInput(input []string) []result {
	var results []result
	re := regexp.MustCompile(`\s+`)
	timeSplit := re.Split(input[0], -1)
	distanceSplit := re.Split(input[1], -1)
	// fmt.Println(timeSplit, distanceSplit)
	for i := 1; i < len(timeSplit); i++ {
		results = append(results, result{strToInt(timeSplit[i]), strToInt(distanceSplit[i])})
	}
	return results
}

// findWinningStrats calculates all possible hold times to beat distance record
func findWinningStrats(result result) []int {
	var winners []int
	for hold := 0; hold < result.time; hold++ {
		speed := hold
		distance := speed * (result.time - hold)
		if distance > result.distance {
			winners = append(winners, hold)
		}
	}
	return winners
}

func part1(results []result) int {
	// fmt.Println(results)
	var total int = 1
	for _, res := range results {
		winningStrategies := findWinningStrats(res)
		// fmt.Println(s, len(s))
		total *= len(winningStrategies)
	}
	return total
}

func part2(input []string) int {
	var res result = result{}
	re := regexp.MustCompile(`\s+`)
	timeSplit := re.Split(input[0], -1)
	distanceSplit := re.Split(input[1], -1)
	// fmt.Println(timeSplit, distanceSplit)
	// read input across erroneous spaces
	for i := 1; i < len(timeSplit); i++ {
		res = result{
			time:     strToInt(fmt.Sprint(res.time) + timeSplit[i]),
			distance: strToInt(fmt.Sprint(res.distance) + distanceSplit[i]),
		}
	}
	// fmt.Println(res)

	return len(findWinningStrats(res))
}
