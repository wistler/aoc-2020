package day09

import (
	"strconv"
	"testing"

	"github.com/wistler/aoc-2020/internal/io"
)

func TestSampleData(t *testing.T) {
	input := []int{
		35,
		20,
		15,
		25,
		47,
		40,
		62,
		55,
		65,
		95,
		102,
		117,
		150,
		182,
		127,
		219,
		299,
		277,
		309,
		576,
	}

	got := part1(input, 5)
	want := 127
	if got != want {
		t.Fatalf("Part 1: Got: %v, but wanted: %v", got, want)
	}

	got = part2(input, 5)
	want = 62
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

	part1(input, 25)
	part2(input, 25)
}
