// https://adventofcode.com/2021/day/6

package main

import (
	"fmt"
	"strconv"
	"strings"

	"2021/input"
)

func main() {
	input := input.GetStringInput("input.txt")

	fmt.Printf("Part1: %d\n", part1(parseInput(input), 80))
	fmt.Printf("Part2: %d\n", part2(parseInput(input), 256))
}

func parseInput(input []string) []int {
	split := strings.Split(input[0], ",")
	var ret []int
	for _, num := range split {
		inum, err := strconv.Atoi(num)
		i64num := int(inum)
		if err != nil {
			panic(err)
		}
		ret = append(ret, i64num)
	}
	return ret
}

func tick(fishes *[]int) {
	for i := range *fishes {
		// fmt.Printf("%d\n", i)
		// birth new fish and reset count if fish int is 0
		if (*fishes)[i] == 0 {
			(*fishes)[i] = 6
			*fishes = append(*fishes, 8)
			continue
		}
		//count down fish num
		(*fishes)[i]--
	}
}

func part1(fishes []int, days int) int {
	for i := 0; i < days; i++ {
		tick(&fishes)
	}
	return len(fishes)
}

// ===================================================

func part2(fishes []int, days int) int64 {
	var fishCount [9]int64
	// populate fishcount
	for _, fish := range fishes {
		fishCount[fish] += 1
	}
	// fmt.Printf("%#v\n", fishes)
	// fmt.Printf("%#v\n", fishCount)
	//simulate ticks
	for i := 0; i < days; i++ {
		newFish := fishCount[0]
		fishCount[0] = fishCount[1]
		fishCount[1] = fishCount[2]
		fishCount[2] = fishCount[3]
		fishCount[3] = fishCount[4]
		fishCount[4] = fishCount[5]
		fishCount[5] = fishCount[6]
		fishCount[6] = fishCount[7]
		fishCount[7] = fishCount[8]
		fishCount[8] = newFish
		fishCount[6] += newFish

		// fmt.Printf("%#v\n", fishCount)
	}
	var total int64
	for _, val := range fishCount {
		total += val
	}
	return total
}
