package day13

import (
	"testing"

	"github.com/wistler/aoc-2020/internal/io"
)

func TestSampleData(t *testing.T) {
	input := []string{
		"939",
		"7,13,x,x,59,x,31,19",
	}

	got := part1(input)
	want := 295
	if got != want {
		t.Fatalf("Part 1: Got: %v, but wanted: %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	testCases := []struct {
		input  string
		output int
	}{
		{input: "17,x,13,19", output: 3_417},
		{input: "67,7,59,61", output: 754_018},
		{input: "67,x,7,59,61", output: 779_210},
		{input: "67,7,x,59,61", output: 1_261_476},
		{input: "7,13,x,x,59,x,31,19", output: 1_068_781},
		{input: "1789,37,47,1889", output: 1_202_161_486},
	}
	for _, tC := range testCases {
		got := part2(tC.input, false)
		if got != tC.output {
			t.Fatalf("Part 2: Got: %v, but wanted: %v", got, tC.output)
		}
	}
}

func TestWithRealData(t *testing.T) {
	input := io.ReadInputFile("./input.txt")

	part1(input)
	part2(input[1], true)
}
