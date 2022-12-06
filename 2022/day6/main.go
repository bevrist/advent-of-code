// https://adventofcode.com/2022/day/6

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

// find when first instance of 4 unique letters is found
func part1(in []string) int {
	var buf [4]string // a ring buffer
	head := 0         // ring buffer pointer
	// fmt.Println("in", in[0])
	for i, v := range in[0] {
		char := string(v)
		// populate first 3 chars of buffer
		if i < 3 {
			buf[i] = char
			head = 3
		} else {
			//add chars through ring buffer
			buf[head] = char
			if head++; head == 4 {
				head = 0
			}
			// fmt.Printf("%v\n", buf)
			// check if all 4 chars are unique
			chk := make(map[string]bool)
			unique := true
			for _, char := range buf {
				if _, ok := chk[char]; ok { // if char exists in map then it repeated
					unique = false
				}
				chk[char] = true
			}
			if unique {
				return i + 1
			}
		}
	}
	return -999 // no unique sequence found
}

// find when first instance of 14 unique letters is found
func part2(in []string) int {
	var buf [14]string // a ring buffer
	head := 0          // ring buffer pointer
	// fmt.Println("in", in[0])
	for i, v := range in[0] {
		char := string(v)
		// populate first 3 chars of buffer
		if i < 13 {
			buf[i] = char
			head = 13
		} else {
			//add chars through ring buffer
			buf[head] = char
			if head++; head == 14 {
				head = 0
			}
			// fmt.Printf("%v\n", buf)
			// check if all 4 chars are unique
			chk := make(map[string]bool)
			unique := true
			for _, char := range buf {
				if _, ok := chk[char]; ok { // if char exists in map then it repeated
					unique = false
				}
				chk[char] = true
			}
			if unique {
				return i + 1
			}
		}
	}
	return -999 // no unique sequence found
}
