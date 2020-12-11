package day11

import (
	"log"
	"strings"

	"github.com/wistler/aoc-2020/internal/matrix"
	"github.com/wistler/aoc-2020/internal/vector"
)

const (
	floor      = "."
	unoccupied = "L"
	occupied   = "#"
)

var dirs []vector.Vec = []vector.Vec{
	vector.Make(-1, -1),
	vector.Make(-1, 0),
	vector.Make(-1, +1),
	vector.Make(0, -1),
	vector.Make(0, +1),
	vector.Make(+1, -1),
	vector.Make(+1, 0),
	vector.Make(+1, +1),
}

func isOccupied(seatMap [][]string, center vector.Vec, dir vector.Vec, firstVisible bool) bool {
	seat := floor
	pos := center
	for {
		var err error
		pos, err = pos.Sum(dir)
		if err != nil {
			panic(err)
		}
		seat = matrix.Get(seatMap, pos, "x")
		if !firstVisible {
			// only look at the vert first position in given direction
			break
		}
		// log.Printf("Pos %v, Seat: %q, Dir: %v", pos, seat, dir)
		if seat != floor {
			break
		}
	}
	return seat == occupied
}

func adjOccupied(seatMap [][]string, r, c int, firstVisible bool) int {
	center := vector.Make(float64(r), float64(c))
	adjOccupied := 0
	for _, dir := range dirs {
		if isOccupied(seatMap, center, dir, firstVisible) {
			adjOccupied++
		}
	}
	return adjOccupied
}

func sim(seatMap [][]string, occupancyLimit int, firstVisible bool) [][]string {
	R := len(seatMap)
	C := len(seatMap[0])
	nextSeatMap := make([][]string, R)
	for r := 0; r < R; r++ {
		nextSeatMap[r] = make([]string, C)

		for c := 0; c < C; c++ {
			nextSeatMap[r][c] = seatMap[r][c]
			if seatMap[r][c] == floor {
				continue
			}
			adj := adjOccupied(seatMap, r, c, firstVisible)
			if adj >= occupancyLimit && seatMap[r][c] == occupied {
				nextSeatMap[r][c] = unoccupied
			}
			if adj == 0 && seatMap[r][c] == unoccupied {
				nextSeatMap[r][c] = occupied
			}
		}
	}
	return nextSeatMap
}

func toSeatMap(input []string) [][]string {
	seatMap := make([][]string, len(input))
	for i, row := range input {
		seatMap[i] = strings.Split(row, "")
	}
	return seatMap
}

func part1(input []string) int {
	log.SetPrefix("Day 11: Part 1: ")
	log.SetFlags(0)

	seatMap := toSeatMap(input)
	steps := 0
	for ; ; steps++ {
		newMap := sim(seatMap, 4, false)
		if matrix.Equal(newMap, seatMap) {
			break
		}
		seatMap = newMap
	}

	sum := matrix.Count(seatMap, occupied)

	log.Printf("At step: %v\n", steps)
	log.Printf("Answer: %v", sum)
	return sum
}

func part2(input []string) int {
	log.SetPrefix("Day 11: Part 2: ")
	log.SetFlags(0)

	seatMap := toSeatMap(input)
	steps := 0
	for ; ; steps++ {
		newMap := sim(seatMap, 5, true)
		if matrix.Equal(newMap, seatMap) {
			break
		}
		seatMap = newMap
	}

	log.Printf("At step: %v\n", steps)
	sum := matrix.Count(seatMap, occupied)

	log.Printf("Answer: %v", sum)
	return sum
}
