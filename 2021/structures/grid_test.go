package structures

import (
	"fmt"
	"testing"
)

func TestPrint(t *testing.T) {

	g := *NewGrid()
	g.IncSize(2, 4)
	g.Set(0, 0, 9)
	g.Set(1, 0, 3)
	fmt.Println(g.Get(1, 10))

	fmt.Printf("%#v\n", g)
	fmt.Printf("%v\n", g)

	result := 1
	expected := 1
	if result != expected {
		t.Errorf("got: %d, want: %d", result, expected)
	}
}
