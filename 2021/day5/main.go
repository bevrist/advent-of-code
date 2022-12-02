// https://adventofcode.com/2021/day/5

package main

import (
	"fmt"
	"strconv"
	"strings"

	"2021/input"
	"2021/structures"
)

func main() {
	input := input.GetStringInput("input.txt")

	parsed := parseInput(input)

	fmt.Printf("Part1: %d\n", part1(parsed))
	fmt.Printf("Part2: %d\n", part2(parsed))
}

// vent stores the coordinates of each vent input
type Vent struct {
	x1, y1, x2, y2 int
}

func parseInput(input []string) []Vent {
	var vents []Vent
	for _, row := range input {
		newVent := Vent{}
		parsed := strings.Fields(row)
		// b, _ := json.Marshal(parsed)
		// fmt.Printf("%v\n", string(b))
		vent1 := strings.Split(parsed[0], `,`)
		vent2 := strings.Split(parsed[2], `,`)
		newVent.x1, _ = strconv.Atoi(vent1[0])
		newVent.y1, _ = strconv.Atoi(vent1[1])
		newVent.x2, _ = strconv.Atoi(vent2[0])
		newVent.y2, _ = strconv.Atoi(vent2[1])

		vents = append(vents, newVent)
	}
	return vents
}

// calculate how many spaces have overlapping vents
func part1(vents []Vent) int {
	g := structures.NewGrid()
	for _, vent := range vents {
		// only add vents that are aligned
		if vent.x1 == vent.x2 || vent.y1 == vent.y2 {
			// expand grid if necessary
			g.IncWidth(vent.x1 + 1)
			g.IncWidth(vent.x2 + 1)
			g.IncHeight(vent.y1 + 1)
			g.IncHeight(vent.y2 + 1)
			// is vent horizontal or vertical
			vertical := false
			if vent.x1 == vent.x2 {
				vertical = true
			}
			// draw line based on vent vertical or horizontal
			var smol, big int
			if vertical {
				if vent.y1 > vent.y2 {
					smol = vent.y2
					big = vent.y1
				} else {
					smol = vent.y1
					big = vent.y2
				}
				//draw line
				for i := smol; i <= big; i++ {
					// fmt.Printf("set:%d,%d, value:%d\n", vent.x1, i, g.Get(vent.x1, i)+1)
					g.Set(vent.x1, i, g.Get(vent.x1, i)+1) // add 1 to current value
				}
			} else { //horizontal
				if vent.x1 > vent.x2 {
					smol = vent.x2
					big = vent.x1
				} else {
					smol = vent.x1
					big = vent.x2
				}
				//draw line
				for i := smol; i <= big; i++ {
					// fmt.Printf("set:%d,%d, value:%d\n", i, vent.y1, g.Get(i, vent.y1)+1)
					g.Set(i, vent.y1, g.Get(i, vent.y1)+1) // add 1 to current value
				}
			}
		} else {
			continue // skip when vents arent aligned with grid
		}
		// fmt.Printf("%d,%d->%d,%d\n", vent.x1, vent.y1, vent.x2, vent.y2)
		// fmt.Println(g)
	}
	// fmt.Println(g)

	// count all spaces and add up any with numbers larger than 2
	total := 0
	for y := 0; y < g.Height(); y++ {
		for x := 0; x < g.Width(); x++ {
			// fmt.Printf("%d", g.Get(x, y))
			if g.Get(x, y) > 1 {
				total += 1
			}
		}
		// fmt.Printf("\n")
	}

	return total
}

// calculate how many spaces have overlapping vents, counting horizontal vents
func part2(vents []Vent) int {
	g := structures.NewGrid()
	for _, vent := range vents {
		// expand grid if necessary
		g.IncWidth(vent.x1 + 1)
		g.IncWidth(vent.x2 + 1)
		g.IncHeight(vent.y1 + 1)
		g.IncHeight(vent.y2 + 1)
		// add vents that are aligned to grid
		if vent.x1 == vent.x2 || vent.y1 == vent.y2 {
			// is vent horizontal or vertical
			vertical := false
			if vent.x1 == vent.x2 {
				vertical = true
			}
			// draw line based on vent vertical or horizontal
			var smol, big int
			if vertical {
				if vent.y1 > vent.y2 {
					smol = vent.y2
					big = vent.y1
				} else {
					smol = vent.y1
					big = vent.y2
				}
				//draw line
				for i := smol; i <= big; i++ {
					// fmt.Printf("set:%d,%d, value:%d\n", vent.x1, i, g.Get(vent.x1, i)+1)
					g.Set(vent.x1, i, g.Get(vent.x1, i)+1) // add 1 to current value
				}
			} else { //horizontal
				if vent.x1 > vent.x2 {
					smol = vent.x2
					big = vent.x1
				} else {
					smol = vent.x1
					big = vent.x2
				}
				// draw line
				for i := smol; i <= big; i++ {
					// fmt.Printf("set:%d,%d, value:%d\n", i, vent.y1, g.Get(i, vent.y1)+1)
					g.Set(i, vent.y1, g.Get(i, vent.y1)+1) // add 1 to current value
				}
			}
		} else { // add vents that are horizontal
			// find if line is heading up "/" or down "\" via slope
			slope := (vent.y2 - vent.y1) / (vent.x2 - vent.x1)
			downward := false
			if slope > 0 { // if positive slope, its upwards
				downward = true
			}
			//get leftmost point
			var leftP, rightP struct{ x, y int }
			if vent.x1 < vent.x2 {
				leftP.x = vent.x1
				leftP.y = vent.y1
				rightP.x = vent.x2
				rightP.y = vent.y2
			} else {
				leftP.x = vent.x2
				leftP.y = vent.y2
				rightP.x = vent.x1
				rightP.y = vent.y1
			}
			if downward {
				// draw downward line from leftmost point
				for i, j := leftP.x, leftP.y; i <= rightP.x; i, j = i+1, j+1 {
					// fmt.Printf("set:%d,%d, value:%d\n", i, j, g.Get(i, j)+1)
					g.Set(i, j, g.Get(i, j)+1) // add 1 to current value
				}
			} else {
				// draw upward line from leftmost point
				for i, j := leftP.x, leftP.y; i <= rightP.x; i, j = i+1, j-1 {
					// fmt.Printf("set:%d,%d, value:%d\n", i, j, g.Get(i, j)+1)
					g.Set(i, j, g.Get(i, j)+1) // add 1 to current value
				}
			}
		}
		// fmt.Printf("%d,%d->%d,%d\n", vent.x1, vent.y1, vent.x2, vent.y2)
		// fmt.Println(g)
	}
	// fmt.Println(g)

	// count all spaces and add up any with numbers larger than 2
	total := 0
	for y := 0; y < g.Height(); y++ {
		for x := 0; x < g.Width(); x++ {
			// fmt.Printf("%d", g.Get(x, y))
			if g.Get(x, y) > 1 {
				total += 1
			}
		}
		// fmt.Printf("\n")
	}

	return total
}
