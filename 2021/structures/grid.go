// map backed 2 dimensional integer grid
// O(1) set,access, and expansion. less optimal iterative access compared to 2d array

// .......1..
// ..1....1..
// ..1....1..
// .......1..
// .112111211
// ..........
// ..........
// ..........
// ..........
// 222111....

package structures

import (
	"errors"
	"fmt"
)

type Grid struct {
	width  int
	height int
	grid   map[string]*int
}

// coordinate returns string representation of coordinate for map lookup
func coordinate(x, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}

// NewGrid creates a new grid object
func NewGrid() *Grid {
	return &Grid{width: 1, height: 1, grid: make(map[string]*int)}
}

// NewGridInitialized creates a new grid object with values set to an initialized value
func NewGridInitialized() *Grid {
	return &Grid{width: 1, height: 1, grid: make(map[string]*int)}
}

// Width returns the length of the y coordinate of the grid
func (g Grid) Height() int {
	return g.height
}

// Width returns the width of the x coordinate of the grid
func (g Grid) Width() int {
	return g.width
}

// IncHeight expands grid to `newHeight`
// returns error if new value is less than existing value
func (g *Grid) IncHeight(newHeight int) error {
	if newHeight < g.height {
		return errors.New("new height value less than existing height")
	}
	g.height = newHeight
	return nil
}

// IncWidth expands grid to `newWidth`
// returns error if new value is less than existing value
func (g *Grid) IncWidth(newWidth int) error {
	if newWidth < g.width {
		return errors.New("new width value less than existing width")
	}
	g.width = newWidth
	return nil
}

// IncSize expands grid to new width and height, 0 value assumes no change
func (g *Grid) IncSize(newWidth, newHeight int) error {
	if newWidth != 0 {
		err := g.IncWidth(newWidth)
		if err != nil {
			return err
		}
	}
	if newHeight != 0 {
		err := g.IncHeight(newHeight)
		if err != nil {
			return err
		}
	}
	return nil
}

// inBounds returns true when the coordinates are within the grid
func inBounds(x, y int, g Grid) bool {
	// check that x and y are within bounds of grid
	if x < 0 || x > g.width {
		return false
	}
	if y < 0 || y > g.height {
		return false
	}
	return true
}

// Set a value at a coordinate in the grid, err when out of bounds
func (g *Grid) Set(x, y int, value int) error {
	// check that x and y are within bounds of grid
	if !inBounds(x, y, *g) {
		return errors.New("coordinate out of bounds: " + fmt.Sprintf("%d,%d", x, y))
	}
	g.grid[coordinate(x, y)] = &value
	return nil
}

// Get a value at a coordinate in the grid, unset values are 0
// value is always zero when out of bounds
func (g Grid) Get(x, y int) int {
	// return zero value if space is unset
	if g.grid[coordinate(x, y)] == nil {
		return 0
	}
	return *g.grid[coordinate(x, y)]
}

// GetRaw returns the raw pointer value at a coordinate
// value is always nil when out of bounds
func (g Grid) GetRaw(x, y int) *int {
	return g.grid[coordinate(x, y)]
}

func (g Grid) String() string {
	var ret string
	for y := 0; y < g.height; y++ {
		for x := 0; x < g.width; x++ {
			// print value at grid location or print nil value if unset
			if val, ok := g.grid[coordinate(x, y)]; ok {
				ret += fmt.Sprintf("%d", *val)
			} else {
				ret += "."
			}
		}
		ret += "\n"
	}
	return ret
}

func (g Grid) GoString() string {
	var ret string
	ret += "  _"
	for x := 0; x < g.width; x++ {
		ret += fmt.Sprintf("%d____", x)
	}
	ret += "\n"
	for y := 0; y < g.height; y++ {
		ret += fmt.Sprintf("%d: ", y)
		for x := 0; x < g.width; x++ {
			// print value at grid location or print nil value if unset
			if val, ok := g.grid[coordinate(x, y)]; ok {
				ret += fmt.Sprintf("%d\t", *val)
			} else {
				ret += ".\t"
			}
		}
		ret += "\n"
	}
	return ret
}
