package day03

import (
	"io/ioutil"
	"strings"
	"testing"
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
	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	input := strings.Split(string(data), "\r\n")

	part1(input)
	part2(input)
}
