package dijkstra

import (
	"fmt"
	"io"
	"math"

	"optimal-route/graph"
)

type Result struct {
	Path      []string
	Cost      int
	Reachable bool
}

func FindShortestPath(g graph.Graph, startName, targetName string, logger io.Writer) (Result, error) {
	start, ok := g.NodeIndex(startName)
	if !ok {
		return Result{}, fmt.Errorf("стартовый узел %q не найден", startName)
	}

	target, ok := g.NodeIndex(targetName)
	if !ok {
		return Result{}, fmt.Errorf("конечный узел %q не найден", targetName)
	}

	nodeCount := len(g.Nodes)
	dist := make([]int, nodeCount)
	prev := make([]int, nodeCount)
	visited := make([]bool, nodeCount)

	for i := range dist {
		dist[i] = math.MaxInt
		prev[i] = -1
	}
	dist[start] = 0

	logln(logger, "Шаги алгоритма:")

	for step := 1; step <= nodeCount; step++ {
		current := minDistanceVertex(dist, visited)
		if current == -1 {
			logln(logger, "  Нет непосещенных достижимых узлов, поиск завершен.")
			break
		}

		visited[current] = true
		logf(logger, "  Шаг %d: выбран узел %s с текущей стоимостью %d\n", step, g.Nodes[current], dist[current])

		if current == target {
			logln(logger, "    Целевой узел достигнут, дальнейшие обновления не требуются.")
			break
		}

		for neighbor, weight := range g.Weights[current] {
			if weight == graph.NoEdge || visited[neighbor] || current == neighbor {
				continue
			}

			candidate := dist[current] + weight
			if candidate < dist[neighbor] {
				oldDistance := formatDistance(dist[neighbor])
				dist[neighbor] = candidate
				prev[neighbor] = current
				logf(
					logger,
					"    Обновляем %s: %s -> %d через %s\n",
					g.Nodes[neighbor],
					oldDistance,
					candidate,
					g.Nodes[current],
				)
				continue
			}

			logf(
				logger,
				"    %s не улучшается: текущая стоимость %s, кандидат %d через %s\n",
				g.Nodes[neighbor],
				formatDistance(dist[neighbor]),
				candidate,
				g.Nodes[current],
			)
		}
	}

	if dist[target] == math.MaxInt {
		return Result{Reachable: false}, nil
	}

	return Result{
		Path:      restorePath(g.Nodes, prev, start, target),
		Cost:      dist[target],
		Reachable: true,
	}, nil
}

func minDistanceVertex(dist []int, visited []bool) int {
	bestIndex := -1
	bestDistance := math.MaxInt

	for i, distance := range dist {
		if !visited[i] && distance < bestDistance {
			bestIndex = i
			bestDistance = distance
		}
	}

	return bestIndex
}

func restorePath(nodes []string, prev []int, start, target int) []string {
	var reversed []string

	for current := target; current != -1; current = prev[current] {
		reversed = append(reversed, nodes[current])
		if current == start {
			break
		}
	}

	path := make([]string, 0, len(reversed))
	for i := len(reversed) - 1; i >= 0; i-- {
		path = append(path, reversed[i])
	}

	return path
}

func formatDistance(distance int) string {
	if distance == math.MaxInt {
		return "∞"
	}

	return fmt.Sprintf("%d", distance)
}

func logln(writer io.Writer, message string) {
	if writer == nil {
		return
	}

	fmt.Fprintln(writer, message)
}

func logf(writer io.Writer, format string, args ...any) {
	if writer == nil {
		return
	}

	fmt.Fprintf(writer, format, args...)
}
