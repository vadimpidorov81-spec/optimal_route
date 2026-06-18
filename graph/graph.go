package graph

import (
	"fmt"
	"strings"
)

const NoEdge = -1

type Graph struct {
	Nodes   []string
	Weights [][]int
}

func (g Graph) NodeIndex(name string) (int, bool) {
	for i, node := range g.Nodes {
		if node == name {
			return i, true
		}
	}

	return -1, false
}

func (g Graph) NodeCount() int {
	return len(g.Nodes)
}

func (g Graph) FormatMatrix() string {
	var builder strings.Builder

	builder.WriteString("Матрица весов (- означает, что ребра нет):\n")
	builder.WriteString("     ")
	for _, node := range g.Nodes {
		builder.WriteString(fmt.Sprintf("%4s", node))
	}
	builder.WriteString("\n")

	for i, node := range g.Nodes {
		builder.WriteString(fmt.Sprintf("%4s ", node))
		for _, weight := range g.Weights[i] {
			if weight == NoEdge {
				builder.WriteString(fmt.Sprintf("%4s", "-"))
				continue
			}
			builder.WriteString(fmt.Sprintf("%4d", weight))
		}
		builder.WriteString("\n")
	}

	return builder.String()
}
