package day09

import (
	"log"

	"github.com/wistler/aoc-2020/internal"
)

func getAllSums(numbers []int) []int {
	sums := []int{}
	for _, n1 := range numbers {
		for _, n2 := range numbers {
			if n1 != n2 {
				if ok, _ := internal.ContainsNumber(sums, n1+n2); !ok {
					sums = append(sums, n1+n2)
				}
			}
		}
	}
	return sums
}

func getInvalidNumber(input []int, preamble int) int {
	invalid := 0
	for i := preamble; i < len(input); i++ {
		possibleSums := getAllSums(input[i-preamble : i])
		v := input[i]
		if ok, _ := internal.ContainsNumber(possibleSums, v); !ok {
			invalid = v
			break
		}
	}
	return invalid
}

func part1(input []int, preamble int) int {
	log.SetPrefix("Day 9: Part 1: ")
	log.SetFlags(0)

	invalid := getInvalidNumber(input, preamble)
	log.Printf("Answer: %v", invalid)
	return invalid
}

func part2(input []int, preamble int) int {
	log.SetPrefix("Day 9: Part 2: ")
	log.SetFlags(0)

	invalid := getInvalidNumber(input, preamble)
	result := 0
	for i := 0; i < len(input)-1; i++ {
		for j := i + 1; j < len(input); j++ {
			if invalid == internal.SumOf(input[i:j]) {
				sm := internal.Min(input[i:j])
				lg := internal.Max(input[i:j])
				result = sm + lg
				// log.Print(invalid, sm, lg, result, input[i:j])
				break
			}
		}
		if result != 0 {
			break
		}
	}

	log.Printf("Answer: %v", result)
	return result
}
