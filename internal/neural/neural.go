package neural

import (
	"fmt"
	"log"
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
		for j := 0; j < len(x[0]); j++ {
			res[i][j] = sigmoid(x[i][j])
		}
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

func sigmoidDerivativeM(x matrix.Matrix) matrix.Matrix {
	res := matrix.Zero(len(x), len(x[0]))
	for i := 0; i < len(x); i++ {
		for j := 0; j < len(x[0]); j++ {
			res[i][j] = sigmoidDerivative(x[i][j])
		}
	}
	return res
}

// Network is
type Network struct {
	nIn    int
	nOut   int
	layer0 matrix.Matrix
}

// CreateNetwork creates a new network with 1 layer, that takes nIn inputs, and nOut output
func CreateNetwork(nIn, nOut int) Network {
	nn := Network{
		nIn:    nIn,
		nOut:   nOut,
		layer0: createNetworkLayer(nIn, nOut),
	}
	return nn
}

// a layer is a matrix of size [nOut x nIn]
func createNetworkLayer(nIn, nOut int) []vector.Vec {
	layer := matrix.Zero(nOut, nIn)
	for o := 0; o < nOut; o++ {
		for i := 0; i < nIn; i++ {
			layer[o][i] = 2*rand.Float64() - 1
		}
	}
	return layer
}

// Train trains the neural net through a process of trial and error
// Given NNet shape is [nIn, ... nOut]
// Input matrix must be of size [? x nIn]
// Result matrix is of size [? x nOut]
func (nn *Network) Train(trainingInputs matrix.Matrix, trainingOutputs matrix.Matrix, maxIterations int, maxError float64) error {
	if len(trainingInputs) != len(trainingOutputs) {
		return fmt.Errorf("number of input sets (%v) do not match number of output sets (%v)", len(trainingInputs), len(trainingOutputs))
	}
	if len(trainingInputs[0]) != nn.nIn {
		return fmt.Errorf("input dimensionality (%v) not compatible with NNet (nIn = %v)", len(trainingInputs[0]), nn.nIn)
	}
	if len(trainingOutputs[0]) != nn.nOut {
		return fmt.Errorf("output dimensionality (%v) not compatible with NNet (nOut = %v)", len(trainingOutputs[0]), nn.nOut)
	}

	for i := 0; i < maxIterations; i++ {
		o, _ := nn.Think(trainingInputs)        // out := Think([? x nIn]) => [? x nOut]
		e, _ := trainingOutputs.Sum(o.Mult(-1)) // e := [? x nOut] - [? x nOut]
		se, _ := sigmoidDerivativeM(o).Prod(e)  // exsd := [? x nOut] * [? x nOut]
		adj, _ := trainingInputs.T().Dot(se)    // adj := [? x nIn]T · [? x nOut] = [nIn x nOut]
		nn.layer0, _ = nn.layer0.Sum(adj.T())   // nn.layer0 := [nOut x nIn] + [nIn x nOut]T

		withinErrorThreshold := true
		for _, er := range e {
			for _, err := range er {
				if math.Abs(err) > maxError {
					withinErrorThreshold = false
					break
				}
			}
			if !withinErrorThreshold {
				break
			}
		}
		// if i%10000 == 0 || withinErrorThreshold {
		// 	log.Printf("-- Iteration %d --\n", i)
		// 	log.Printf("Output (Act) = %v\n", o)
		// 	log.Printf("Output (Exp) = %v\n", trainingOutputs)
		// 	log.Printf("Output Error = %v\n", e)
		// 	log.Printf("Correction   = %v\n", se)
		// 	log.Printf("Adjustment   = %v\n", adj)
		// }

		if withinErrorThreshold {
			log.Printf("Error threshold reached. Stopping at iteration: %v", i)
			break
		}
	}

	log.Printf("New Weights  = %v\n", nn.layer0)
	return nil
}

// Think processes matrix of inputs together
// Input matrix must be of size [? x nIn]
// Result matrix is of size [? x nOut]
func (nn Network) Think(inputs matrix.Matrix) (matrix.Matrix, error) {
	if len(inputs[0]) != nn.nIn {
		return nil, fmt.Errorf("input dimensionality (%v) not compatible with NNet (nIn = %v)", len(inputs[0]), nn.nIn)
	}
	// [? x nIn] · [nOut x nIn]T => [? x nOut]
	dot, err := inputs.Dot(nn.layer0.T())
	if err != nil {
		return nil, fmt.Errorf("Think: error occurred: %v", err)
	}
	return sigmoidM(dot), nil
}
