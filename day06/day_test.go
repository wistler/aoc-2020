package day06

import (
	"testing"

	"github.com/wistler/aoc-2020/internal/io"
)

func Equal(p []int, b []int) bool {
	if len(p) != len(b) {
		return false
	}
	for i, v := range p {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestSampleData(t *testing.T) {
	input := []string{
		"abc",
		"",
		"a",
		"b",
		"c",
		"",
		"ab",
		"ac",
		"",
		"a",
		"a",
		"a",
		"a",
		"",
		"b",
	}
	uniq := []int{3, 3, 3, 1, 1}
	common := []int{3, 0, 1, 1, 1}
	g := getGroups(input)
	if len(g) != 5 {
		t.Fatalf("Expected 4 groups, but got %v.\n%v", len(g), g)
	}
	for i, gr := range g {
		u := getUniqueAnswers(gr)
		if len(u) != uniq[i] {
			t.Fatalf("Expected uniq ans %v, but got %v.\n%v => %v", uniq[i], len(u), gr, u)
		}
		c := getCommonAnswers(gr)
		if len(c) != common[i] {
			t.Fatalf("Expected common ans %v, but got %v.\n%v => %v", common[i], len(c), gr, c)
		}
	}

	part1(input)
	part2(input)
}

func TestWithRealData(t *testing.T) {
	input := io.ReadInputFile("./input.txt")

	part1(input)
	part2(input)
}
