package generator

import (
	"reflect"
	"testing"
)

func TestGenerateRandomMatrixGraphIsRepeatable(t *testing.T) {
	first := GenerateRandomMatrixGraph(3, 2, 1, 9, 100)
	second := GenerateRandomMatrixGraph(3, 2, 1, 9, 100)

	if !reflect.DeepEqual(first, second) {
		t.Fatal("expected same seed to generate identical graphs")
	}
}

func TestGenerateRandomMatrixGraphCreatesSymmetricMatrix(t *testing.T) {
	g := GenerateRandomMatrixGraph(3, 2, 1, 9, 100)

	for i := range g.Weights {
		for j := range g.Weights[i] {
			if g.Weights[i][j] != g.Weights[j][i] {
				t.Fatalf("expected symmetric weights at [%d][%d] and [%d][%d]", i, j, j, i)
			}
		}
	}
}

func TestValidateConfigRejectsInvalidWeightRange(t *testing.T) {
	if err := ValidateConfig(3, 2, 10, 1); err == nil {
		t.Fatal("expected error when min weight is greater than max weight")
	}
}
