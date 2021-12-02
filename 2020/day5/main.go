// https://adventofcode.com/2020/day/5

package main

import (
	"2020/input"
	"fmt"
	"strconv"
)

func main() {
	input := input.GetStringInput("input.txt")

	fmt.Printf("Part1: %d\n", part1(input))
	fmt.Printf("Part2: %d\n", part2(input))
}

func parseInput(input []rune) int {
	//parse first 7 characters
	row := make([]rune, 7)
	for i := 0; i < 7; i++ {
		if input[i] == 'F' {
			row[i] = '0'
		} else {
			row[i] = '1'
		}
	}
	//parse last 3 characters
	seat := make([]rune, 3)
	for i := 0; i < 3; i++ {
		if input[i+7] == 'L' {
			seat[i] = '0'
		} else {
			seat[i] = '1'
		}
	}

	//convert binary input to integers for seats and rows
	rowNum64, err := strconv.ParseInt(string(row), 2, 0)
	if err != nil {
		panic(err)
	}
	rowNum := int(rowNum64)
	seatNum64, err := strconv.ParseInt(string(seat), 2, 0)
	if err != nil {
		panic(err)
	}
	seatNum := int(seatNum64)
	// fmt.Println(string(row))
	// fmt.Println(rowNum)
	// fmt.Println(string(seat))
	// fmt.Println(seatNum)

	//return arbitrary seat ID
	return rowNum*8 + seatNum
}

// calculate largest Seat ID
func part1(input []string) int {
	highestSeat := 0
	for _, pass := range input {
		currentSeat := parseInput([]rune(pass))
		if currentSeat > highestSeat {
			highestSeat = currentSeat
		}
	}
	return highestSeat
}

// calculate missing Seat ID
func part2(input []string) int {
	var seatMap map[int]bool = map[int]bool{}

	//find populated seats
	for _, pass := range input {
		currentSeat := parseInput([]rune(pass))
		seatMap[currentSeat] = true
	}

	//check all possible seats and return empty seat
	skipEmpty := true
	for i := 0; i < 127; i++ {
		for j := 0; j < 7; j++ {
			seatId := (i*8 + j)
			if seatMap[seatId] {
				skipEmpty = false
			} else if !seatMap[seatId] && !skipEmpty {
				return seatId
			}
		}
	}

	return -1
}
