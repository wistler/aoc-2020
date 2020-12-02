package day02

import (
	"fmt"
	"log"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func part1(input []string) (int, error) {
	log.SetPrefix("Day 2: Part 1: ")
	log.SetFlags(0)

	valid := 0
	for i, str := range input {
		var min, max int
		var ch, password string
		n, err := fmt.Sscanf(str, "%d-%d %1s: %s", &min, &max, &ch, &password)
		if err != nil {
			panic(err)
		}
		if n != 4 {
			return 0, fmt.Errorf("Input format error @ %v", i)
		}

		count := strings.Count(password, ch)
		if min <= count && count <= max {
			valid++
		}
	}

	log.Printf("Answer: %v", valid)
	return valid, nil
}

func part2(input []string) (int, error) {
	log.SetPrefix("Day 2: Part 2: ")
	log.SetFlags(0)

	valid := 0
	for i, str := range input {
		var min, max int
		var ch, password string
		n, err := fmt.Sscanf(str, "%d-%d %1s: %s", &min, &max, &ch, &password)
		if err != nil {
			panic(err)
		}
		if n != 4 {
			return 0, fmt.Errorf("Input format error @ %v", i)
		}

		count := 0
		if string(password[min-1]) == ch {
			count++
		}
		if string(password[max-1]) == ch {
			count++
		}
		if count == 1 {
			valid++
		}
	}

	log.Printf("Answer: %v", valid)
	return valid, nil
}
