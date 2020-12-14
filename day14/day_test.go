package day14

import (
	"testing"

	"github.com/wistler/aoc-2020/internal/io"
)

func TestSampleData(t *testing.T) {
	input := []string{
		"mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X",
		"mem[8] = 11",
		"mem[7] = 101",
		"mem[8] = 0",
	}

	got := part1(input)
	want := 165
	if got != want {
		t.Fatalf("Part 1: Got: %v, but wanted: %v", got, want)
	}

	input2 := []string{
		"mask = 000000000000000000000000000000X1001X",
		"mem[42] = 100",
		"mask = 00000000000000000000000000000000X0XX",
		"mem[26] = 1",
	}

	got = part2(input2)
	want = 208
	if got != want {
		t.Fatalf("Part 2: Got: %v, but wanted: %v", got, want)
	}
}

func TestWithRealData(t *testing.T) {
	input := io.ReadInputFile("./input.txt")

	part1(input)
	part2(input)
}
