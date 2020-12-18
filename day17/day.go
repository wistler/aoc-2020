package day17

import (
	"log"
)

type coord [4]int
type space map[coord]bool

var neighbors = []coord{}
var neighbors4D = []coord{}

func init() {
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			for z := -1; z <= 1; z++ {
				if !(x == 0 && y == 0 && z == 0) {
					neighbors = append(neighbors, coord{x, y, z})
				}
				for w := -1; w <= 1; w++ {
					if !(x == 0 && y == 0 && z == 0 && w == 0) {
						neighbors4D = append(neighbors4D, coord{x, y, z, w})
					}
				}
			}
		}
	}
}

func neighborsOf(o coord, use4D bool) []coord {
	nArr := neighbors
	if use4D {
		nArr = neighbors4D
	}

	result := make([]coord, len(nArr))
	for i, n := range nArr {
		nc := coord{
			o[0] + n[0],
			o[1] + n[1],
			o[2] + n[2],
			o[3] + n[3],
		}
		result[i] = nc
	}
	return result
}

func activatedNeighborsOf(currState space, o coord, use4D bool) int {
	active := 0
	for _, nc := range neighborsOf(o, use4D) {
		if currState[nc] {
			active++
		}
	}
	return active
}

func countActivated(currState space, pts []coord) int {
	active := 0
	for _, nc := range pts {
		if currState[nc] {
			active++
		}
	}
	return active
}

func simCycle(currState space, use4D bool, debug bool) space {
	newState := make(space)

	for o := range currState {
		nc := neighborsOf(o, use4D)
		activatedNeighbors := countActivated(currState, nc)
		if debug {
			log.Print("\n\no:", o, "activatedNeighbors:", activatedNeighbors)
		}

		// if key o exists, then it's active ..
		if activatedNeighbors == 2 || activatedNeighbors == 3 {
			newState[o] = true
		}

		// check neighbors because we were active in currState...
		for _, no := range nc {
			if currState[no] {
				continue // only interested in inactive neighbors
			}
			ano := activatedNeighborsOf(currState, no, use4D)
			if debug {
				log.Print("? no:", no, "activatedNeighbors:", ano)
			}

			if ano == 3 {
				newState[no] = true
			}
		}

	}

	return newState
}

func toStateSpace(input []string) space {
	myState := make(space)
	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[0]); x++ {
			if input[y][x] == '#' {
				myState[coord{x, y}] = true
			}
		}
	}
	return myState
}

func part1(input []string, debug bool) int {
	log.SetPrefix("Day 17: Part 1: ")
	log.SetFlags(0)

	myState := toStateSpace(input)
	for i := 0; i < 6; i++ {
		myState = simCycle(myState, false, debug)
	}
	sum := len(myState)

	log.Printf("Answer: %v", sum)
	return sum
}

func part2(input []string, debug bool) int {
	log.SetPrefix("Day 17: Part 2: ")
	log.SetFlags(0)

	myState := toStateSpace(input)
	for i := 0; i < 6; i++ {
		myState = simCycle(myState, true, debug)
	}
	sum := len(myState)

	log.Printf("Answer: %v", sum)
	return sum
}
