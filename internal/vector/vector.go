package vector

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Vec is a n-dimensional vector
type Vec []float64

// Make creates an instance of Vec
func Make(x ...float64) Vec {
	return x
}

// Zero creates and returns a zero vector of size n
func Zero(n int) Vec {
	return make([]float64, n, n)
}

func (p Vec) String() string {
	pieces := []string{}
	for _, v := range p {
		pieces = append(pieces, strconv.FormatFloat(v, 'f', 3, 64))
	}
	return fmt.Sprintf("[%v]", strings.Join(pieces, ", "))
	// return fmt.Sprintf("<%vx1 Vec>[%v]", len(p), strings.Join(pieces, ", "))
}

// Add adds a vector to self
func (p *Vec) Add(o Vec) error {
	if len(*p) != len(o) {
		return errors.New("invalid operation: vectors of different size")
	}
	for i, v := range *p {
		(*p)[i] = v + o[i]
	}
	return nil
}

// Sum returns the sum of this vector with another
func (p Vec) Sum(o Vec) (Vec, error) {
	if len(p) != len(o) {
		return nil, errors.New("invalid operation: vectors of different size")
	}
	res := make([]float64, len(p))
	for i, v := range p {
		res[i] = v + o[i]
	}
	return res, nil
}

// Dot returns dot product of this vector with other vector.
func (p Vec) Dot(o Vec) (float64, error) {
	if len(p) != len(o) {
		return 0, errors.New("invalid operation: vectors of different size")
	}
	var sum float64 = 0.0
	for i, v := range p {
		sum += v * o[i]
	}
	return sum, nil
}

// Scale scales a vector.
func (p *Vec) Scale(f float64) {
	for i, v := range *p {
		(*p)[i] = v * f
	}
}

// Mult returns a scaled copy of this vector
func (p Vec) Mult(f float64) Vec {
	res := make([]float64, len(p))
	for i, v := range p {
		res[i] = v * f
	}
	return res
}

// Equal tells whether a and b contain the same elements.
// A nil argument is equivalent to an empty slice.
func (p Vec) Equal(b Vec) bool {
	if len(p) != len(b) {
		return false
	}
	for i, v := range p {
		if v != b[i] {
			return false
		}
	}
	return true
}
