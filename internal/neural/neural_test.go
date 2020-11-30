package neural

import (
	"math"
	"testing"

	"github.com/wistler/aoc-2020/internal/matrix"
	"github.com/wistler/aoc-2020/internal/vector"
)

func Test(t *testing.T) {

	trainingInputs := matrix.New(
		vector.Make(0, 0, 0),
		vector.Make(0, 0, 1),
		vector.Make(0, 1, 0),
		vector.Make(0, 1, 1),

		vector.Make(1, 0, 0),
		vector.Make(1, 0, 1),
		vector.Make(1, 1, 0),
		vector.Make(1, 1, 1),
	)
	trainingOutputs := matrix.New(
		vector.Make(0, 0),
		vector.Make(0, 1),
		vector.Make(1, 0),
		vector.Make(1, 1),

		vector.Make(0, 0),
		vector.Make(0, 1),
		vector.Make(1, 0),
		vector.Make(1, 1),
	)

	testInputs := matrix.New(
		// vector.Make(0, 0, 0),
		// vector.Make(0, 0, 1),
		vector.Make(0, 1, 0),
		vector.Make(0, 1, 1),

		// vector.Make(1, 0, 0),
		vector.Make(1, 0, 1),
		vector.Make(1, 1, 0),
		// vector.Make(1, 1, 1),
	)
	testOutputs := matrix.New(
		// vector.Make(0, 0),
		// vector.Make(0, 1),
		vector.Make(1, 0),
		vector.Make(1, 1),

		// vector.Make(0, 0),
		vector.Make(0, 1),
		vector.Make(1, 0),
		// vector.Make(1, 1),
	)

	nn := CreateNetwork(len(trainingInputs[0]), len(trainingOutputs[0]))
	err := nn.Train(trainingInputs, trainingOutputs, 250, 0.10)
	if err != nil {
		t.Fatalf("Training error: %v", err)
	}

	for i := 0; i < len(testInputs); i++ {
		testInput := testInputs[i]
		testOutput := testOutputs[i]

		estimate, err := nn.Think(matrix.New(testInput))
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		// t.Logf("Estimated output %v", estimate)
		for idx, eo := range estimate[0] {
			if math.Round(eo) != testOutput[idx] {
				t.Logf("Estimated output %v, but want %v", estimate, testOutput)
				break
			}
		}
	}

}
