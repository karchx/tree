package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/bubbles/key"
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

func buildNodeTree(root string, maxDepth int) tree.Nodes {
  allNodes := make([]*pathNode, 0)

  rootPath := func(p string) string {
    if p == "." {
      return root
    }
    return p
  }
  fs.WalkDir(os.DirFS(root), ".", func(p string, d fs.DirEntry, err error) error {
    if err != nil {
      return fs.SkipDir
    }
    
    cnt := len(strings.Split(p, string(os.PathSeparator)))
    if maxDepth != -1 && cnt > maxDepth {
      return fs.SkipDir
    }

    st := tree.NodeNone
    if d.IsDir() {
      st |= tree.NodeCollapsible
    }
    p = rootPath(p)
    parent := findNodeByPath(allNodes, rootPath(filepath.Dir(p)))

    node := new(pathNode)
    node.path = p
    node.state = st
    node.children = make([]*pathNode, 0)

    if parent == nil {
      allNodes = append(allNodes, node)
    } else {
      node.parent = parent
      node.state |= tree.NodeCollapsed
      parent.children = append(parent.children, node)
    }
    return nil
  })

  return treeNodes(allNodes)
}

func treeNodes(pathNodes []*pathNode) tree.Nodes {
  nodes := make(tree.Nodes, len(pathNodes))
  for i, n := range pathNodes {
    nodes[i] = n
  }
  return nodes
}

func findNodeByPath(nodes []*pathNode, path string) *pathNode {
  for _, node := range nodes {
    if filepath.Clean(node.path) == filepath.Clean(path) {
      return node
    }
    if child := findNodeByPath(node.children, path); child != nil {
      return child
    }
  }
  return nil
}

type quittingTree struct {
	tree.Model
}

func (e *quittingTree) Update(m tea.Msg) (tea.Model, tea.Cmd) {
	if msg, ok := m.(tea.KeyMsg); ok && key.Matches(msg, key.NewBinding(key.WithKeys("q"))) {
		return e, tea.Quit
	}
  _, cmd := e.Model.Update(m)
	return e, cmd
}

func main() {
  var depth int
  flag.IntVar(&depth, "depth", 2, "The maximum depth to read the directory structure")
  flag.Parse()

  path := RootPath
  if flag.NArg() > 0 {
    abs, err := filepath.Abs(flag.Arg(0))
    if err != nil {
      fmt.Fprintf(os.Stderr, "%s\n", err.Error())
      os.Exit(1)
    }
    path = abs
  }

  t := tree.New(buildNodeTree(path, depth))
  m := quittingTree{Model: t}


  if err := tea.NewProgram(&m).Start(); err != nil {
    fmt.Fprintf(os.Stderr, "%s\n", err.Error())
    os.Exit(1)
  }
}
