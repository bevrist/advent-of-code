// https://adventofcode.com/2023/day/4

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

	// input := []string{
	// 	"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
	// 	"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
	// 	"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
	// 	"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
	// 	"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
	// 	"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
	// }

	parsed := parseInput(input)

	fmt.Printf("Part1: %d\n", part1(parsed))
	fmt.Printf("Part2: %d\n", part2(parsed))
}

// card represents a scratch card
type card struct {
	id          int
	winningNums []int
	yourNums    []int
	copies      int
}

// parseInput extracts card structures from the input
func parseInput(input []string) []card {
	var cards []card
	re := regexp.MustCompile(`\s+`) //empty space regex
	for _, line := range input {
		colSplit := strings.Split(line, ": ")
		idStr := strings.Split(colSplit[0], " ")[1]
		allNums := strings.Split(colSplit[1], " | ")
		winningNumsStr := re.Split(allNums[0], -1)
		yourNumsStr := re.Split(allNums[1], -1)
		// convert string numbers to integers
		id, _ := strconv.Atoi(idStr)
		var winningNums, yourNums []int
		for _, winNum := range winningNumsStr {
			if winNum == "" { // skip empty values
				continue
			}
			num, _ := strconv.Atoi(winNum)
			winningNums = append(winningNums, num)
		}
		for _, yourNum := range yourNumsStr {
			if yourNum == "" { // skip empty values
				continue
			}
			num, _ := strconv.Atoi(yourNum)
			yourNums = append(yourNums, num)
		}

		cards = append(cards, card{id, winningNums, yourNums, 0})
	}
	return cards
}

// part1 calculate the point total of each card and add them up
func part1(cards []card) int {
	type cardMatch struct {
		card    card
		matches []int
	}
	var cardMatches []cardMatch
	// number of winning matches per card
	for _, card := range cards {
		// fmt.Println(card)
		newMatch := cardMatch{card: card}
		for _, win := range card.winningNums {
			for _, yours := range card.yourNums {
				if yours == win {
					newMatch.matches = append(newMatch.matches, win)
					break
				}
			}
		}
		cardMatches = append(cardMatches, newMatch)
		// fmt.Println(newMatch)
	}

	var total int
	// calculate card totals
	for _, cardMatch := range cardMatches {
		cardValue := int(math.Pow(float64(2), float64(len(cardMatch.matches)-1)))
		total += cardValue
	}
	return total
}

// part2 calculate how many winning copies of scratchcards you end up with
func part2(cards []card) int {
	type cardMatch struct {
		card    card
		matches []int
	}
	var cardMatches []cardMatch
	// number of winning matches per card
	for _, card := range cards {
		// fmt.Println(card)
		newMatch := cardMatch{card: card}
		for _, win := range card.winningNums {
			for _, yours := range card.yourNums {
				if yours == win {
					newMatch.matches = append(newMatch.matches, win)
					break
				}
			}
		}
		cardMatches = append(cardMatches, newMatch)
		// fmt.Println(newMatch)
	}

	var total int
	// iterate through cards + copies and add up total card count
	for c, cardMatch := range cardMatches {
		total += cardMatch.card.copies + 1            // add up current card copy total + original(1)
		for i := 0; i <= cardMatch.card.copies; i++ { // repeat for count of copies of card (at least once for og card)
			for j := c; j <= c+len(cardMatch.matches); j++ { // increment copy counter for following cards
				cardMatches[j].card.copies++
			}
		}
	}
	return total
}
