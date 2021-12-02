// https://adventofcode.com/2020/day/7

package main

import (
	"2020/input"
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	input := input.GetStringInput("input.txt")

	fmt.Printf("Part1: %d\n", part1(input))
	fmt.Printf("Part2: %d\n", part2(input))
}

type bag struct {
	name  string
	count int
}

// get map of bags and what they contain from input
func parseInput(input []string) map[string][]bag {
	//parse input into bag objects
	re := regexp.MustCompile(`^(\w+ \w+) bags contain (\d|no) (\w+ \w+)`)
	sre := regexp.MustCompile(`(\d) (\w+ \w+) bag.?`)

	var bagMap = make(map[string][]bag)
	for _, line := range input {
		topBagName := re.FindStringSubmatch(line)[1]
		subBags := sre.FindAllStringSubmatch(line, -1)

		subBagList := make([]bag, len(subBags))
		//get list of bags that are inside this bag
		for i, subBag := range subBags {
			if subBag[1] == "no" {
				continue
			}
			count, _ := strconv.Atoi(subBag[1])
			newBag := bag{count: count, name: subBag[2]}
			subBagList[i] = newBag
		}
		bagMap[topBagName] = subBagList
	}
	return bagMap
}

// recursive function to find if current bag contains the target bag
func containsBag(bagMap map[string][]bag, currentBag, targetBag string) bool {
	if currentBag == targetBag {
		return true
	}
	for _, subBag := range bagMap[currentBag] {
		//return true if any sub bag contains the target bag
		if containsBag(bagMap, subBag.name, targetBag) {
			return true
		}
	}
	return false
}

// recursive function to count how many bags a target bag contains
func countBags(bagMap map[string][]bag, currentBag string) int {
	if len(bagMap[currentBag]) == 0 {
		return 1 // if the bag contains nothing, count only self
	}
	count := 0
	// for each sub bag, count how many bags it contains
	for _, subBag := range bagMap[currentBag] {
		count += (countBags(bagMap, subBag.name) * subBag.count)
	}
	return count + 1 //add 1 for the bag itself
}

// find number of bags that could contain the target bag
func part1(input []string) int {
	targetBag := "shiny gold"
	bagMap := parseInput(input)

	// for each bag in map, check if it contains the target bag
	bagsThatHoldTarget := map[string]bool{}
	for bag, contains := range bagMap {
		if bag == targetBag {
			continue
		}
		for _, subBag := range contains {
			if containsBag(bagMap, subBag.name, targetBag) {
				//add found bag to a map to only count unique bags
				bagsThatHoldTarget[bag] = true
			}
		}
	}

	return len(bagsThatHoldTarget)
}

// count total number of bags a target bag contains
func part2(input []string) int {
	targetBag := "shiny gold"
	bagMap := parseInput(input)

	return countBags(bagMap, targetBag) - 1 // remove 1 to not count the initial bag itself
}
