// https://adventofcode.com/2023/day/1

package main

import (
	"2023/input"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input := input.FileStringInput("input.txt")
	// input := []string{
	// 	"two1nine",
	// 	"eightwothree",
	// 	"abcone2threexyz",
	// 	"xtwone3four",
	// 	"4nineeightseven2",
	// 	"zoneight234",
	// 	"7pqrstsixteen",
	// 	"eighthree",
	// 	"sevenine",
	// }

	parsed := parseInput(input)

	fmt.Printf("Part1: %d\n", part1(parsed))
	fmt.Printf("Part2: %d\n", part1(parsed))
}

type line struct {
	raw      string
	firstNum int
	lastNum  *int
}

func (l line) GoString() string {
	var lastNumString string
	if l.lastNum == nil {
		lastNumString = "nil"
	} else {
		lastNumString = fmt.Sprintf("%d", *l.lastNum)
	}
	return fmt.Sprintf("{%s, %d, %s}", l.raw, l.firstNum, lastNumString)
}

// convertNumberToInteger converts string literal "numbers" to integers
func convertNumberToInteger(word string) (int, error) {
	numberMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
		"zero":  0,
	}
	// Convert to lowercase to handle case variations
	lowercaseWord := strings.ToLower(word)
	num, found := numberMap[lowercaseWord]
	if !found {
		return -1, fmt.Errorf("unsupported word: %s", word)
	}
	return num, nil
}

func stringToInt(num string) string {
	if len(num) > 1 {
		numInt, err := convertNumberToInteger(num)
		if err != nil {
			log.Fatalln("failed to convert number.", num, err)
		}
		return fmt.Sprint(numInt)
	}
	return num
}

// parseInput extracts first and last digits from input strings to store in []line
func parseInput(input []string) []line {
	var lines []line
	// get first match
	reFirst := regexp.MustCompile(`(one|two|three|four|five|six|seven|eight|nine|zero|\d)`)
	//get last match
	reLast := regexp.MustCompile(`.*(one|two|three|four|five|six|seven|eight|nine|zero|\d).*$`)
	for _, rawLine := range input {
		first := stringToInt(reFirst.FindAllString(rawLine, -1)[0])
		lastMatches := reLast.FindStringSubmatch(rawLine)
		last := stringToInt(lastMatches[len(lastMatches)-1])
		fmt.Println(first, last)
		firstNum, _ := strconv.Atoi(first)
		rawLast, _ := strconv.Atoi(last)
		lastNum := &rawLast
		newLine := line{raw: rawLine, firstNum: firstNum, lastNum: lastNum}
		fmt.Printf("%#v\n", newLine)
		lines = append(lines, newLine)
	}
	return lines
}

// get first and last digit in each line and add up the total
func part1(input []line) int {
	total := 0
	for _, line := range input {
		// fmt.Printf("%#v\n", line)
		var value int
		// concatenate the first and last numbers to form 2 digit integer
		if line.lastNum != nil {
			concatenatedStr := strconv.Itoa(line.firstNum) + strconv.Itoa(*line.lastNum)
			value, _ = strconv.Atoi(concatenatedStr)
		} else {
			value = line.firstNum
		}
		// fmt.Println(value)
		total += value
	}
	return total
}
