package day10

import (
	"log"
	"sort"
)

func part1(input []int) int {
	log.SetPrefix("Day 10: Part 1: ")
	log.SetFlags(0)

	jolt := 0
	diffs := make(map[int]int)

	sort.Ints(input) // sorted adapters
	// log.Print(input)
	for _, adapter := range input {
		diffs[adapter-jolt]++
		jolt = adapter
	}
	diffs[3]++
	// log.Print(diffs)
	ans := diffs[1] * diffs[3]

	log.Printf("Answer: %v", ans)
	return ans
}

func part2(input []int) int {
	log.SetPrefix("Day 10: Part 2: ")
	log.SetFlags(0)

	sort.Ints(input) // sorted adapters
	// log.Print(input)

	outputWanted := input[len(input)-1] + 3
	table := make([]int, outputWanted+1)
	input = append(input, outputWanted)

	table[0] = 1 // seed value

	for i := 0; i < len(table); i++ {
		if table[i] == 0 {
			continue
		}
		for _, adapter := range input {
			if 0 < adapter-i && adapter-i <= 3 && adapter < len(table) {
				table[adapter] += table[i]
				// log.Print(i, adapter, table[i], table[adapter])
			}
		}
	}

	// log.Print(table)
	ans := table[len(table)-1]

	log.Printf("Answer: %v", ans)
	return ans
}
