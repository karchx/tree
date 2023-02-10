package tree

import (
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// NodeState is used for passing information from a Treesih element to the view itself
type NodeState uint16

type Model struct {
	focus  bool
	cursor int
	tree   Nodes

	viewport viewport.Model
}

type Node interface {
	tea.Model
	Parent() Node
	Children() Nodes
	State() NodeState
}

type Nodes []Node

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

var (
	width = lipgloss.Width

	defaultStyle         = lipgloss.NewStyle()
	defaultSelectedStyle = defaultStyle.Reverse(true)
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
