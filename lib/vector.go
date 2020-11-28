package lib

import "fmt"

// Vec is a 2D vector
type Vec struct {
	x int
	y int
}

// MakeVector creates an instance of Vec
func MakeVector(x, y int) Vec {
	return Vec{x, y}
}

// ToString hello
func (p Vec) String() string {
	return fmt.Sprintf("[x: %d, y: %d]", p.x, p.y)
}

// Add adds another
func (p Vec) Add(o Vec) Vec {
	return Vec{
		x: p.x + o.x,
		y: p.y + o.y,
	}
}

// Scale scales a vector.
func (p Vec) Scale(f int) Vec {
	return Vec{
		x: p.x * f,
		y: p.y * f,
	}
}
