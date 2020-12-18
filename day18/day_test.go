package day18

import (
	"testing"

	"github.com/wistler/aoc-2020/internal/io"
)

func TestPart1WithSampleData(t *testing.T) {
	testCases := []struct {
		line string
		ans  int
	}{
		{line: "1 + 2 * 3 + 4 * 5 + 6", ans: 71},
		{line: "1 + (2 * 3) + (4 * (5 + 6))", ans: 51},
		{line: "2 * 3 + (4 * 5)", ans: 26},
		{line: "5 + (8 * 3 + 9 + 3 * 4 * 3)", ans: 437},
		{line: "5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", ans: 12240},
		{line: "((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", ans: 13632},
	}
	for _, tC := range testCases {
		got := solve(tC.line, false, false)
		want := tC.ans
		if got != want {
			solve(tC.line, false, true)
			t.Fatalf("Part 1: Got: %v, but wanted: %v", got, want)
		}
	}
}

func TestPart2WithSampleData(t *testing.T) {
	testCases := []struct {
		line string
		ans  int
	}{
		{line: "1 + 2 * 3 + 4 * 5 + 6", ans: 231},
		{line: "1 + (2 * 3) + (4 * (5 + 6))", ans: 51},
		{line: "2 * 3 + (4 * 5)", ans: 46},
		{line: "5 + (8 * 3 + 9 + 3 * 4 * 3)", ans: 1445},
		{line: "5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", ans: 669060},
		{line: "((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", ans: 23340},
	}
	for _, tC := range testCases {
		got := solve(tC.line, true, false)
		want := tC.ans
		if got != want {
			solve(tC.line, true, true)
			t.Fatalf("Part 2: Got: %v, but wanted: %v", got, want)
		}
	}
}

func TestWithRealData(t *testing.T) {
	input := io.ReadInputFile("./input.txt")

	part1(input)
	part2(input)
}
