package neural

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/wistler/aoc-2020/internal/matrix"
	"github.com/wistler/aoc-2020/internal/vector"
)

// Sigmoid function, which describes an S shaped curve.
// We pass the weighted sum of the inputs through this function to
// normalise them between 0 and 1.
func sigmoid(x float64) float64 {
	return 1 / (1 + math.Exp(-x))
}

func sigmoidV(x vector.Vec) vector.Vec {
	res := vector.Zero(len(x))
	for i := 0; i < len(x); i++ {
		res[i] = sigmoid(x[i])
	}
	return res
}

func sigmoidM(x matrix.Matrix) matrix.Matrix {
	res := matrix.Zero(len(x), len(x[0]))
	for i := 0; i < len(x); i++ {
		res[i] = sigmoidV(x[i])
	}
	return res
}

// SigmoidDerivative returns first derivative of Sigmoid fn.
// This is the gradient of the Sigmoid curve.
// It indicates how confident we are about the existing weight.
func sigmoidDerivative(x float64) float64 {
	return x * (1 - x)
}

func sigmoidDerivativeV(x vector.Vec) vector.Vec {
	res := vector.Zero(len(x))
	for i := 0; i < len(x); i++ {
		res[i] = sigmoidDerivative(x[i])
	}
	return res
}

// Network is
type Network struct {
	weights vector.Vec
}

// NewNetwork creates a new network that takes n inputs, and 1 output
func NewNetwork(n int) Network {

	nn := Network{
		weights: vector.Zero(n),
	}

	for i := 0; i < len(nn.weights); i++ {
		nn.weights[i] = 2*rand.Float64() - 1
	}

	return nn
}

// Train trains the neural net through a process of trial and error
func (nn *Network) Train(inputs matrix.Matrix, outputs vector.Vec, iterations int) error {
	if len(inputs) != len(outputs) {
		return fmt.Errorf("number of inputs (%v) do not match number of outputs (%v)", len(inputs), len(outputs))
	}
	if len(inputs[0]) != len(nn.weights) {
		return fmt.Errorf("input dimensionality (%v) not compatible with NNet (%v)", len(inputs[0]), len(nn.weights))
	}

	for i := 0; i < iterations; i++ {
		out, err := nn.ThinkM(inputs)
		if err != nil {
			return fmt.Errorf("Train: error: %v", err)
		}
		e, _ := outputs.Sum(out.Mult(-1))
		exsd, _ := matrix.New(sigmoidDerivativeV(out)).Prod(matrix.New(e))
		adj, _ := inputs.T().Dot(exsd.T())
		adjV, _ := adj.AsVec()
		// log.Printf("-- Iteration %d --\n", i)
		// log.Printf("Output (Act) = %v\n", out)
		// log.Printf("Output (Exp) = %v\n", outputs)
		// log.Printf("Output Error = %v\n", e)
		// log.Printf("Correction   = %v\n", exsd)
		// log.Printf("Adjustment   = %v\n", adjV)
		err = nn.weights.Add(adjV)
		if err != nil {
			return err
		}
	}

	return nil
}

// Think returns output for 1 node
func (nn *Network) Think(inputs vector.Vec) (float64, error) {
	dot, err := inputs.Dot(nn.weights)
	if err != nil {
		return 0, err
	}
	return sigmoid(dot), nil
}

// ThinkM processes matrix of inputs together
func (nn Network) ThinkM(inputs matrix.Matrix) (vector.Vec, error) {
	if len(nn.weights) != len(inputs[0]) {
		return nil, fmt.Errorf("Think: expected input width: %v, got: %v", len(nn.weights), len(inputs[0]))
	}
	dot, err := inputs.Dot(matrix.New(nn.weights).T())
	if err != nil {
		return nil, fmt.Errorf("Think: error occurred: %v", err)
	}
	v, err := dot.AsVec()
	if err != nil {
		return nil, fmt.Errorf("Think: error occurred: %v", err)
	}
	return sigmoidV(v), nil
}
