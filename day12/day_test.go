package day12

import (
	"testing"

	"github.com/wistler/aoc-2020/internal/io"
)

func TestRotate(t *testing.T) {
	switch true {
	case rotate('E', 1) != 'S':
		t.Fatalf("E+1 not S")
	case rotate('E', -1) != 'N':
		t.Fatalf("E+1 not N")
	case rotate('E', -3) != 'S':
		t.Fatalf("E-3 not S")
	}
}

func TestSampleData(t *testing.T) {
	input := []string{
		"F10",
		"N3",
		"F7",
		"R90",
		"F11",
	}

	got := part1(input)
	want := 25
	if got != want {
		t.Fatalf("Part 1: Got: %v, but wanted: %v", got, want)
	}

	got = part2(input)
	want = 286
	if got != want {
		t.Fatalf("Part 2: Got: %v, but wanted: %v", got, want)
	}
}

func TestWithRealData(t *testing.T) {
	input := io.ReadInputFile("./input.txt")

	part1(input)
	part2(input)
}
