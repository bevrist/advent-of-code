// https://adventofcode.com/2022/day/3

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

func part1(in []string) int {
	score := 0
	for _, line := range in {
		// split bag into 2
		l := line[:len(line)/2]
		r := line[len(line)/2:]
		// fmt.Printf("l: %v, r: %v\n", l, r)

		// find shared character
		var shared string
		done := false
		for _, a := range l {
			for _, b := range r {
				if a == b {
					shared = string(a)
					done = true
					break
				}
			}
			if done {
				break
			}
		}

		// get score and add to total
		score += priorityScore[shared]
	}
	return score
}

func part2(in []string) int {
	score := 0
	// read 3 lines at a time
	for i := 0; i < len(in); i += 3 {
		// fmt.Printf("i:%v ,%v %v %v\n", i, in[i], in[i+1], in[i+2])
		//add all items to compare map to find duplicate across all 3
		var cmp map[string]int = make(map[string]int)
		for j := 0; j < 3; j++ { // 3 lists at a time
			var curr map[rune]bool = make(map[rune]bool)
			for _, char := range in[i+j] {
				if _, found := curr[char]; !found { //only add char to cmp map once per list
					curr[char] = true
					cmp[string(char)]++
				}
			}
		}
		// fmt.Printf("cmp: %#v\n", cmp)
		// find item in cmp map that was in all 3 lists
		var cmn string
		for key, char := range cmp {
			if char == 3 {
				cmn = string(key)
			}
		}

		// lookup score and add to total
		score += priorityScore[cmn]
	}
	return score
}
