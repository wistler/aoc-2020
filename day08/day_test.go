package day08

import (
	"testing"

	"github.com/wistler/aoc-2020/internal/io"
)

func TestSampleData(t *testing.T) {
	input := []string{
		"nop +0",
		"acc +1",
		"jmp +4",
		"acc +3",
		"jmp -3",
		"acc -99",
		"acc +1",
		"jmp -4",
		"acc +6",
	}

	got := part1(input)
	want := 5
	if got != want {
		t.Fatalf("Part 1: Got: %v, but wanted: %v", got, want)
	}

	got = part2(input)
	want = 8
	if got != want {
		t.Fatalf("Part 2: Got: %v, but wanted: %v", got, want)
	}
}

func TestWithRealData(t *testing.T) {
	input := io.ReadInputFile("./input.txt")

	part1(input)
	part2(input)
}
