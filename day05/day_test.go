package day05

import (
	"testing"

	"github.com/wistler/aoc-2020/internal/io"
)

func Equal(p []int, b []int) bool {
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

func TestGetHalf(t *testing.T) {
	h := getHalf([]int{0, 127}, true)
	e := []int{0, 63}
	if !Equal(h, e) {
		t.Fatalf("Expected: %v, Got: %v", e, h)
	}

	h = getHalf([]int{0, 127}, false)
	e = []int{64, 127}
	if !Equal(h, e) {
		t.Fatalf("Expected: %v, Got: %v", e, h)
	}

	h = getHalf([]int{64, 127}, true)
	e = []int{64, 95}
	if !Equal(h, e) {
		t.Fatalf("Expected: %v, Got: %v", e, h)
	}

	h = getHalf([]int{0, 63}, false)
	e = []int{32, 63}
	if !Equal(h, e) {
		t.Fatalf("Expected: %v, Got: %v", e, h)
	}

	h = getHalf([]int{0, 63}, true)
	e = []int{0, 31}
	if !Equal(h, e) {
		t.Fatalf("Expected: %v, Got: %v", e, h)
	}
}

func TestGetRow(t *testing.T) {
	r := "BFFFBBFRRR"
	R := getRow(r)
	e := 70
	if R != e {
		t.Fatalf("Row: %v, Expected: %v, Got: %v", r, e, R)
	}

	r = "FFFBBBFRRR"
	R = getRow(r)
	e = 14
	if R != e {
		t.Fatalf("Row: %v, Expected: %v, Got: %v", r, e, R)
	}

	r = "BBFFBBFRLL"
	R = getRow(r)
	e = 102
	if R != e {
		t.Fatalf("Row: %v, Expected: %v, Got: %v", r, e, R)
	}
}

func TestWithRealData(t *testing.T) {
	input := io.ReadInputFile("./input.txt")

	part1(input)
	part2(input)
}
