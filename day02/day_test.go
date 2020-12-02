package day02

import (
	"io/ioutil"
	"strings"
	"testing"
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

func TestPart1(t *testing.T) {
	data, err := ioutil.ReadFile("./input.txt")
	check(err)
	input := strings.Split(string(data), "\r\n")

	_, err = part1(input)
	check(err)

	_, err = part2(input)
	check(err)
}
