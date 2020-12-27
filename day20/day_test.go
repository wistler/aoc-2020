package day20

import (
	"log"
	"os"
	"testing"

	"github.com/wistler/aoc-2020/internal/io"
)

func TestSyncState(t *testing.T) {
	testCases := []struct {
		src           [4]int
		target        [4]int
		state         tileState
		skipApplyTest bool
	}{
		{
			src:    [4]int{1, 2, 3, 4},
			target: [4]int{1, 2, 3, 4},
			state:  tileState{0, false},
		},
		{
			src:    [4]int{1, 2, 3, 4},
			target: [4]int{2, 3, 4, 1},
			state:  tileState{1, false},
		},
		{
			src:    [4]int{1, 2, 3, 4},
			target: [4]int{3, 4, 1, 2},
			state:  tileState{2, false},
		},
		{
			src:    [4]int{1, 2, 3, 4},
			target: [4]int{4, 1, 2, 3},
			state:  tileState{3, false},
		},
		{
			src:    [4]int{1, 2, 3, 4},
			target: [4]int{3, 2, 1, 4},
			state:  tileState{3, true},
		},
		{
			src:    [4]int{1, 2, 3, 4},
			target: [4]int{2, 1, 4, 3},
			state:  tileState{2, true},
		},
		{
			src:    [4]int{1, 2, 3, 4},
			target: [4]int{1, 4, 3, 2},
			state:  tileState{1, true},
		},
		{
			src:    [4]int{1, 2, 3, 4},
			target: [4]int{4, 3, 2, 1},
			state:  tileState{0, true},
		},
		{
			src:           [4]int{1, 0, 0, 4},
			target:        [4]int{0, -1, -1, 0},
			state:         tileState{2, false},
			skipApplyTest: true,
		},
		{
			src:           [4]int{0, 0, 0, 4},
			target:        [4]int{4, -1, -1, -1},
			state:         tileState{3, false},
			skipApplyTest: true,
		},
	}

	for _, tC := range testCases {
		state, err := findSyncState(tC.src, tC.target)
		if err != nil {
			t.Log("src:", tC.src, "target:", tC.target)
			findSyncState(tC.src, tC.target)
			panic(err)
		}
		if tC.state != state {
			t.Log("src:", tC.src, "target:", tC.target)
			findSyncState(tC.src, tC.target)
			t.Fatalf("Wanted %v, got %v", tC.state, state)
		}
		if !tC.skipApplyTest {
			target := applyState(tC.src, tC.state)
			if target != tC.target {
				t.Fatalf("applyState(%v): Wanted %v, got %v", tC.state, tC.target, target)
			}
		}
	}
}

func TestZip(t *testing.T) {
	rawInput := `
Tile 0001:
..##.#..#.
##..#.....
#...##..#.
####.#...#
##.##.###.
##...#.###
.#.#.#..##
..#....#..
###...#.#.
..###..###

Tile 0002:
#.##...##.
#.####...#
.....#..##
#...######
.##.#....#
.###.#####
###.##.##.
.###....#.
..#.#..#.#
#...##.#..

Tile 0003:
####...##.
#..##.#..#
##.#..#.#.
.###.####.
..###.####
.##....##.
.#...####.
#.##.####.
####..#...
.....##...

`
	input := io.SplitOnNewLines(rawInput)
	tiles := parseInput(input)

	ids := []image{tiles[1].rawData, tiles[2].rawData, tiles[3].rawData}
	z := zip(ids...)
	t.Log(len(z), len(z[0]))
	for _, zr := range z {
		t.Log(toString(zr))
	}
}

func TestTransform(t *testing.T) {
	testCases := []struct {
		initial string
		state   []tileState
		final   string
	}{
		{
			initial: `
			..........
			.#####....
			..........
			..........
			..........
			..........
			..........
			..........
			..........
			..........
			`,
			state: []tileState{{1, false}},
			final: `
			..........
			..........
			..........
			..........
			.#........
			.#........
			.#........
			.#........
			.#........
			..........
			`,
		},
		{
			initial: `
			..........
			.#####....
			..........
			..........
			..........
			..........
			..........
			..........
			..........
			..........
			`,
			state: []tileState{{3, false}},
			final: `
			..........
			........#.
			........#.
			........#.
			........#.
			........#.
			..........
			..........
			..........
			..........
			`,
		},
		{
			initial: `
			..........
			.#####....
			..........
			..........
			..........
			..........
			..........
			..........
			..........
			..........
			`,
			state: []tileState{{0, true}},
			final: `
			..........
			.#........
			.#........
			.#........
			.#........
			.#........
			..........
			..........
			..........
			..........
			`,
		},
		{
			initial: `
			..........
			.#####....
			..........
			..........
			..........
			..........
			..........
			..........
			..........
			..........
			`,
			state: []tileState{{1, true}},
			final: `
			..........
			....#####.
			..........
			..........
			..........
			..........
			..........
			..........
			..........
			..........
			`,
		},
		{
			initial: `
			..........
			.#####....
			..........
			..........
			..........
			..........
			..........
			..........
			..........
			..........
			`,
			state: []tileState{{2, true}},
			final: `
			..........
			..........
			..........
			..........
			........#.
			........#.
			........#.
			........#.
			........#.
			..........
			`,
		},
		{
			initial: `
			..........
			.#####....
			..........
			..........
			..........
			..........
			..........
			..........
			..........
			..........
			`,
			state: []tileState{{2, false}, {0, true}},
			final: `
			..........
			..........
			..........
			..........
			........#.
			........#.
			........#.
			........#.
			........#.
			..........
			`,
		},
		{
			initial: `
			..........
			.#####....
			..........
			..........
			..........
			..........
			..........
			..........
			..........
			..........
			`,
			state: []tileState{{1, false}, {0, true}, {1, false}},
			final: `
			..........
			..........
			..........
			..........
			........#.
			........#.
			........#.
			........#.
			........#.
			..........
			`,
		},
	}
	for _, tC := range testCases {
		initial := stringsToBool(tC.initial)
		final := stringsToBool(tC.final)

		t0 := tile{}
		t0.rawData = initial
		results := []image{}
		results = append(results, initial)
		for _, st := range tC.state {
			t0 = t0.transform(st)
			results = append(results, t0.rawData)
		}
		results = append(results, final)

		for i := 0; i < len(t0.rawData); i++ {
			if !pixelsMatch(t0.rawData[i], final[i]) {
				t1 := tile{}
				t1.rawData = initial
				for _, st := range tC.state {
					display(t1.rawData)
					t2 := t1.transform(st)
					display(t2.rawData)
					t1 = t2
				}

				display(results...)
				t.Fatalf("Match Error for transform: %v", tC.state)
			}
		}
	}
}

func TestMonsterHunt(t *testing.T) {
	testCases := []struct {
		haystack string
		needle   string
		r        int
		m        int
		err      bool
	}{
		{
			haystack: `
			........................
			........................
			........................
			........................
			........................
				`,
			needle: `
			...#.
			.###.
			`,
			r:   0,
			m:   0,
			err: true,
		},
		{
			haystack: `
			........................
			........#...............
			......###...............
			........................
			........................
				`,
			needle: `
			...#.
			.###.
			`,
			r:   0,
			m:   1,
			err: false,
		},
		{
			haystack: `
			........................
			........................
			......###...............
			........#...............
			........................
				`,
			needle: `
			...#.
			.###.
			`,
			r:   0,
			m:   0,
			err: true,
		},
		{
			haystack: `
			........................
			........#...............
			......###.......#.......
			..............###.......
			........................
				`,
			needle: `
			...#.
			.###.
			`,
			r:   0,
			m:   2,
			err: false,
		},
		{
			haystack: `
			............#...........
			.....##.#......#........
			......###..#....##.#....
			............#.###.......
			......#.................
				`,
			needle: `
			...#.
			.###.
			`,
			r:   9,
			m:   2,
			err: false,
		},
	}

	log.SetOutput(os.Stdout)
	for _, tC := range testCases {
		image := stringsToBool(tC.haystack)
		duck := stringsToBool(tC.needle)
		r, m, w, err := findRoughness(image, duck)

		if tC.err && err == nil {
			t.Log("r:", r, "m:", m, "err:", err)
			display(w)

			t.Fatal("Expected error, but got not error.")
		}
		if tC.r != r {
			t.Log("r:", r, "m:", m, "err:", err)
			display(w)

			t.Fatalf("Expected r = %v, but got r = %v", tC.r, r)
		}
		if tC.m != m {
			t.Log("r:", r, "m:", m, "err:", err)
			t.Log(w)

			t.Fatalf("Expected m = %v, but got m = %v", tC.m, m)
		}
	}
}

func TestMonsterHunt2(t *testing.T) {
	testCases := []struct {
		haystack string
		needle   string
		r        int
		m        int
	}{
		{
			haystack: `
			..........
			..........
			..........
			..........
			..........
			..........
			..........
			..........
			..........
			..........
				`,
			needle: `
			...#.
			.###.
			`,
			r: 0,
			m: 0,
		},
		{
			haystack: `
			..........
			........#.
			......###.
			..........
			..........
			..........
			..........
			..........
			..........
			..........
				`,
			needle: `
			...#.
			.###.
			`,
			r: 0,
			m: 1,
		},
		{
			haystack: `
			..........
			..........
			......###.
			........#.
			..........
			..........
			..........
			..........
			..........
			..........
				`,
			needle: `
			...#.
			.###.
			`,
			r: 0,
			m: 1,
		},
		{
			haystack: `
			..........
			........#.
			......###.
			..........
			..........
			..........
			..........
			......#...
			....###...
			..........
				`,
			needle: `
			...#.
			.###.
			`,
			r: 0,
			m: 2,
		},
		{
			haystack: `
			.........#
			..#.##....
			...###..#.
			.........#
			...#......
			..........
			..#.......
			...##.#...
			.###......
			..........
				`,
			needle: `
			...#.
			.###.
			`,
			r: 9,
			m: 2,
		},
	}

	for _, tC := range testCases {
		t.Log("-------- tC --------")
		image := stringsToBool(tC.haystack)
		duck := stringsToBool(tC.needle)
		r, m := findRoughnessAndMosters(image, duck)

		if tC.r != r {
			t.Log("r:", r, "m:", m)
			t.Fatalf("Expected r = %v, but got r = %v", tC.r, r)
		}
		if tC.m != m {
			t.Log("r:", r, "m:", m)
			t.Fatalf("Expected m = %v, but got m = %v", tC.m, m)
		}
	}

}

func TestReadData(t *testing.T) {
	rawInput := `
Tile 0001:
..#.......
.######...
..#....#..
........#.
........#.
........#.
.......#.#
..........
#.........
....#.....

`
	input := io.SplitOnNewLines(rawInput)
	tiles := parseInput(input)

	t1 := tiles[1]
	t1.links[0] = 0
	t1.links[1] = 1
	t1.links[2] = 2
	t1.links[3] = 3
	t.Log(t1.String())

	for _, st := range states {
		t2 := t1.transform(st)
		t.Log(t2.String())
		t.Log("-----------")
	}

}

func TestWithSampleData(t *testing.T) {
	rawInput := `
Tile 2311:
..##.#..#.
##..#.....
#...##..#.
####.#...#
##.##.###.
##...#.###
.#.#.#..##
..#....#..
###...#.#.
..###..###

Tile 1951:
#.##...##.
#.####...#
.....#..##
#...######
.##.#....#
.###.#####
###.##.##.
.###....#.
..#.#..#.#
#...##.#..

Tile 1171:
####...##.
#..##.#..#
##.#..#.#.
.###.####.
..###.####
.##....##.
.#...####.
#.##.####.
####..#...
.....##...

Tile 1427:
###.##.#..
.#..#.##..
.#.##.#..#
#.#.#.##.#
....#...##
...##..##.
...#.#####
.#.####.#.
..#..###.#
..##.#..#.

Tile 1489:
##.#.#....
..##...#..
.##..##...
..#...#...
#####...#.
#..#.#.#.#
...#.#.#..
##.#...##.
..##.##.##
###.##.#..

Tile 2473:
#....####.
#..#.##...
#.##..#...
######.#.#
.#...#.#.#
.#########
.###.#..#.
########.#
##...##.#.
..###.#.#.

Tile 2971:
..#.#....#
#...###...
#.#.###...
##.##..#..
.#####..##
.#..####.#
#..#.#..#.
..####.###
..#.#.###.
...#.#.#.#

Tile 2729:
...#.#.#.#
####.#....
..#.#.....
....#..#.#
.##..##.#.
.#.####...
####.#.#..
##.####...
##..#.##..
#.##...##.

Tile 3079:
#.#.#####.
.#..######
..#.......
######....
####.#..#.
.#...#.##.
#.#####.##
..#.###...
..#.......
..#.###...

`
	input := io.SplitOnNewLines(rawInput)
	part1Ans := 20899048083289
	part2Ans := 273

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
