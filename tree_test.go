package tree

import (
	"testing"

	"github.com/charmbracelet/bubbles/viewport"
)

type n struct {
	n string
	p *n
	c []*n
	s NodeState
}

var child = tn("two")
var oneWithChild = tn("one", c(child))
var oneWithChildCollapsed = tn("one collapsed", st(NodeCollapsed), c(child))
var oneWithChildCollapsedExpected = Nodes{oneWithChildCollapsed}

func tn(name string, fns ...func(*n)) *n {
	n := &n{n: name}
	for _, fn := range fns {
		fn(n)
	}
	if len(n.c) > 0 {
		n.s |= NodeCollapsible
	}
	return n
}

func st(st NodeState) func(*n) {
	return func(nn *n) {
		nn.s = st
	}
}

func c(c ...*n) func(*n) {
	return func(nn *n) {
		for i, nnn := range c {
			if i == len(c)-1 {
				nnn.s |= NodeLastChild
			}
			nnn.p = nn
			nn.c = append(nn.c, nnn)
		}
	}
}

func TestNodeState_Is(t *testing.T) {
	tests := []struct {
		name   string
		given  NodeState
		states []NodeState
		want   bool
	}{
		{
			name:   "nil",
			given:  0,
			states: nil,
			want:   true,
		},
		{
			name:   "Collapsible.Is_Collapsible",
			given:  NodeCollapsed,
			states: []NodeState{NodeCollapsed},
			want:   true,
		},
		{
			name:   "Collapsed.Is_Collapsed",
			given:  NodeCollapsed,
			states: []NodeState{NodeCollapsed},
			want:   true,
		},
		{
			name:   "Collapsed.IsNot",
			given:  NodeCollapsed,
			states: []NodeState{NodeCollapsible, NodeSelected, NodeCollapsed | NodeCollapsible},
			want:   false,
		},
		{
			name:   "Collapsed|Collapsible",
			given:  NodeCollapsed | NodeCollapsible,
			states: []NodeState{NodeCollapsible, NodeCollapsed, NodeCollapsed | NodeCollapsible},
			want:   true,
		},
		{
			name:   "Collapsed.IsNot_Collapsible|Collapsed",
			given:  NodeCollapsed,
			states: []NodeState{NodeCollapsed | NodeCollapsible},
			want:   false,
		},
		{
			name:   "Collapsed|Collapsible.IsNot",
			given:  NodeCollapsed | NodeCollapsible,
			states: []NodeState{NodeSelected, NodeSelected | NodeCollapsed},
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, st := range tt.states {
				if got := tt.given.Is(st); got != tt.want {
					t.Errorf("Is() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func mockModel(nn ...*n) Model {
	m := Model{
		viewport: viewport.New(0, 1),
		focus:    true,
	}
}

func TestModel_ToggleExpand(t *testing.T) {
	type fields struct {
		tree   *n
		cursor int
	}

	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name:    "one with child collapsed",
			fields:  fields{tree: oneWithChildCollapsed, cursor: 0},
			wantErr: false,
		},
		{
			name:    "one with child expanded",
			fields:  fields{tree: oneWithChild, cursor: 0},
			wantErr: false,
		},
		{
			name:    "/tmp",
			fields:  fields{tree: treeOne, cursor: 0},
			wantErr: false,
		},
		{
			name:    "/tmp/example1 - not expandable",
			fields:  fields{tree: treeOne, cursor: 1},
			wantErr: false,
		},
		{
			name:    "/tmp/example2 - not expandable",
			fields:  fields{tree: treeOne, cursor: 2},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := mockModel(tt.fields.tree)
			state := m.currentNode().State()
			if !state.Is(NodeCollapsible) {
				t.Skipf("Current node isn't expandable")
			}

			expected := state.Is(NodeCollapsed)
			m.ToggleExpand()

			nodeCollapsed := m.currentNode().State().Is(NodeCollapsed)
			if expected == nodeCollapsed {
				t.Errorf("Current node collapsed state = %t, expected it to be: %t", nodeCollapsed, !expected)
			}
		})
	}
}
