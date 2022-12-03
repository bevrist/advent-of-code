// https://adventofcode.com/2022/day/2

package main

import (
	"2022/input"
	"fmt"
)

func main() {
	input := input.GetStringInput("input.txt")

	fmt.Printf("Part1: %d\n", part1(input))
	fmt.Printf("Part2: %d\n", part2(input))
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

// get total score for the provided play
func part1(input []string) int {
	var rounds []Round
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
		rounds = append(rounds, *currRound)
	}

	total := 0
	for _, round := range rounds {
		total += round.score
	}
	return total
}

// get total score for the optimal play
func part2(input []string) int {
	var rounds []Round
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
		// get desired outcome
		switch string(entry[2]) {
		case "X":
			currRound.outcome = loss
		case "Y":
			currRound.outcome = tie
		case "Z":
			currRound.outcome = win
		}
		// get player's play based on desired outcome
		switch currRound.opponent {
		case rock:
			switch currRound.outcome {
			case win:
				currRound.player = paper
			case loss:
				currRound.player = scissors
			case tie:
				currRound.player = rock
			}
		case paper:
			switch currRound.outcome {
			case win:
				currRound.player = scissors
			case loss:
				currRound.player = rock
			case tie:
				currRound.player = paper
			}
		case scissors:
			switch currRound.outcome {
			case win:
				currRound.player = rock
			case loss:
				currRound.player = paper
			case tie:
				currRound.player = scissors
			}
		}
		// add up score
		currRound.score = currRound.player + currRound.outcome

		// fmt.Printf("%v\n", entry)
		// fmt.Printf("%+v\n\n", currRound)
		rounds = append(rounds, *currRound)
	}

	total := 0
	for _, round := range rounds {
		total += round.score
	}
	return total
}
