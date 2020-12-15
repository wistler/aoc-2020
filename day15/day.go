package day15

import (
	"log"
)

func lastIndexOf(haystack []int, needle int, skip int) (int, bool) {
	for i := len(haystack) - 1; i >= 0; i-- {
		if haystack[i] == needle {
			if skip > 0 {
				skip--
				continue
			}
			return i, true
		}
	}
	return -1, false
}

func part1(input []int) int {
	log.SetPrefix("Day 15: Part 1: ")
	log.SetFlags(0)

	var mem [2020]int
	n := -1
	for i := 0; i < 2020; i++ {
		if i < len(input) {
			n = input[i]
		} else {
			lastIndex, found := lastIndexOf(mem[:i], n, 1)
			if found {
				n = (i - 1) - lastIndex
			} else {
				n = 0
			}
		}
		mem[i] = n
	}
	log.Printf("Answer: %v", n)
	return n
}

func part2(input []int) int {
	log.SetPrefix("Day 15: Part 2: ")
	log.SetFlags(0)

	mem := make(map[int]int) // map of value to last known index
	lastValue, currValue := -1, -1
	for i := 0; i < 30_000_000; i++ {
		if i < len(input) {
			currValue = input[i]
			mem[currValue] = i
		} else {
			lastIndex, ok := mem[lastValue]
			if ok {
				currValue = (i - 1) - lastIndex
			} else {
				currValue = 0
			}
		}
		if i > 0 {
			mem[lastValue] = i - 1
		}
		lastValue = currValue
	}
	log.Printf("Answer: %v", currValue)
	return currValue
}
