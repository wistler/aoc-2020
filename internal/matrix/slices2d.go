package matrix

import "github.com/wistler/aoc-2020/internal/vector"

// Count returns number of occurances in matrix
func Count(matrix [][]string, value string) int {
	R := len(matrix)
	C := len(matrix[0])
	count := 0
	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			if matrix[r][c] == value {
				count++
			}
		}
	}
	return count
}

// Equal check for matrix equality
func Equal(matA [][]string, matB [][]string) bool {
	if len(matA) != len(matB) {
		return false
	}
	R := len(matA)
	if R == 0 {
		return true
	}
	if len(matA[0]) != len(matB[0]) {
		return false
	}
	C := len(matA[0])
	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			if matA[r][c] != matB[r][c] {
				return false
			}
		}
	}
	return true
}

// Get returns value of matrix at index
func Get(matrix [][]string, index vector.Vec, def string) string {
	x := int(index[0])
	y := int(index[1])
	if x >= 0 && x < len(matrix) &&
		y >= 0 && y < len(matrix[0]) {
		return matrix[x][y]
	}
	return def
}
