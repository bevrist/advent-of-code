// https://adventofcode.com/2023/day/1

package main

import (
	"2023/input"
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	input := input.FileStringInput("input.txt")
	// input := []string{
	// 	"1abc2",
	// 	"pqr3stu8vwx",
	// 	"a1b2c3d4e5f",
	// 	"treb7uchet",
	// }
	// input := []string{
	// 	"two1nine",
	// 	"eightwothree",
	// 	"abcone2threexyz",
	// 	"xtwone3four",
	// 	"4nineeightseven2",
	// 	"zoneight234",
	// 	"7pqrstsixteen",
	// 	// "eighthree",
	// 	// "sevenine",
	// }

	fmt.Printf("Part1: %d\n", part1(input))
	fmt.Printf("Part2: %d\n", part2(input))
}

// part1 get first and last digit in each line and add up the total
func part1(input []string) int {
	re := regexp.MustCompile(`(\d)`)
	var total int
	for _, line := range input {
		allDigits := re.FindAllString(line, -1)
		value, _ := strconv.Atoi(allDigits[0] + allDigits[len(allDigits)-1])
		total += value
	}
	return total
}

// stringtoDigit converts string representation of number to digit
func toDigit(number string) string {
	// simply return digits as is
	if len(number) == 1 {
		return number
	}

	//convert spelled out numbers
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
	return fmt.Sprint(numberMap[number])
}

// part2 get first and last digits including spelled out
func part2(input []string) int {
	//get first match
	reFirst := regexp.MustCompile(`(one|two|three|four|five|six|seven|eight|nine|zero|\d)`)
	//get last match
	reLast := regexp.MustCompile(`.*(one|two|three|four|five|six|seven|eight|nine|zero|\d).*$`)
	var total int
	for _, line := range input {
		first := reFirst.FindAllString(line, -1)
		last := reLast.FindStringSubmatch(line)
		value, _ := strconv.Atoi(toDigit(first[0]) + toDigit(last[len(last)-1]))
		total += value
	}
	return total
}
