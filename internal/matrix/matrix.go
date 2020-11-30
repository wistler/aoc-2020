package matrix

import (
	"errors"
	"fmt"
	"strings"

	"github.com/wistler/aoc-2020/internal/vector"
)

// Matrix is a 2D vector
type Matrix []vector.Vec

func (p Matrix) String() string {
	pieces := []string{}
	for _, v := range p {
		pieces = append(pieces, v.String())
	}
	return fmt.Sprintf("<%vx%v Mat>[%v]", len(p), len(p[0]), strings.Join(pieces, ";"))
}

// New returns a matrix instance
func New(v ...vector.Vec) Matrix {
	return v
}

// Zero returns a zero matrix of size MxN
func Zero(M, N int) Matrix {
	var mat Matrix = make([]vector.Vec, M)
	for i := 0; i < M; i++ {
		mat[i] = vector.Zero(N)
	}
	return mat
}

// AsVec returns a Vector from a Row/Column Matrix
func (p Matrix) AsVec() (vector.Vec, error) {
	if len(p) != 1 && len(p[0]) != 1 {
		return nil, errors.New("Not a 1-dimensional matrix")
	}
	if len(p) == 1 {
		return p[0], nil
	}
	return p.T()[0], nil
}

// Sum performs element-wise addition and returns new matrix
func (p Matrix) Sum(o Matrix) (Matrix, error) {
	if !p.SameShape(o) {
		return nil, fmt.Errorf("invalid operation: matrices size incompatible: %v.%v", p.ShapeString(), o.ShapeString())
	}
	M, N := p.Shape()
	sum := Zero(M, N)
	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			sum[i][j] = p[i][j] + o[i][j]
		}
	}
	return sum, nil
}

// Dot returns matrix multiplication of 2 matrices
// [MxN] * [N*P] => [MxP]
func (p Matrix) Dot(o Matrix) (Matrix, error) {
	M, N := p.Shape()
	O, P := o.Shape()
	if N != O {
		return nil, fmt.Errorf("invalid operation: matrices size incompatible: %v.%v", p.ShapeString(), o.ShapeString())
	}
	var prod Matrix = make([]vector.Vec, M)
	for i := 0; i < M; i++ {
		prod[i] = vector.Zero(P)
	}
	for i := 0; i < M; i++ {
		for j := 0; j < P; j++ {
			for k := 0; k < N; k++ {
				prod[i][j] += p[i][k] * o[k][j]
			}
		}
	}
	return prod, nil
}

// Mult returns matrix multiplied by scalar
func (p Matrix) Mult(f float64) Matrix {
	prod, _ := p.Prod(New([]float64{f}))
	return prod
}

// Prod returns element-wise product of 2 matrices
// [MxN] * [MxN] => [MxN]  // case 1
// [MxN] * [1x1] => [MxN]  // case 2
// [1x1] * [1x1] => [1x1]  // case 3 - same as case 1
func (p Matrix) Prod(o Matrix) (Matrix, error) {
	M, N := p.Shape()
	O, P := o.Shape()
	if M == O && N == P {
		var prod Matrix = make([]vector.Vec, M)
		for i := 0; i < M; i++ {
			prod[i] = vector.Zero(N)
		}
		for i := 0; i < M; i++ {
			for j := 0; j < N; j++ {
				prod[i][j] = p[i][j] * o[i][j]
			}
		}
		return prod, nil
	}
	if (M*N == 1) || (O*P == 1) {
		scalar := o[0][0]
		mat := p
		if M*N == 1 {
			scalar = p[0][0]
			mat = o
		}
		prod := make([]vector.Vec, len(mat))
		for i := 0; i < len(mat); i++ {
			prod[i] = vector.Zero(len(mat[0]))
		}
		for i := 0; i < len(mat); i++ {
			for j := 0; j < len(mat[0]); j++ {
				prod[i][j] = mat[i][j] * scalar
			}
		}
		return prod, nil
	}
	return nil, errors.New("invalid operation: matrices size incompatible")
}

// T returns the transformed matrix [M x N] => [N x M]
func (p Matrix) T() Matrix {
	var mat Matrix = make([]vector.Vec, len(p[0]))
	for i := 0; i < len(p[0]); i++ {
		mat[i] = vector.Zero(len(p))
	}
	for i := 0; i < len(p); i++ {
		for j := 0; j < len(p[0]); j++ {
			mat[j][i] = p[i][j]
		}
	}
	return mat
}

// Equal tells whether a and b contain the same elements.
// A nil argument is equivalent to an empty slice.
func (p Matrix) Equal(b Matrix) bool {
	if !p.SameShape(b) {
		return false
	}
	for i, v := range p {
		if !v.Equal(b[i]) {
			return false
		}
	}
	return true
}

// Shape returns the shape of the matrix as a [2x1] vector
func (p Matrix) Shape() (int, int) {
	return len(p), len(p[0])
}

// ShapeString returns a string
func (p Matrix) ShapeString() string {
	M, N := p.Shape()
	return fmt.Sprintf("[%vx%v]", M, N)
}

// SameShape returns true if matrics are of same shape
func (p Matrix) SameShape(o Matrix) bool {
	M, N := p.Shape()
	O, P := o.Shape()
	return M == O && N == P
}
