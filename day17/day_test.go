package day17

import (
	"testing"

	"github.com/wistler/aoc-2020/internal/io"
)

func TestSampleData(t *testing.T) {

	input := []string{
		".#.",
		"..#",
		"###",
	}

	got := part1(input, false)
	want := 112
	if got != want {
		t.Fatalf("Part 1: Got: %v, but wanted: %v", got, want)
	}

	got = part2(input, false)
	want = 848
	if got != want {
		t.Fatalf("Part 2: Got: %v, but wanted: %v", got, want)
	}
}

func TestWithRealData(t *testing.T) {
	input := io.ReadInputFile("./input.txt")

	part1(input, false)
	part2(input, false)
}
