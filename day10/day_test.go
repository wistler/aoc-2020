package day10

import (
	"strconv"
	"testing"

	"github.com/wistler/aoc-2020/internal/io"
)

func TestSampleData(t *testing.T) {
	input := []int{
		16,
		10,
		15,
		5,
		1,
		11,
		7,
		19,
		6,
		12,
		4,
	}

	got := part1(input)
	want := 35
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
	tmp := io.ReadInputFile("./input.txt")
	input := make([]int, len(tmp))
	for i, t := range tmp {
		ti, err := strconv.Atoi(t)
		if err != nil {
			panic(err)
		}
		input[i] = ti
	}

	part1(input)
	part2(input)
}
