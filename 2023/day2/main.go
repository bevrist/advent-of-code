// https://adventofcode.com/2023/day/2

package main

import (
	"2023/input"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input := input.FileStringInput("input.txt")

	fmt.Printf("Part1: %d\n", part1(input))
	fmt.Printf("Part2: %d\n", part2(input))
}

// part1 find what games would have been possible with the max number of colored cubes
func part1(input []string) int {
	// game holds count of cubes per game
	type game struct {
		id    int
		red   int
		green int
		blue  int
	}
	var games []game
	for _, line := range input {
		//extract color totals from each game
		gameSplit := strings.Split(line, ": ")
		gameId, _ := strconv.Atoi(strings.Split(gameSplit[0], " ")[1])
		// split games based on "," or ";" separators
		re := regexp.MustCompile("(, |; )")
		gameSteps := re.Split(gameSplit[1], -1)
		currGame := game{id: gameId}
		for _, step := range gameSteps {
			// store the largest number of cubes seen at one time
			switch cube := strings.Split(step, " "); cube[1] {
			case "red":
				val, _ := strconv.Atoi(cube[0])
				currGame.red = int(math.Max(float64(val), float64(currGame.red)))
			case "green":
				val, _ := strconv.Atoi(cube[0])
				currGame.green = int(math.Max(float64(val), float64(currGame.green)))
			case "blue":
				val, _ := strconv.Atoi(cube[0])
				currGame.blue = int(math.Max(float64(val), float64(currGame.blue)))
			}
		}
		games = append(games, currGame)
	}
	// fmt.Printf("%+v\n", games)
	// check if game is possible with these parameters
	maxRed := 12
	maxGreen := 13
	maxBlue := 14
	var total int
	for _, game := range games {
		isGood := true
		if game.red > maxRed || game.blue > maxBlue || game.green > maxGreen {
			isGood = false
		}
		if isGood {
			total += game.id
		}
		// fmt.Printf("%+v, %v\n", game, isGood)
	}
	return total
}

// part2 find what the minimum number of cubes in each game would have been
func part2(input []string) int {
	// game holds count of cubes per game
	type game struct {
		id    int
		red   int
		green int
		blue  int
		power int
	}
	var games []game
	for _, line := range input {
		//extract color totals from each game
		gameSplit := strings.Split(line, ": ")
		gameId, _ := strconv.Atoi(strings.Split(gameSplit[0], " ")[1])
		// split games based on "," or ";" separators
		re := regexp.MustCompile("(, |; )")
		gameSteps := re.Split(gameSplit[1], -1)
		currGame := game{id: gameId}
		for _, step := range gameSteps {
			// store the largest number of cubes seen at one time
			switch cube := strings.Split(step, " "); cube[1] {
			case "red":
				val, _ := strconv.Atoi(cube[0])
				currGame.red = int(math.Max(float64(val), float64(currGame.red)))
			case "green":
				val, _ := strconv.Atoi(cube[0])
				currGame.green = int(math.Max(float64(val), float64(currGame.green)))
			case "blue":
				val, _ := strconv.Atoi(cube[0])
				currGame.blue = int(math.Max(float64(val), float64(currGame.blue)))
			}
		}
		currGame.power = currGame.green * currGame.red * currGame.blue
		games = append(games, currGame)
	}
	// fmt.Printf("%+v\n", games)

	var total int
	for _, game := range games {
		total += game.power
	}
	return total
}
