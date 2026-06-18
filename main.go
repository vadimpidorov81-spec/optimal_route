package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"optimal-route/dijkstra"
	"optimal-route/generator"
	"optimal-route/graph"
	"optimal-route/validation"
)

const (
	Width     = 10
	Height    = 10
	MinWeight = 1
	MaxWeight = 20
	Seed      = 42

	LogFileName = "dijkstra_steps.log"
)

type testCase struct {
	Name   string
	Width  int
	Height int
	Graph  graph.Graph
	Start  string
	Target string
}

func main() {
	if err := generator.ValidateConfig(Width, Height, MinWeight, MaxWeight); err != nil {
		fmt.Println("Ошибка параметров генерации:", err)
		return
	}

	randomGraph := generator.GenerateRandomMatrixGraph(Width, Height, MinWeight, MaxWeight, Seed)
	if err := validation.ValidateGraph(randomGraph); err != nil {
		fmt.Println("Ошибка в графе:", err)
		return
	}

	logFile, err := os.Create(LogFileName)
	if err != nil {
		fmt.Println("Ошибка создания лог-файла:", err)
		return
	}
	defer logFile.Close()

	cases := []testCase{
		{
			Name:   "random-matrix",
			Width:  Width,
			Height: Height,
			Graph:  randomGraph,
			Start:  randomGraph.Nodes[0],
			Target: randomGraph.Nodes[randomGraph.NodeCount()-1],
		},
	}

	for _, tc := range cases {
		if err := runCase(tc, logFile); err != nil {
			fmt.Printf("Case: %s | error: %v\n", tc.Name, err)
		}
	}
}

func runCase(tc testCase, logWriter io.Writer) error {
	startTime := time.Now()
	result, err := dijkstra.FindShortestPath(tc.Graph, tc.Start, tc.Target, nil)
	elapsed := time.Since(startTime)
	if err != nil {
		return err
	}

	if err := writeCaseLog(tc, logWriter); err != nil {
		return err
	}
	loggedResult, err := dijkstra.FindShortestPath(tc.Graph, tc.Start, tc.Target, logWriter)
	if err != nil {
		return err
	}
	fmt.Fprintln(logWriter)

	if !result.Reachable {
		fmt.Printf(
			"Case: %s | size: %dx%d | nodes: %d | time: %s | path: not found\n",
			tc.Name,
			tc.Width,
			tc.Height,
			tc.Graph.NodeCount(),
			elapsed,
		)
		return nil
	}

	fmt.Printf(
		"Case: %s | size: %dx%d | nodes: %d | time: %s | cost: %d | path: %s\n",
		tc.Name,
		tc.Width,
		tc.Height,
		tc.Graph.NodeCount(),
		elapsed,
		result.Cost,
		strings.Join(result.Path, " -> "),
	)

	if strings.Join(result.Path, " -> ") != strings.Join(loggedResult.Path, " -> ") || result.Cost != loggedResult.Cost {
		return fmt.Errorf("результат логируемого запуска отличается от измеренного")
	}

	return nil
}

func writeCaseLog(tc testCase, writer io.Writer) error {
	_, err := fmt.Fprintf(
		writer,
		"Case: %s\nSize: %dx%d\nNodes: %d\nStart: %s\nTarget: %s\n\n%s\n",
		tc.Name,
		tc.Width,
		tc.Height,
		tc.Graph.NodeCount(),
		tc.Start,
		tc.Target,
		tc.Graph.FormatMatrix(),
	)
	return err
}
