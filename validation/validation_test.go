package validation

import (
	"testing"

	"optimal-route/graph"
)

func TestValidateGraphRejectsNonSquareMatrix(t *testing.T) {
	g := graph.Graph{
		Nodes: []string{"A", "B"},
		Weights: [][]int{
			{0, 1},
		},
	}

	if err := ValidateGraph(g); err == nil {
		t.Fatal("expected error for non-square matrix")
	}
}

func TestValidateGraphRejectsInvalidDiagonal(t *testing.T) {
	g := graph.Graph{
		Nodes: []string{"A"},
		Weights: [][]int{
			{5},
		},
	}

	if err := ValidateGraph(g); err == nil {
		t.Fatal("expected error for non-zero diagonal")
	}
}

func TestValidateGraphRejectsNegativeWeight(t *testing.T) {
	g := graph.Graph{
		Nodes: []string{"A", "B"},
		Weights: [][]int{
			{0, -2},
			{1, 0},
		},
	}

	if err := ValidateGraph(g); err == nil {
		t.Fatal("expected error for negative weight")
	}
}

func TestValidateGraphAllowsNoEdge(t *testing.T) {
	g := graph.Graph{
		Nodes: []string{"A", "B"},
		Weights: [][]int{
			{0, graph.NoEdge},
			{graph.NoEdge, 0},
		},
	}

	if err := ValidateGraph(g); err != nil {
		t.Fatalf("expected NoEdge marker to be valid, got %v", err)
	}
}
