package day01

import (
	"errors"
	"log"
)

func part1(input []int) (int, error) {
	log.SetPrefix("Day 1: Part 1: ")
	log.SetFlags(0)

	for i := 0; i < len(input)-1; i++ {
		for j := i + 1; j < len(input); j++ {
			if input[i]+input[j] == 2020 {
				ans := input[i] * input[j]
				log.Println("Answer: ", ans)
				return ans, nil
			}
		}
	}

	return 0, errors.New("Unable to find pair of numbers")
}

func part2(input []int) (int, error) {
	log.SetPrefix("Day 1: Part 2: ")
	log.SetFlags(0)

	for i := 0; i < len(input)-1; i++ {
		for j := i + 1; j < len(input); j++ {
			for k := j + 1; k < len(input); k++ {
				if input[i]+input[j]+input[k] == 2020 {
					ans := input[i] * input[j] * input[k]
					log.Println("Answer: ", ans)
					return ans, nil
				}
			}
		}
	}

	return 0, errors.New("Unable to find pair of numbers")
}
