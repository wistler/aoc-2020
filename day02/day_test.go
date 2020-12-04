package day02

import (
	"testing"

	"github.com/wistler/aoc-2020/internal/io"
)

func TestWithSampleData(t *testing.T) {
	input := []string{
		"1-3 a: abcde",
		"1-3 b: cdefg",
		"2-9 c: ccccccccc",
	}
	part1Ans := 2
	part2Ans := 1

	got, err := part1(input)
	check(err)
	if got != part1Ans {
		t.Fatalf(`Part 1: got %v, but want %v`, got, part1Ans)
	}

	got, err = part2(input)
	check(err)
	if got != part2Ans {
		t.Fatalf(`Part 2: got %v, but want %v`, got, part1Ans)
	}
}

func TestWithRealData(t *testing.T) {
	input := io.ReadInputFile("./input.txt")

	_, err := part1(input)
	check(err)

	_, err = part2(input)
	check(err)
}
