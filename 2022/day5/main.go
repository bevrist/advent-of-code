// https://adventofcode.com/2022/day/5

package main

import (
	"2022/input"
	"fmt"
	"strconv"
	"strings"
)

type Step struct {
	count int
	from  int
	dest  int
}

type Stacks struct {
	stacks [][]string
	steps  []Step
}

func main() {
	input := input.GetStringInput("input.txt")

	fmt.Printf("Part1: %s\n", part1(input))
	fmt.Printf("Part2: %s\n", part2(input))
}

func parseInput(in []string) Stacks {
	// get starting stacks
	// get length of starting stacks
	var stackCount, stackHeight int
	for i := 0; i < len(in); i++ {
		if len(in[i]) == 0 {
			stackHeight = i - 1
			labels := strings.Fields(in[i-1])
			count, err := strconv.Atoi(labels[len(labels)-1])
			if err != nil {
				panic(err)
			}
			stackCount = count
			break
		}
	}
	// fmt.Println("Stack Count: ", stackCount)
	// fmt.Println("Stack Height: ", stackHeight)
	// load stacks
	var newStacks Stacks
	newStacks.stacks = make([][]string, stackCount)
	for i := 0; i < stackHeight; i++ {
		// read rows
		counter := 0
		for j := 1; j < (stackCount * 4); j += 4 {
			if len(in[i]) < j { //handle hanging blank lines in input
				// fmt.Printf(" ")
				break
			}
			if string(in[i][j]) == " " { // dont add empty entries
				// fmt.Printf(" ")
				counter++
				continue
			}
			// fmt.Printf("%v", string(in[i][j]))
			t := []string{string(in[i][j])}
			newStacks.stacks[counter] = append(t, newStacks.stacks[counter]...)
			counter++
		}
		// fmt.Print("|\n")
	}

	// populate instruction list
	for i := stackHeight + 2; i < len(in); i++ {
		f := strings.Fields(in[i])
		// fmt.Println("instruction: ", f)
		cnt, _ := strconv.Atoi(f[1])
		frm, _ := strconv.Atoi(f[3])
		to, _ := strconv.Atoi(f[5])
		newStep := Step{
			count: cnt,
			from:  frm - 1,
			dest:  to - 1,
		}
		newStacks.steps = append(newStacks.steps, newStep)
	}
	return newStacks
}

func part1(in []string) string {
	parsed := parseInput(in)

	// perform steps
	for _, ins := range parsed.steps {
		for i := 0; i < ins.count; i++ {
			// fmt.Printf("%+v\n", ins)
			// fmt.Printf("%+v\n", parsed.stacks)
			// remove from stack
			rm := len(parsed.stacks[ins.from]) - 1
			tmp := parsed.stacks[ins.from][rm]
			parsed.stacks[ins.from] = parsed.stacks[ins.from][:rm]
			parsed.stacks[ins.dest] = append(parsed.stacks[ins.dest], tmp)
		}
	}

	//get top items from each stack
	ret := ""
	for _, s := range parsed.stacks {
		ret += s[len(s)-1]
	}

	// fmt.Printf("%+v\n", parsed)
	return ret
}

func part2(in []string) string {
	parsed := parseInput(in)

	// perform steps
	for _, ins := range parsed.steps {
		//move multiple containers at once
		var move []string
		for i := 0; i < ins.count; i++ {
			// fmt.Printf("%+v\n", ins)
			// fmt.Printf("%+v\n", parsed.stacks)
			// remove group from stack in order
			rm := len(parsed.stacks[ins.from]) - 1
			move = append([]string{parsed.stacks[ins.from][rm]}, move...)
			parsed.stacks[ins.from] = parsed.stacks[ins.from][:rm]
		}
		// append entire group on new stack
		parsed.stacks[ins.dest] = append(parsed.stacks[ins.dest], move...)
	}

	//get top items from each stack
	ret := ""
	for _, s := range parsed.stacks {
		ret += s[len(s)-1]
	}

	// fmt.Printf("%+v\n", parsed)
	return ret
}
