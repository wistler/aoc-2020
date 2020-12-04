package day01

import (
	"io/ioutil"
	"strconv"
	"strings"
	"testing"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func TestWithSampleData(t *testing.T) {
	input := []int{
		1721,
		979,
		366,
		299,
		675,
		1456,
	}
	part1Ans := 514579
	part2Ans := 241861950

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
	data, err := ioutil.ReadFile("./input.txt")
	check(err)
	tmp := strings.Split(string(data), "\r\n")

	input := make([]int, len(tmp))
	for i, t := range tmp {
		input[i], err = strconv.Atoi(t)
		check(err)
	}

	_, err = part1(input)
	check(err)

	_, err = part2(input)
	check(err)
}
