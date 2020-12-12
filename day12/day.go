package day12

import (
	"fmt"
	"log"

	"github.com/wistler/aoc-2020/internal"

	"github.com/wistler/aoc-2020/internal/vector"
)

var dirs map[rune]vector.Vec = map[rune]vector.Vec{
	'N': vector.Make(0, 1),
	'S': vector.Make(0, -1),
	'W': vector.Make(-1, 0),
	'E': vector.Make(1, 0),
}

var dirSequence []rune = []rune{'E', 'S', 'W', 'N'}

func rotate(facing rune, quads int) rune {
	_, facingIndex := internal.ContainsRune(dirSequence, facing)
	return dirSequence[(facingIndex+quads+len(dirSequence))%len(dirSequence)]
}

func part1(input []string) int {
	log.SetPrefix("Day 12: Part 1: ")
	log.SetFlags(0)

	pos := vector.Make(0, 0)
	facing := 'E'

	for _, step := range input {
		var ch rune
		var n int
		_, err := fmt.Sscanf(step, "%c%d", &ch, &n)
		if err != nil {
			panic(err)
		}
		switch ch {
		case 'N', 'S', 'E', 'W':
			pos.Add(dirs[ch].Mult(float64(n)))
		case 'L', 'R':
			n = n / 90
			if ch == 'L' {
				n = -n
			}
			facing = rotate(facing, n)
		case 'F':
			pos.Add(dirs[facing].Mult(float64(n)))
		}
	}

	sum := int(pos.ManhattenDistance())

	log.Printf("Answer: %v", sum)
	return sum
}

func part2(input []string) int {
	log.SetPrefix("Day 12: Part 2: ")
	log.SetFlags(0)

	pos := vector.Make(0, 0)
	waypoint := []struct {
		dir   rune
		units int
	}{
		{'E', 10},
		{'N', 1},
		{'S', 0},
		{'W', 0},
	}

	for _, step := range input {
		var ch rune
		var n int
		_, err := fmt.Sscanf(step, "%c%d", &ch, &n)
		if err != nil {
			panic(err)
		}
		switch ch {
		case 'N', 'S', 'E', 'W':
			for i := 0; i < len(waypoint); i++ {
				if waypoint[i].dir == ch {
					waypoint[i].units += n
				}
			}
		case 'L', 'R':
			n = n / 90
			if ch == 'L' {
				n = -n
			}
			for i := 0; i < len(waypoint); i++ {
				waypoint[i].dir = rotate(waypoint[i].dir, n)
			}
		case 'F':
			w := vector.Make(0, 0)
			for _, wp := range waypoint {
				w.Add(dirs[wp.dir].Mult(float64(wp.units)))
			}
			pos.Add(w.Mult(float64(n)))
		}
	}

	sum := int(pos.ManhattenDistance())

	log.Printf("Answer: %v", sum)
	return sum
}
