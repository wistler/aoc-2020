package day05

import (
	"log"
)

func getHalfOfRange(r []int, first bool) []int {
	s := r[0]
	e := r[1]
	if first {
		return []int{s, s + (e-s+1)/2 - 1}
	}
	return []int{s + (e-s+1)/2, e}
}

func getRow(ticket string) int {
	r := []int{0, 127}
	for _, t := range ticket {
		r = getHalfOfRange(r, t == 'F')
	}
	return r[0]
}

func getColumn(ticket string) int {
	r := []int{0, 7}
	for _, t := range ticket {
		r = getHalfOfRange(r, t == 'L')
	}
	return r[0]
}

func getSeatID(ticket string) int {
	return getRow(ticket[:7])*8 + getColumn(ticket[7:])
}

func contains(arr []int, search int) bool {
	for _, value := range arr {
		if search == value {
			return true
		}
	}
	return false
}

func part1(input []string) int {
	log.SetPrefix("Day 5: Part 1: ")
	log.SetFlags(0)

	max := 0
	for _, ticket := range input {
		id := getSeatID(ticket)
		if id > max {
			max = id
		}
	}

	log.Printf("Answer: %v", max)
	return max
}

func part2(input []string) int {
	log.SetPrefix("Day 5: Part 2: ")
	log.SetFlags(0)

	ids := make([]int, len(input))
	max := 0
	min := 0
	for i, ticket := range input {
		id := getSeatID(ticket)
		ids[i] = id
		if i == 0 {
			min = id
			max = id
		} else {
			if id > max {
				max = id
			}
			if id < min {
				min = id
			}
		}
	}

	found := -1
	for i := min; i <= max; i++ {
		if !contains(ids, i) && contains(ids, i-1) && contains(ids, i+1) {
			found = i
			break
		}
	}

	log.Printf("Answer: %v", found)
	return found
}
