package matrix

import "github.com/wistler/aoc-2020/internal/vector"

// String is a matrix where with string elements
type String [][]string

// Count returns number of occurances in matrix
func (mat String) Count(value string) int {
	R := len(mat)
	C := len(mat[0])
	count := 0
	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			if mat[r][c] == value {
				count++
			}
		}
	}
	return count
}

// Equal check for matrix equality
func (mat String) Equal(matB String) bool {
	A, B := mat, matB
	if len(A) != len(B) {
		return false
	}
	R := len(A)
	if R == 0 {
		return true
	}
	if len(A[0]) != len(B[0]) {
		return false
	}
	C := len(A[0])
	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			if A[r][c] != B[r][c] {
				return false
			}
		}
	}
	return true
}

// StringMatrix creates a new string matrix
func StringMatrix(rows, cols int) String {
	mat := make([][]string, rows)
	for r := 0; r < rows; r++ {
		mat[r] = make([]string, cols)
		for c := 0; c < cols; c++ {
			mat[r][c] = ""
		}
	}
	return mat
}

// Shape ...
func (mat String) Shape() (int, int) {
	return len(mat), len(mat[0])
}

// Get returns value of matrix at index
func (mat String) Get(index vector.Vec, def string) string {
	x := int(index[0])
	y := int(index[1])
	if x >= 0 && x < len(mat) &&
		y >= 0 && y < len(mat[0]) {
		return mat[x][y]
	}
	return def
}
