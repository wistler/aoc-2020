package vector

import "testing"

func TestSum(t *testing.T) {
	v1 := Make(1, 2, 3, 4)
	v2 := Make(5, 6, 7, 0)

	v3, _ := v1.Sum(v2)
	v3e := Make(6, 8, 10, 4)
	if !v3e.Equal(v3) {
		t.Fatalf("Expected Sum: %v, but got: %v", v3e, v3)
	}
}

func TestMult(t *testing.T) {
	v1 := Make(1, 2, 3, 4)

	v3 := v1.Mult(-1)
	v3e := Make(-1, -2, -3, -4)
	if !v3e.Equal(v3) {
		t.Fatalf("Expected Sum: %v, but got: %v", v3e, v3)
	}
}

func TestDiff(t *testing.T) {
	v1 := Make(1, 2, 3, 4)
	v2 := Make(5, 6, 7, 0)

	v3, _ := v1.Sum(v2.Mult(-1))
	v3e := Make(-4, -4, -4, 4)
	if !v3e.Equal(v3) {
		t.Fatalf("Expected Sum: %v, but got: %v", v3e, v3)
	}
}
