package tree

import "testing"

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
