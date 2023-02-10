package tree

import "testing"

func Test_showTreeSymbolAtPos(t *testing.T) {
	got := Sum(1, 2)
	if got != 3 {
		t.Errorf("Invalid result %d", got)
	}
}
