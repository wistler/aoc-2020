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

// // DotV returns dot product of each row of this matrix with given vector.
// func (p Matrix) DotV(o vector.Vec) (vector.Vec, error) {
// 	if len(p[0]) != len(o) {
// 		return nil, errors.New("invalid operation: matrix and vector of different size")
// 	}
// 	sum := vector.Zero(len(p))
// 	for i, v := range p {
// 		s, err := v.Dot(o)
// 		if err != nil {
// 			return nil, err
// 		}
// 		sum[i] = s
// 	}
// 	return sum, nil
// }

// Dot returns matrix multiplication of 2 matrices
// [MxN] * [N*P] => [MxP]
func (p Matrix) Dot(o Matrix) (Matrix, error) {
	M := len(p)
	N := len(p[0])
	O := len(o)
	P := len(o[0])
	if N != O {
		return nil, fmt.Errorf("invalid operation: matrices size incompatible: [%vx%v].[%vx%v]", M, N, O, P)
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

// Prod returns element-wise product of 2 matrices
// [MxN] * [MxN] => [MxN]  // case 1
// [MxN] * [1x1] => [MxN]  // case 2
// [1x1] * [1x1] => [1x1]  // case 3 - same as case 1
func (p Matrix) Prod(o Matrix) (Matrix, error) {
	M := len(p)
	N := len(p[0])
	O := len(o)
	P := len(o[0])
	if (M == O && N == P) || (M*N == 1 && O*P == 1) {
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
		// } else if M*N == 1 && O*P == 1 {
		// 	return New(vector.Make(p[0][0] * o[0][0])), nil
	} else {
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
	}
	return nil, errors.New("invalid operation: matrices size incompatible")
}

// T returns the transformed matrix
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
	if len(p) != len(b) || len(p[0]) != len(b[0]) {
		return false
	}
	for i, v := range p {
		if !v.Equal(b[i]) {
			return false
		}
	}
	return true
}
