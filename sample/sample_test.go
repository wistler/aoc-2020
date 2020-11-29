package main

import (
	"testing"

	"github.com/wistler/aoc-2020/internal/vector"
)

func TestSum(t *testing.T) {
	// given
	pts := []vector.Vec{
		vector.Make(1, 2),
		vector.Make(3, 4),
	}
	want := vector.Make(4, 6)

	// when
	sum := Sum(pts)

	// then
	if !sum.Equal(want) {
		t.Fatalf(`Sum(%q) = %q, want match for %q`, pts, sum, want)
	}
}
