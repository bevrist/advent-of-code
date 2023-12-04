// https://adventofcode.com/2023/day/3

package main

import (
	"2023/input"
	"fmt"
	"math"
	"strconv"
	"unicode"
)

func main() {
	input := input.FileStringInput("input.txt")

	// input := []string{
	// 	"467..114..",
	// 	"...*......",
	// 	"..35..633.",
	// 	"......#...",
	// 	"617*......",
	// 	".....+.58.",
	// 	"..592.....",
	// 	"......755.",
	// 	"...$.*....",
	// 	".664.598..",
	// }

	parsed := parseInput(input)

	fmt.Printf("Part1: %d\n", part1(parsed))
	// fmt.Printf("Part2: %d\n", part2(parsed))
}

// number holds location and values of numbers in grid
type number struct {
	startPos coord //coordinate value starting position (leftmost digit coordinate)
	value    int   //integer value of number
	length   int   //number of digits in number
}

func (v number) GoString() string {
	return fmt.Sprintf("%#v:%d,%d", v.startPos, v.value, v.length)
}

// symbol holds symbols with coordinates
type symbol struct {
	pos    coord
	symbol string
}

func (s symbol) GoString() string {
	return fmt.Sprintf("%#v:%v", s.pos, s.symbol)
}

// coord points to a position on a grid
type coord struct {
	y int
	x int
}

func (c coord) GoString() string {
	return fmt.Sprintf("[%d,%d]", c.x, c.y)
}

// grid holds parsed info about the input
type grid struct {
	rawInput []string
	numbers  []number
	symbols  []symbol
}

// RawInput   | Values       | Symbols
// 467..114.. | [0,0]:467,3  | [1,3]:*
// ...*...... | [0,5]:114,3  | [3,6]:#
// ..35..633. | [2,2]:35,2   | [5,5]:+
// ......#... | [2,6]:633,3  | [8,3]:$
// 617*...... | [4,0]:617,3  | [8,5]:*
// .....+.58. | [5,7]:58,2   |
// ..592..... | [6,2]:592,3  |
// ......755. | [7,6]:755,3  |
// ...$.*.... | [9,1]:664,3  |
// .664.598.. | [9,5]:598,3  |

// GoString prints pretty formatted representation for grid struct
func (g grid) GoString() string {
	// make separate arrays for each printable type, then combine all at the end with measured padding
	var printVal []string = make([]string, len(g.rawInput))
	valMod := len(g.rawInput) % (len(g.numbers) + 1) // +1 to prevent divide by zero
	valCounter := 0
	var printSym []string = make([]string, len(g.rawInput))
	symMod := len(g.symbols) % len(g.rawInput) // +1 to prevent divide by zero
	symCounter := 0
	for range g.rawInput {
		// create values string array
		for i := 0; i < valMod; i++ {
			if valCounter >= len(g.numbers) {
				break // dont print out of bounds
			}
			strNumber := fmt.Sprintf("%#v", g.numbers[valCounter])
			printVal[i] += fmt.Sprintf("%-*v", len(strNumber)+1, strNumber)
			valCounter++
		}
		// create symbols string array
		for i := 0; i < symMod; i++ {
			if symCounter >= len(g.symbols) {
				break // dont print out of bounds
			}
			strSymbol := fmt.Sprintf("%#v", g.symbols[symCounter])
			printSym[i] += fmt.Sprintf("%-*v", len(strSymbol)+1, strSymbol)
			symCounter++
		}
	}
	var ret string
	// add titles at appropriate padding above printouts
	rawPad := len(g.rawInput) + 1
	valPad := int(math.Max(float64(len(printVal[0])+1), float64(7))) // minimum padding for middle row
	symPad := len(printSym[0]) + 1
	ret += fmt.Sprintf("%-*s| %-*s| %-*s\n", rawPad, "RawInput", valPad, "Values", symPad, "Symbols")
	// print raw grid, then numbers, then symbols next to each other
	for i := range g.rawInput {
		ret += fmt.Sprintf("%-*v| %-*v| %-*v\n", rawPad, g.rawInput[i], valPad, printVal[i], symPad, printSym[i])
	}
	return ret
}

// =======================================================================================================================

func parseInput(input []string) grid {
	var grid grid
	grid.rawInput = input

	// get numbers and symbols from grid
	for i, line := range grid.rawInput {
		for j := 0; j < len(line); j++ {
			// if number found, create new value struct and find length and value
			if unicode.IsDigit(rune(line[j])) {
				intNum, _ := strconv.Atoi(string(line[j]))
				newNum := number{startPos: coord{i, j}, value: intNum, length: 1}
				j++
				// check for out of bounds
				if j < len(line) {
					for unicode.IsDigit(rune(line[j])) {
						newNum.length++
						// literally concatenate digits and convert to int
						val, _ := strconv.Atoi(fmt.Sprint(newNum.value) + string(line[j]))
						newNum.value = val
						j++
						if j >= len(line) { // check for out of bounds again
							break
						}
					}
					j-- // back up if last character isnt a number so it can be checked as a symbol
				}
				grid.numbers = append(grid.numbers, newNum)
			} else { // if symbol found, create and save new symbol struct
				if string(line[j]) != "." {
					grid.symbols = append(grid.symbols, symbol{coord{i, j}, string(line[j])})
				}
			}
		}
	}

	return grid
}

// inBounds returns true if the coordinate is valid for grid
func (g grid) isValidPosition(c coord) bool {
	// negative is always invalid
	if c.x < 0 || c.y < 0 {
		return false
	}
	// greater than length of grid is invalid
	if c.x >= len(g.rawInput[0]) {
		return false
	}
	if c.y >= len(g.rawInput) {
		return false
	}
	return true
}

// part1 find what
func part1(grid grid) int {
	// fmt.Printf("%#v\n", grid)
	// make map of symbol coordinates for quick search
	symMap := map[string]symbol{}
	for _, symbol := range grid.symbols {
		symMap[fmt.Sprint(symbol.pos.x, symbol.pos.y)] = symbol
	}
	numMap := map[number]bool{} // map of number structs to avoid duplicates

	// walk around numbers searching for symbols
	for _, num := range grid.numbers {
		// walk positions around number and search for symbol
		for y := num.startPos.y - 1; y < num.startPos.y+2; y++ { //horizontal rows (y)
			for x := num.startPos.x - 1; x < num.startPos.x+num.length+1; x++ { //vertical rows (x)
				if grid.isValidPosition(coord{y, x}) {
					_, exists := symMap[fmt.Sprint(x, y)]
					if exists {
						numMap[num] = true // add to map so duplicates are skipped
					}
				}
			}
		}
	}

	// add up matching numbers
	var total int
	for num := range numMap {
		total += num.value
	}
	return total
}
