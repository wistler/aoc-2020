package neural

import (
	"math"
	"testing"

	"github.com/wistler/aoc-2020/internal/matrix"
	"github.com/wistler/aoc-2020/internal/vector"
)

func Test(t *testing.T) {
	nn := NewNetwork(3)

	inputs := matrix.New(
		vector.Make(0, 0, 1),
		vector.Make(1, 1, 1),
		vector.Make(1, 0, 1),
		vector.Make(0, 1, 1),
	)
	outputs := vector.Make(0, 1, 1, 0)

	err := nn.Train(inputs, outputs, 500)
	if err != nil {
		t.Fatalf("Training error: %v", err)
	}

	want := 1.0
	estimate, err := nn.Think(vector.Make(1, 0, 0))
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if math.Round(estimate) != want {
		t.Logf("Estimated output %v, but want %v", estimate, want)
	}
}
