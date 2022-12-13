// https://adventofcode.com/2022/day/8

package main

import (
	"2022/input"
	"2022/structures"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	input := input.GetStringInput("input.txt")

	fmt.Printf("Part1: %d\n", part1(input))
	fmt.Printf("Part2: %d\n", part2(input))
}

// Pos a coordinate position
type Pos struct {
	x int
	y int
}

// Location stores data for point on grid
type Location struct {
	tCount int
}

// declare cardinal type and constants
type Direction string

const (
	up    Direction = "U"
	right           = "R"
	down            = "D"
	left            = "L"
)

// Command from input
type Command struct {
	direction Direction
	distance  int
}

// get list of commands
func parseInput(in []string) []Command {
	var c []Command
	for _, row := range in {
		split := strings.Fields(row)
		dir := split[0]
		count, _ := strconv.Atoi(split[1])
		c = append(c, Command{direction: Direction(dir), distance: count})
	}
	return c
}

func printGrid(g structures.Grid[Location], head, tail Pos) {
	for y := g.Height() - 1; y >= 0; y-- {
		for x := 0; x < g.Width(); x++ {
			if x == head.x && y == head.y {
				fmt.Print("H")
			} else if x == tail.x && y == tail.y {
				fmt.Print("T")
			} else {
				fmt.Printf("%v", g.Get(x, y).tCount)
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func updateTail(g *structures.Grid[Location], tail, head, lastHead Pos) Pos {
	ret := tail
	// move tail to last head position if more than 1 position away
	xDistance := math.Abs(float64(tail.x - head.x))
	yDistance := math.Abs(float64(tail.y - head.y))
	if xDistance > 1 || yDistance > 1 {
		// fmt.Printf("x: %v, y: %v\n", xDistance, yDistance)
		// increment Location count
		g.Set(lastHead.x, lastHead.y, Location{g.Get(lastHead.x, lastHead.y).tCount + 1})
		ret = lastHead
	}
	return ret
}

func part1(in []string) int {
	g := structures.NewGrid[Location]()
	g.IncSize(1000, 1000) // give it a minimum size
	startSpace := Pos{500, 500}
	g.Set(startSpace.x, startSpace.y, Location{1}) // mark starting space as visited
	c := parseInput(in)

	headLoc := startSpace
	var lastHead Pos
	tailLoc := startSpace

	// fmt.Println("== Initial State ==")
	// printGrid(*g, headLoc, tailLoc)

	for _, cmd := range c {
		// fmt.Printf("== %v %v ==\n", cmd.direction, cmd.distance)
		// move head according to command
		for j := 0; j < cmd.distance; j++ {
			// move head number of times
			lastHead = headLoc
			switch cmd.direction {
			case up:
				// if headLoc.y >= g.Height()-1 {
				// 	g.IncHeight(g.Height() + 1)
				// }
				headLoc.y++
			case right:
				// if headLoc.x >= g.Width() {
				// 	g.IncWidth(g.Width() + 1)
				// }
				headLoc.x++
			case down:
				headLoc.y--
			case left:
				headLoc.x--
			}
			tailLoc = updateTail(g, tailLoc, headLoc, lastHead)
			// printGrid(*g, headLoc, tailLoc)
		}
	}

	// fmt.Println("== Final State ==")
	// printGrid(*g, Pos{-1, -1}, Pos{-1, -1})

	// count unique visited spaces
	total := 0
	for y := 0; y < g.Height(); y++ {
		for x := 0; x < g.Width(); x++ {
			if g.Get(x, y).tCount > 0 {
				total++
			}
		}
	}
	return total
}

// =============================================

func printGrid2(g structures.Grid[Location], rope [10]Pos) {
	for y := g.Height() - 1; y >= 0; y-- {
		for x := 0; x < g.Width(); x++ {
			passed := false
			for i := 0; i < len(rope); i++ {
				if x == rope[i].x && y == rope[i].y {
					fmt.Print(i)
					passed = true
					break
				}
			}
			if !passed {
				if g.Get(x, y).tCount == 0 {
					fmt.Print(".")
				} else {
					fmt.Print("#")
				}
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

// Clamp returns f clamped to [low, high]
func Clamp(f, low, high float64) float64 {
	if f < low {
		return low
	}
	if f > high {
		return high
	}
	return f
}

func updateBody(g *structures.Grid[Location], rope [10]Pos, direction Pos) [10]Pos {
	//update head
	rope[0].x += direction.x
	rope[0].y += direction.y
	// fmt.Println()
	for i := 1; i < len(rope); i++ {
		// move part to last head position if more than 1 position away
		xDistance := float64(rope[i-1].x - rope[i].x)
		yDistance := float64(rope[i-1].y - rope[i].y)
		if (math.Abs(xDistance) >= 1 && math.Abs(yDistance) > 1) || (math.Abs(xDistance) > 1 && math.Abs(yDistance) >= 1) {
			// move diagonal
			// fmt.Printf("\ndiag: %v, pos: %v, front: %v, clampx/y %v %v", i, rope[i], rope[i-1], xDistance, yDistance)
			rope[i].x += int(Clamp(xDistance, -1, 1))
			rope[i].y += int(Clamp(yDistance, -1, 1))
			// fmt.Printf(" new: %v", rope[i])
			if i == 9 {
				// increment Location count if tail moved
				g.Set(rope[i].x, rope[i].y, Location{g.Get(rope[i].x, rope[i].y).tCount + 1})
			}
		} else if math.Abs(xDistance) > 1 || math.Abs(yDistance) > 1 {
			// fmt.Printf("\nstraight: %v, pos: %v, front: %v, clampx/y %v %v", i, rope[i], rope[i-1], int(Clamp(xDistance, -1, 1)), int(Clamp(yDistance, -1, 1)))
			//move vertical/horizontal
			// fmt.Printf("x: %v, y: %v\n", xDistance, yDistance)
			rope[i].x += int(Clamp(xDistance, -1, 1))
			rope[i].y += int(Clamp(yDistance, -1, 1))
			// fmt.Printf(" new: %v", rope[i])
			if i == 9 {
				// increment Location count if tail moved
				g.Set(rope[i].x, rope[i].y, Location{g.Get(rope[i].x, rope[i].y).tCount + 1})
			}
		} else {
			// fmt.Println("no move")
			// fmt.Printf("|skip:%v,%v", i, rope[i])
		}
	}
	// fmt.Println()
	return rope
}

func part2(in []string) int {
	g := structures.NewGrid[Location]()
	g.IncSize(1000, 1000) // give it a minimum size
	startSpace := Pos{500, 500}
	g.Set(startSpace.x, startSpace.y, Location{1}) // mark starting space as visited
	c := parseInput(in)

	var rope [10]Pos
	for i := 0; i < len(rope); i++ {
		rope[i] = startSpace
	}
	var direction Pos

	// fmt.Println("== Initial State ==")
	// printGrid2(*g, rope)

	for _, cmd := range c {
		// fmt.Printf("== %v %v ==\n", cmd.direction, cmd.distance)
		// move head according to command
		for j := 0; j < cmd.distance; j++ {
			// move head number of times
			switch cmd.direction {
			case up:
				direction.y = 1
				direction.x = 0
			case right:
				direction.x = 1
				direction.y = 0
			case down:
				direction.y = -1
				direction.x = 0
			case left:
				direction.x = -1
				direction.y = 0
			}
			rope = updateBody(g, rope, direction)
			// printGrid2(*g, rope)
		}
		// fmt.Println(rope)
		// printGrid2(*g, rope)
	}

	// fmt.Println("== Final State ==")
	// printGrid2(*g, rope)
	// printGrid2(*g, [10]Pos{})

	// count unique visited spaces
	total := 0
	for y := 0; y < g.Height(); y++ {
		for x := 0; x < g.Width(); x++ {
			if g.Get(x, y).tCount > 0 {
				total++
			}
		}
	}
	return total
}

// 2423 too low
