package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/karchx/tree"
)

const RootPath = "/tmp"

type pathNode struct {
  parent *pathNode
  path string
  state tree.NodeState
  children []*pathNode
}

func (n *pathNode) Parent() tree.Node {

  return n.parent
}

func (n *pathNode) Init() tea.Cmd {
  return nil
}

const (
  Collapsed = "⊞"
	Expanded  = "⊟"
)

func (n *pathNode) Children() tree.Nodes {
  return treeNodes(n.children)
}

func (n *pathNode) View() string {
  return ""
}

func (n *pathNode) State() tree.NodeState {
  return n.state
}

func (n *pathNode) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  switch m := msg.(type) {
  case tree.NodeState:
    n.state = m
  }
  return n, nil
}

func treeNodes(pathNodes []*pathNode) tree.Nodes {
  nodes := make(tree.Nodes, len(pathNodes))
  for i, n := range pathNodes {
    nodes[i] = n
  }
  return nodes
}

func main() {
  if err := tea.NewProgram(nil).Start(); err != nil {
    fmt.Fprintf(os.Stderr, "%s\n", err.Error())
    os.Exit(1)
  }
}
