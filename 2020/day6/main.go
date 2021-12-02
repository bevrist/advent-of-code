// https://adventofcode.com/2020/day/6

package main

import (
	"2020/input"
	"fmt"
)

func main() {
	input := input.GetStringInput("input.txt")

	fmt.Printf("Part1: %d\n", part1(input))
	fmt.Printf("Part2: %d\n", part2(input))
}

type group struct {
	numPeople int    //total number of people who answered questions
	answers   string //all people's answers concatenated together
}

//parse input to group objects
func parseInput(input []string) []group {
	var groups []group
	answers := ""
	count := 0
	for _, item := range input {
		count++
		if item == "" {
			newGroup := group{numPeople: count - 1, answers: answers}
			groups = append(groups, newGroup)
			answers = ""
			count = 0
		}
		answers += item
	}
	newGroup := group{numPeople: count, answers: answers}
	groups = append(groups, newGroup) //append final group
	return groups
}

// get count of questions where all people answered
func countAllAnswered(groups []group) []int {
	questionCount := []int{}
	for _, group := range groups {
		//add answers to map to count
		questions := map[rune]int{}
		for _, char := range group.answers {
			questions[char] += 1
		}
		//remove answers not provided by all people
		for a, b := range questions {
			if group.numPeople != b {
				delete(questions, a)
			}
		}
		//count questions answered by all users
		questionCount = append(questionCount, len(questions))
	}
	return questionCount
}

// get count of answers per group
func countAnswers(groups []group) []int {
	answerCount := []int{}
	for _, group := range groups {
		questions := map[rune]bool{}
		for _, char := range group.answers {
			questions[char] = true
		}
		answerCount = append(answerCount, len(questions))
	}
	return answerCount
}

//return sum of all items in array
func countArrayEntries(arr []int) int {
	var total int
	for _, count := range arr {
		total += count
	}
	return total
}

//get total sum of questions answered by each group
func part1(input []string) int {
	parsedInput := parseInput(input)
	questionCount := countAnswers(parsedInput)
	return countArrayEntries(questionCount)
}

//get sum of questions all people in each group answered
func part2(input []string) int {
	parsedInput := parseInput(input)
	questionCount := countAllAnswered(parsedInput)
	return countArrayEntries(questionCount)
}
