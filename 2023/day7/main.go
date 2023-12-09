// https://adventofcode.com/2023/day/7

package main

import (
	"2023/input"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input := input.FileStringInput("input.txt")
	// input := []string{
	// 	"32T3K 765",
	// 	"T55J5 684",
	// 	"KK677 28",
	// 	"KTJJT 220",
	// 	"QQQJA 483",
	// }

	parsed := parseInput(input)

	fmt.Printf("Part1: %d\n", part1(parsed))
	fmt.Printf("Part2: %d\n", part2(parsed))
}

type hand struct {
	cards string
	bet   int
}

func parseInput(input []string) []hand {
	var hands []hand
	for _, line := range input {
		split := strings.Split(line, " ")
		num, _ := strconv.Atoi(split[1])
		hands = append(hands, hand{cards: split[0], bet: num})
	}
	return hands
}

func cardRank(cards string) int {
	//count up number of each type of card
	var totals map[string]int = make(map[string]int)
	for i := range cards {
		totals[string(cards[i])] += 1
	}
	// fmt.Println(totals)

	// determine type of hand the cards represent
	// Five of a kind => 6
	if len(totals) == 1 {
		return 6
	}
	if len(totals) == 2 {
		// Four of a kind => 5
		for _, val := range totals {
			if val == 4 {
				return 5
			}
		}
		// Full house => 4
		var hasThree, hasTwo bool
		for _, val := range totals {
			if val == 3 {
				hasThree = true
			}
			if val == 2 {
				hasTwo = true
			}
		}
		if hasThree && hasTwo {
			return 4
		}
	}
	if len(totals) == 3 {
		for _, val := range totals {
			// Three of a kind => 3
			if val == 3 {
				return 3
			}
		}
		// Two pair => 2
		return 2
	}
	// One pair => 1
	if len(totals) == 4 {
		return 1
	}
	// Trash => 0
	return 0
}

// handValues returns sortable string representation of cards
func handValues(cards string) string {
	// convert card values to alphabetical representation for sorting
	strengthMap := map[string]string{
		"A": "A",
		"K": "B",
		"Q": "C",
		"J": "D",
		"T": "E",
		"9": "F",
		"8": "G",
		"7": "H",
		"6": "I",
		"5": "J",
		"4": "K",
		"3": "L",
		"2": "M",
	}
	var ret string
	for i := range cards {
		ret += strengthMap[string(cards[i])]
	}
	return ret
}

func part1(hands []hand) int {
	// sort hands by rank
	// fmt.Println(hands)
	sort.Slice(hands, func(i, j int) bool {
		iRank := cardRank(hands[i].cards)
		jRank := cardRank(hands[j].cards)
		// if rank is the same, sort by hand values
		if iRank == jRank {
			return handValues(hands[i].cards) > handValues(hands[j].cards)
		} else {
			return iRank < jRank
		}
	})
	// fmt.Println(hands)
	// calculate total winnings
	var total int
	for i, hand := range hands {
		// fmt.Println(i+1, hand)
		total += (i + 1) * hand.bet
	}
	return total
}

// ================================================================================

func cardRankJoker(cards string) int {
	//sub joker for most common card then get cardRank
	var totals map[string]int = make(map[string]int)
	var jokers int
	for i := range cards {
		if string(cards[i]) == "J" {
			jokers++
			continue
		}
		totals[string(cards[i])] += 1
	}

	// handle all joker edge case
	if jokers == 5 {
		return cardRank("AAAAA")
	}

	most := 0
	mostCard := ""
	for card, count := range totals {
		if count > most {
			most = count
			mostCard = card
		}
	}
	return cardRank(strings.ReplaceAll(cards, "J", mostCard))
}

// handValues returns sortable string representation of cards, where joker is the weakest
func handValuesJoker(cards string) string {
	// convert card values to alphabetical representation for sorting
	strengthMap := map[string]string{
		"A": "A",
		"K": "B",
		"Q": "C",
		"J": "Z", // joker is now the weakest card
		"T": "E",
		"9": "F",
		"8": "G",
		"7": "H",
		"6": "I",
		"5": "J",
		"4": "K",
		"3": "L",
		"2": "M",
	}
	var ret string
	for i := range cards {
		ret += strengthMap[string(cards[i])]
	}
	return ret
}

func part2(hands []hand) int {
	// sort hands by rank
	// fmt.Println(hands)
	sort.Slice(hands, func(i, j int) bool {
		iRank := cardRankJoker(hands[i].cards)
		jRank := cardRankJoker(hands[j].cards)
		// if rank is the same, sort by hand values
		if iRank == jRank {
			return handValuesJoker(hands[i].cards) > handValuesJoker(hands[j].cards)
		} else {
			return iRank < jRank
		}
	})
	// fmt.Println(hands)
	// calculate total winnings
	var total int
	for i, hand := range hands {
		// fmt.Println(i+1, hand)
		total += (i + 1) * hand.bet
	}
	return total
}
