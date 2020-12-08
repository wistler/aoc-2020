package day00

import (
	"testing"

	"github.com/wistler/aoc-2020/internal/io"
)

func TestSampleData(t *testing.T) {
	input := []string{
		"",
	}

	got := part1(input)
	want := 0
	if got != want {
		t.Fatalf("Part 1: Got: %v, but wanted: %v", got, want)
	}

	got = part2(input)
	want = 0
	if got != want {
		t.Fatalf("Part 2: Got: %v, but wanted: %v", got, want)
	}
}

func TestWithRealData(t *testing.T) {
	input := io.ReadInputFile("./input.txt")

	part1(input)
	part2(input)
}
