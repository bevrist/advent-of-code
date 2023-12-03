// https://adventofcode.com/2021/day/8

package main

import (
	"fmt"
	"strings"
)

func main() {
	// input := input.GetStringInput("input.txt")
	input := []string{
		"be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe",
		"edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc",
		"fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg",
		"fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega | efabcd cedba gadfec cb",
		"aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga | gecf egdcabf bgf bfgea",
		"fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf | gebdcfa ecba ca fadegcb",
		"dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf | cefg dcbef fcge gbcadfe",
		"bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd | ed bcgafe cdgba cbgef",
		"egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg | gbdfcae bgc cg cgb",
		"gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc | fgae cfgab fg bagce",
	}

	fmt.Printf("Part1: %d\n", part1(parseInput(input)))
	// fmt.Printf("Part2: %d\n===\n", part2(parseInput(input)))
}

type message struct {
	signal []string
	output []string
}

// extract list of integers from input
func parseInput(input []string) []message {
	var messages []message
	for _, in := range input {
		rawSignal := strings.Split(in, " | ")[0]
		rawOutput := strings.Split(in, " | ")[1]

		newMessage := message{
			signal: strings.Split(rawSignal, " "),
			output: strings.Split(rawOutput, " "),
		}
		messages = append(messages, newMessage)
		// fmt.Printf("%s\n", newMessage)
	}
	return messages
}

// digitMap maps the number of characters to the digit it represents
var digitMap = map[int]int{
	2: 1, // represents the display "1"
	4: 4, // represents the display "4"
	3: 7, // represents the display "7"
	7: 8, // represents the display "8"
}

// count the number of 1,4,7,8 in the output column
func part1(messages []message) int {
	var counter int
	for _, msg := range messages {
		// fmt.Printf("\n%s\n", msg.output)
		for _, chars := range msg.output {
			// fmt.Println(len(chars))
			// if the number is a 1,4,7,8 count it
			if digitMap[len(chars)] != 0 {
				// fmt.Printf("%s:%d, ", chars, digitMap[len(chars)])
				counter++
			}
		}
	}
	return counter
}

// ===================================================

// // use the signal bits to decode the mappings for the rest of the bits
// func part2(messages []message) int {

// 	return 1
// }
