package structures

import (
	"fmt"
	"testing"
)

// create a test grid with some initialization
// 9,0
// 3,.(nil)
func new2x2Grid() *Grid {
	g := NewGrid()
	g.IncSize(2, 4)
	g.Set(0, 0, 9)
	g.Set(0, 1, 0)
	g.Set(1, 0, 3)
	g.Set(1, 3, 1)
	return g
}

func TestPrint(t *testing.T) {
	g := new2x2Grid()

	//manual print grid
	fmt.Println("Manual print")
	fmt.Printf("%d%d\n", g.Get(0, 0), g.Get(1, 0))
	fmt.Printf("%d%d\n", g.Get(0, 1), g.Get(1, 1))
	fmt.Printf("%d%d\n", g.Get(0, 2), g.Get(1, 2))
	fmt.Printf("%d%d", g.Get(0, 3), g.Get(1, 3))
	fmt.Printf(".%d.\n\n", g.Get(-1, -1))

	//manual print rawgrid
	fmt.Println("Manual print raw 2x2 (0 means nil pointer)")
	fmt.Printf("%d,%d\n", g.GetRaw(0, 0), g.GetRaw(0, 1))
	fmt.Printf("%d,%d,%d\n\n", g.GetRaw(1, 0), g.GetRaw(1, 1), g.GetRaw(-1, -1))

	fmt.Println("%#V")
	fmt.Printf("%#v\n", g)
	fmt.Println("%V")
	fmt.Printf("%v\n", g)

	result := 1
	expected := 1
	if result != expected {
		t.Errorf("got: %d, want: %d", result, expected)
	}
}
