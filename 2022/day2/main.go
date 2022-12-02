// https://adventofcode.com/2022/day/1

package main

import (
	"2022/input"
	"fmt"
)

func main() {
	input := input.GetStringInput("input.txt")

	parsed := parseInput(input)

	fmt.Printf("Part1: %d\n", part1(parsed))
	// fmt.Printf("Part2: %d\n", part2(parsed))
}

// stats for each round
type Round struct {
	opponent int
	player   int
	outcome  int
	score    int
}

const (
	rock     int = 1
	paper    int = 2
	scissors int = 3

	win  int = 6
	tie  int = 3
	loss int = 0
)

func parseInput(input []string) []Round {
	var ret []Round
	for _, entry := range input {
		currRound := new(Round)
		// get opponent play
		switch string(entry[0]) {
		case "A":
			currRound.opponent = rock
		case "B":
			currRound.opponent = paper
		case "C":
			currRound.opponent = scissors
		}
		// get player play
		switch string(entry[2]) {
		case "X":
			currRound.player = rock
		case "Y":
			currRound.player = paper
		case "Z":
			currRound.player = scissors
		}
		// get game outcome
		switch currRound.opponent {
		case rock:
			switch currRound.player {
			case rock:
				currRound.outcome = tie
			case paper:
				currRound.outcome = win
			case scissors:
				currRound.outcome = loss
			}
		case paper:
			switch currRound.player {
			case rock:
				currRound.outcome = loss
			case paper:
				currRound.outcome = tie
			case scissors:
				currRound.outcome = win
			}
		case scissors:
			switch currRound.player {
			case rock:
				currRound.outcome = win
			case paper:
				currRound.outcome = loss
			case scissors:
				currRound.outcome = tie
			}
		}
		// add up score
		currRound.score = currRound.player + currRound.outcome

		// fmt.Printf("%v\n", entry)
		// fmt.Printf("%+v\n\n", currRound)
		ret = append(ret, *currRound)
	}
	return ret
}

// get total score for the provided play
func part1(rounds []Round) int {
	total := 0
	for _, round := range rounds {
		total += round.score
	}
	return total
}

// // get total calories of top 3 elves
// func part2(invs []Inventory) int {
// 	// sort inventory list from greatest to least
// 	sort.Slice(invs, func(i, j int) bool {
// 		return invs[i].total > invs[j].total
// 	})
// 	// return total of top 3 inventories
// 	return invs[0].total + invs[1].total + invs[2].total
// }
