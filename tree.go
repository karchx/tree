package tree

// NodeState is used for passing information from a Treesih element to the view itself
type NodeState uint16

func (s NodeState) Is(st NodeState) bool {
	return s&st == st
}

const (
	NodeNone NodeState = 0

	// NodeCollapsed hints that the current node is collapsed
	NodeCollapsed NodeState = 1 << iota
	NodeSelected
	// NodeCollapsible hints that the current node can be collapsed
	NodeCollapsible
	// NodeHidden hints that the current node is not going to be displayed
	NodeHidden
	// NodeLastChild shows the node to be the last in the children list
	NodeLastChild
)

//
func DefaultSymbols() {}

func max(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return a
	}
	return b
}
