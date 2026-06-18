package validation

import (
	"errors"
	"fmt"

	"optimal-route/graph"
)

func ValidateGraph(g graph.Graph) error {
	if len(g.Nodes) == 0 {
		return errors.New("список узлов пуст")
	}

	if len(g.Weights) != len(g.Nodes) {
		return fmt.Errorf("размер матрицы %d не совпадает с количеством узлов %d", len(g.Weights), len(g.Nodes))
	}

	seen := make(map[string]struct{}, len(g.Nodes))
	for _, node := range g.Nodes {
		if node == "" {
			return errors.New("имя узла не может быть пустым")
		}
		if _, exists := seen[node]; exists {
			return fmt.Errorf("узел %q повторяется", node)
		}
		seen[node] = struct{}{}
	}

	for i, row := range g.Weights {
		if len(row) != len(g.Nodes) {
			return fmt.Errorf("строка %d матрицы имеет длину %d вместо %d", i, len(row), len(g.Nodes))
		}

		for j, weight := range row {
			if i == j && weight != 0 {
				return fmt.Errorf("диагональный элемент [%d][%d] должен быть равен 0", i, j)
			}

			if weight < 0 && weight != graph.NoEdge {
				return fmt.Errorf("ребро %s -> %s имеет отрицательный вес %d", g.Nodes[i], g.Nodes[j], weight)
			}
		}
	}

	return nil
}
