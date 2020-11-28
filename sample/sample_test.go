package main

import (
	"testing"

	"github.com/wistler/aoc-2020/lib"
)

func TestSum(t *testing.T) {
	// given
	pts := []lib.Vec{
		lib.MakeVector(1, 2),
		lib.MakeVector(3, 4),
	}
	want := lib.MakeVector(4, 6)

	// when
	sum := Sum(pts)

	// then
	if sum != want {
		t.Fatalf(`Sum(%q) = %q, want match for %q`, pts, sum, want)
	}
}
