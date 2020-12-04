package day03

import (
	"testing"

	"github.com/wistler/aoc-2020/internal/io"
)

func TestWithSampleData(t *testing.T) {
	input := []string{
		"..##.......",
		"#...#...#..",
		".#....#..#.",
		"..#.#...#.#",
		".#...##..#.",
		"..#.##.....",
		".#.#.#....#",
		".#........#",
		"#.##...#...",
		"#...##....#",
		".#..#...#.#",
	}
	part1Ans := 7
	part2Ans := 336

	got := part1(input)
	if got != part1Ans {
		t.Fatalf(`Part 1: got %v, but want %v`, got, part1Ans)
	}

	got = part2(input)
	if got != part2Ans {
		t.Fatalf(`Part 2: got %v, but want %v`, got, part1Ans)
	}
}

func TestWithRealData(t *testing.T) {
	input := io.ReadInputFile("./input.txt")

	part1(input)
	part2(input)
}
