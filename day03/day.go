package day03

import (
	"log"

	"github.com/wistler/aoc-2020/internal/vector"
)

func checkSlope(mapStrip []string, pos vector.Vec, slope vector.Vec) int {
	treesHit := 0
	y := int(pos[1])
	for y < len(mapStrip) {
		line := mapStrip[y]
		if string(line[int(pos[0])%len(line)]) == "#" {
			treesHit++
		}
		pos.Add(slope)
		y = int(pos[1])
	}
	return treesHit
}

func part1(input []string) int {
	log.SetPrefix("Day 3: Part 1: ")
	log.SetFlags(0)

	pos := vector.Make(0, 0)
	slope := vector.Make(3, 1)
	treesHit := checkSlope(input, pos, slope)
	log.Printf("Answer: %v", treesHit)
	return treesHit
}

func part2(input []string) int {
	log.SetPrefix("Day 3: Part 2: ")
	log.SetFlags(0)

	slopes := []vector.Vec{
		vector.Make(1, 1),
		vector.Make(3, 1),
		vector.Make(5, 1),
		vector.Make(7, 1),
		vector.Make(1, 2),
	}
	prod := 1
	for _, slope := range slopes {
		pos := vector.Make(0, 0)
		treesHit := checkSlope(input, pos, slope)
		prod *= treesHit
	}

	log.Printf("Answer: %v", prod)
	return prod
}
