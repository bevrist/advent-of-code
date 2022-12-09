// https://adventofcode.com/2022/day/8

package main

import (
	"2022/input"
	"2022/structures"
	"fmt"
	"strconv"
)

func main() {
	input := input.GetStringInput("input.txt")

	// fmt.Printf("Part1: %d\n", part1(input))
	fmt.Printf("Part2: %d\n", part2(input))
}

type Tree struct {
	height       int
	visFromNorth bool
	visFromEast  bool
	visFromSouth bool
	visFromWest  bool
	senicNorth   int
	senicEast    int
	senicSouth   int
	senicWest    int
}

func (t Tree) String() string {
	return fmt.Sprint(t.height)
}
func (t Tree) GoString() string {
	// // print letter representing cardinal direction is visible, otherwise print "."
	// var n, e, s, w string = ".", ".", ".", "."
	// if t.visFromNorth {
	// 	n = "N"
	// }
	// if t.visFromEast {
	// 	e = "E"
	// }
	// if t.visFromSouth {
	// 	s = "S"
	// }
	// if t.visFromWest {
	// 	w = "W"
	// }
	// // cardinal visibility
	// return fmt.Sprintf("{%d,%s%s%s%s}", t.height, n, e, s, w)
	// senic scores for each direction
	return fmt.Sprintf("{%d,%d%d%d%d}", t.height, t.senicNorth, t.senicEast, t.senicSouth, t.senicWest)
}

func parseInput(in []string) *structures.Grid[Tree] {
	g := structures.NewGrid[Tree]()
	g.IncHeight(len(in))
	for i, row := range in {
		g.IncWidth(len(row))
		for j, num := range row {
			h, _ := strconv.Atoi(string(num))
			newTree := Tree{height: h}
			g.Set(j, i, newTree)
		}
	}
	return g
}

func getTreeVisibility(g *structures.Grid[Tree]) {
	//calculate if trees are visible from each direction
	//from north
	for i := 0; i < g.Width(); i++ {
		nBiggest := -1
		for j := 0; j < g.Height(); j++ {
			if currTree := g.Get(i, j); currTree.height > nBiggest {
				nBiggest = currTree.height
				currTree.visFromNorth = true
				g.Set(i, j, currTree)
			}
		}
	}
	//from west
	for j := 0; j <= g.Height(); j++ {
		wBiggest := -1
		for i := 0; i < g.Width(); i++ {
			if currTree := g.Get(i, j); currTree.height > wBiggest {
				wBiggest = currTree.height
				currTree.visFromWest = true
				g.Set(i, j, currTree)
			}
		}
	}
	//from east
	for j := g.Height() - 1; j >= 0; j-- {
		eBiggest := -1
		for i := g.Width() - 1; i >= 0; i-- {
			if currTree := g.Get(i, j); currTree.height > eBiggest {
				eBiggest = currTree.height
				currTree.visFromEast = true
				g.Set(i, j, currTree)
			}
		}
	}
	//from south
	for i := g.Width() - 1; i >= 0; i-- {
		sBiggest := -1
		for j := g.Height() - 1; j >= 0; j-- {
			if currTree := g.Get(i, j); currTree.height > sBiggest {
				sBiggest = currTree.height
				currTree.visFromSouth = true
				g.Set(i, j, currTree)
			}
		}
	}
}

func part1(in []string) int {
	g := parseInput(in)
	getTreeVisibility(g)

	// count number of trees that are visible
	visCount := 0
	for j := 0; j < g.Height(); j++ {
		for i := 0; i < g.Width(); i++ {
			if t := g.Get(i, j); t.visFromNorth || t.visFromEast || t.visFromSouth || t.visFromWest {
				visCount++
			}
		}
	}

	// fmt.Printf("%v\n", g)
	// fmt.Printf("%#v\n", g)
	return visCount
}

func getSenicScores(g *structures.Grid[Tree]) {
	// for every tree in grid, calculate visibility of other trees in each direction from the treehouse
	for j := 0; j < g.Height(); j++ {
		// fmt.Println()
		for i := 0; i < g.Width(); i++ {
			// fmt.Println()
			// fmt.Printf("(%v,%v)%v|", i, j, g.Get(i, j).height)
			treeHouse := g.Get(i, j)
			// west view
			for y := i - 1; y >= 0; y-- {
				// fmt.Printf("[%v,%v]{%v,%v}(%v)(%v)", i, j, y, j, g.Get(i, j).height, g.Get(y, j).height)
				// tree is equal or taller than treehouse, stop looking west
				if currTree := g.Get(y, j); currTree.height >= treeHouse.height {
					treeHouse.senicWest++
					break
				} else { // tree is smaller, add to count
					treeHouse.senicWest++
				}
			}
			// north view
			for y := j - 1; y >= 0; y-- {
				// fmt.Printf("[%v,%v]{%v,%v}(%v)(%v)", i, j, y, j, g.Get(i, j).height, g.Get(y, j).height)
				// tree is equal or taller than treehouse, stop looking north
				if currTree := g.Get(i, y); currTree.height >= treeHouse.height {
					treeHouse.senicNorth++
					break
				} else { // tree is smaller, add to count
					treeHouse.senicNorth++
				}
			}
			// east view
			for y := i + 1; y <= g.Width()-1; y++ {
				// fmt.Printf("[%v,%v]{%v,%v}(%v)(%v)", i, j, y, j, g.Get(i, j).height, g.Get(y, j).height)
				// tree is equal or taller than treehouse, stop looking east
				if currTree := g.Get(y, j); currTree.height >= treeHouse.height {
					treeHouse.senicEast++
					break
				} else { // tree is smaller, add to count
					treeHouse.senicEast++
				}
			}
			// south view
			for y := j + 1; y <= g.Height()-1; y++ {
				// fmt.Printf("[%v,%v]{%v,%v}(%v)(%v)", i, j, x, j, g.Get(i, j).height, g.Get(x, j).height)
				// tree is equal or taller than treehouse, stop looking south
				if currTree := g.Get(i, y); currTree.height >= treeHouse.height {
					treeHouse.senicSouth++
					break
				} else { // tree is smaller, add to count
					treeHouse.senicSouth++
				}
			}
			g.Set(i, j, treeHouse)
		}
	}
}

func part2(in []string) int {
	g := parseInput(in)
	getSenicScores(g)

	//calculate senic scores
	highSenicScore := 0
	for j := 0; j < g.Height(); j++ {
		for i := 0; i < g.Width(); i++ {
			t := g.Get(i, j)
			score := t.senicNorth * t.senicEast * t.senicSouth * t.senicWest
			if score > highSenicScore {
				highSenicScore = score
			}
		}
	}

	// fmt.Printf("%#v\n", g)

	return highSenicScore
}
