package generator

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"

	"optimal-route/graph"
)

func ValidateConfig(width, height, minWeight, maxWeight int) error {
	if width <= 0 {
		return errors.New("ширина должна быть больше 0")
	}
	if height <= 0 {
		return errors.New("высота должна быть больше 0")
	}
	if minWeight < 0 {
		return errors.New("минимальный вес должен быть неотрицательным")
	}
	if maxWeight < 0 {
		return errors.New("максимальный вес должен быть неотрицательным")
	}
	if minWeight > maxWeight {
		return fmt.Errorf("минимальный вес %d больше максимального %d", minWeight, maxWeight)
	}

	return nil
}

func GenerateRandomMatrixGraph(width, height, minWeight, maxWeight int, seed int64) graph.Graph {
	if err := ValidateConfig(width, height, minWeight, maxWeight); err != nil {
		panic(err)
	}

	nodeCount := width * height
	nodes := make([]string, nodeCount)
	weights := make([][]int, nodeCount)

	nameWidth := len(strconv.Itoa(nodeCount - 1))
	for i := 0; i < nodeCount; i++ {
		nodes[i] = fmt.Sprintf("N%0*d", nameWidth, i)
		weights[i] = make([]int, nodeCount)
	}

	random := rand.New(rand.NewSource(seed))
	for i := 0; i < nodeCount; i++ {
		for j := i; j < nodeCount; j++ {
			if i == j {
				weights[i][j] = 0
				continue
			}

			weight := minWeight + random.Intn(maxWeight-minWeight+1)
			weights[i][j] = weight
			weights[j][i] = weight
		}
	}

	return graph.Graph{
		Nodes:   nodes,
		Weights: weights,
	}
}
