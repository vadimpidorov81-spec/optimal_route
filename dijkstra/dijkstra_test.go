package dijkstra

import (
	"io"
	"reflect"
	"testing"

	"optimal-route/graph"
)

func TestFindShortestPath(t *testing.T) {
	g := graph.Graph{
		Nodes: []string{"A", "B", "C", "D", "E", "F"},
		Weights: [][]int{
			{0, 7, 2, graph.NoEdge, graph.NoEdge, graph.NoEdge},
			{7, 0, 3, 4, 9, graph.NoEdge},
			{2, 3, 0, graph.NoEdge, 8, graph.NoEdge},
			{graph.NoEdge, 4, graph.NoEdge, 0, 6, 5},
			{graph.NoEdge, 9, 8, 6, 0, 7},
			{graph.NoEdge, graph.NoEdge, graph.NoEdge, 5, 7, 0},
		},
	}

	result, err := FindShortestPath(g, "A", "F", io.Discard)
	if err != nil {
		t.Fatalf("FindShortestPath returned error: %v", err)
	}

	expectedPath := []string{"A", "C", "B", "D", "F"}
	if !result.Reachable {
		t.Fatal("expected target to be reachable")
	}
	if result.Cost != 14 {
		t.Fatalf("expected cost 14, got %d", result.Cost)
	}
	if !reflect.DeepEqual(result.Path, expectedPath) {
		t.Fatalf("expected path %v, got %v", expectedPath, result.Path)
	}
}

func TestFindShortestPathUnknownNode(t *testing.T) {
	g := graph.Graph{
		Nodes: []string{"A"},
		Weights: [][]int{
			{0},
		},
	}

	if _, err := FindShortestPath(g, "A", "Z", io.Discard); err == nil {
		t.Fatal("expected error for unknown target node")
	}
}
