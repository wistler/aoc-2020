package day03

import (
	"log"

	"github.com/wistler/aoc-2020/internal/vector"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func checkSlope(mapStrip []string, pos vector.Vec, slope vector.Vec) (int, error) {
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
	return treesHit, nil
}

func part1(input []string) (int, error) {
	log.SetPrefix("Day 3: Part 1: ")
	log.SetFlags(0)

	pos := vector.Make(0, 0)
	slope := vector.Make(3, 1)
	treesHit, err := checkSlope(input, pos, slope)
	if err != nil {
		return treesHit, err
	}
	log.Printf("Answer: %v", treesHit)
	return treesHit, nil
}

func part2(input []string) (int, error) {
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
		treesHit, err := checkSlope(input, pos, slope)
		if err != nil {
			return treesHit, err
		}
		prod *= treesHit
	}

	log.Printf("Answer: %v", prod)
	return prod, nil
}
