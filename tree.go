package tree

// NodeState is used for passing information from a Treesih element to the view itself
type NodeState uint16

func (s NodeState) Is(st NodeState) bool {
	return s&st == st
}

func Sum(a int, b int) int {
	return a + b
}
