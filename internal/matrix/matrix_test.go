package matrix

import (
	"testing"

	"github.com/wistler/aoc-2020/internal/vector"
)

func TestTransform(t *testing.T) {
	// given
	m := New(
		vector.Make(1, 1),
		vector.Make(2, 1),
		vector.Make(3, 1),
	)

	// when
	t.Logf("Matrix m = %v", m)
	mT := m.T()

	// then
	t.Logf("m.T() = %v", mT)
	if len(mT) != len(m[0]) || len(mT[0]) != len(m) {
		t.Fatalf("Transform size not correct. Expected: [%vx%v], but was: [%vx%v]",
			len(m[0]), len(m),
			len(mT), len(mT[0]))
	}
	if mT[0][2] != m[2][0] {
		t.Fatalf("Element on transformed index [0][2] incorrect. Expected: %v, but was: %v", m[2][0], mT[0][2])
	}
}

func TestDot(t *testing.T) {
	// given
	m := New(
		vector.Make(1, 1),
		vector.Make(2, 1),
		vector.Make(3, 1),
	).T()

	v := New(vector.Make(1, 2, 3)).T()
	expectedDot := New(vector.Make(1+4+9, 1+2+3)).T()

	// when
	t.Logf("Matrix m = %v", m)
	t.Logf("Vector v = %v", v)
	dot, err := m.Dot(v)

	// then
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if !expectedDot.Equal(dot) {
		t.Fatalf("Expected dot: %v, but got: %v", expectedDot, dot)
	}
}

func TestProd(t *testing.T) {
	// case 1
	m := New(vector.Make(1, 2, 3))
	v := New(vector.Make(1, 2, 3))
	expectedProd := New(vector.Make(1, 4, 9))
	prod, err := m.Prod(v)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if !expectedProd.Equal(prod) {
		t.Fatalf("Expected prod: %v, but got: %v", expectedProd, prod)
	}

	// case 2
	m2 := New(vector.Make(4))
	v2 := New(vector.Make(5))
	expectedProd2 := New(vector.Make(20))
	prod2, err2 := m2.Prod(v2)
	if err2 != nil {
		t.Fatalf("Unexpected error: %v", err2)
	}
	if !expectedProd2.Equal(prod2) {
		t.Fatalf("Expected prod: %v, but got: %v", expectedProd2, prod2)
	}

	// case 3
	m3 := New(vector.Make(1, 2, 3))
	v3 := New(vector.Make(4))
	expectedProd3 := New(vector.Make(4, 8, 12))
	prod3, err3 := m3.Prod(v3)
	if err3 != nil {
		t.Fatalf("Unexpected error: %v", err3)
	}
	if !expectedProd3.Equal(prod3) {
		t.Fatalf("Expected prod: %v, but got: %v", expectedProd3, prod3)
	}
}
