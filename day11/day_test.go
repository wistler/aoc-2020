package day11

import (
	"testing"

	"github.com/wistler/aoc-2020/internal/io"
)

func TestAdj(t *testing.T) {
	seatMap := [][]string{
		{"L", "L", "L"},
		{"L", "L", "L"},
		{"L", "L", "L"},
	}
	testCases := []struct {
		row      int
		col      int
		expected int
	}{
		{row: 1, col: 1, expected: 0},
		{row: 0, col: 0, expected: 0},
		{row: 1, col: 2, expected: 0},
		{row: 2, col: 0, expected: 0},
		{row: 2, col: 2, expected: 0},
	}
	for _, tC := range testCases {
		got := adjOccupied(seatMap, tC.row, tC.col, false)
		if got != tC.expected {
			t.Fatalf("Wanted %v, but got %v", tC.expected, got)
		}
	}

	seatMap = [][]string{
		{"L", "#", "L"},
		{"L", "L", "L"},
		{"L", "L", "L"},
	}
	testCases = []struct {
		row      int
		col      int
		expected int
	}{
		{row: 1, col: 1, expected: 1},
		{row: 0, col: 0, expected: 1},
		{row: 0, col: 2, expected: 1},
		{row: 1, col: 0, expected: 1},
		{row: 1, col: 2, expected: 1},
		{row: 2, col: 0, expected: 0},
		{row: 2, col: 2, expected: 0},
	}
	for _, tC := range testCases {
		got := adjOccupied(seatMap, tC.row, tC.col, false)
		if got != tC.expected {
			t.Fatalf("Wanted %v, but got %v", tC.expected, got)
		}
	}
}

func TestMoreAdj(t *testing.T) {
	testCases := []struct {
		seatMap  [][]string
		row      int
		col      int
		expected int
	}{
		{
			seatMap: toSeatMap([]string{
				".......#.",
				"...#.....",
				".#.......",
				".........",
				"..#L....#",
				"....#....",
				".........",
				"#........",
				"...#.....",
			}),
			row:      4,
			col:      3,
			expected: 8,
		},
		{
			seatMap: toSeatMap([]string{
				".............",
				".L.L.#.#.#.#.",
				".............",
			}),
			row:      1,
			col:      1,
			expected: 0,
		},
		{
			seatMap: toSeatMap([]string{
				".##.##.",
				"#.#.#.#",
				"##...##",
				"...L...",
				"##...##",
				"#.#.#.#",
				".##.##.",
			}),
			row:      3,
			col:      3,
			expected: 0,
		},
	}
	for _, tC := range testCases {
		got := adjOccupied(tC.seatMap, tC.row, tC.col, true)
		if got != tC.expected {
			t.Fatalf("SeatMap: %v\nWanted %v, but got %v", tC.seatMap, tC.expected, got)
		}
	}
}

func TestSampleData(t *testing.T) {
	input := []string{
		"L.LL.LL.LL",
		"LLLLLLL.LL",
		"L.L.L..L..",
		"LLLL.LL.LL",
		"L.LL.LL.LL",
		"L.LLLLL.LL",
		"..L.L.....",
		"LLLLLLLLLL",
		"L.LLLLLL.L",
		"L.LLLLL.LL",
	}

	got := part1(input)
	want := 37
	if got != want {
		t.Fatalf("Part 1: Got: %v, but wanted: %v", got, want)
	}

	got = part2(input)
	want = 26
	if got != want {
		t.Fatalf("Part 2: Got: %v, but wanted: %v", got, want)
	}
}

func TestWithRealData(t *testing.T) {
	input := io.ReadInputFile("./input.txt")

	part1(input)
	part2(input)
}
