package day08

import (
	"errors"
	"fmt"
	"log"
)

func process(code []string, pc int, acc int) (int, int) {
	op := ""
	arg := 0
	n, err := fmt.Sscanf(code[pc], "%s %d", &op, &arg)
	if err != nil {
		panic(err)
	}
	if n != 2 {
		panic(fmt.Sprintf("Error parsing instruction on line %v: %q", pc, code[pc]))
	}
	switch op {
	case "nop":
		return pc + 1, acc
	case "acc":
		return pc + 1, acc + arg
	case "jmp":
		return pc + arg, acc
	}
	panic(fmt.Sprintf("Unknown op code at %v: %v", pc, op))
}

func run(code []string, pc int, acc int) (int, int, error) {
	visited := make([]int, len(code))
	for pc < len(code) && visited[pc] == 0 {
		visited[pc] = 1
		pc, acc = process(code, pc, acc)
	}
	if pc < len(code) && visited[pc] == 1 {
		return pc, acc, errors.New("Infinite Loop detected")
	}
	return pc, acc, nil
}

func part1(input []string) int {
	log.SetPrefix("Day 8: Part 1: ")
	log.SetFlags(0)

	_, acc, _ := run(input, 0, 0)

	log.Printf("Answer: %v", acc)
	return acc
}

func part2(input []string) int {
	log.SetPrefix("Day 8: Part 2: ")
	log.SetFlags(0)

	lastModified := 0
	acc := 0
	var err error = nil
	code := make([]string, len(input))
	for {
		modified := false
		for i, line := range input {
			if !modified && i > lastModified {
				if line[0:3] == "nop" {
					line = "jmp" + line[3:]
					log.Printf("Modified line %v to: %q", i, line)
					lastModified = i
					modified = true
				} else if line[0:3] == "jmp" {
					line = "nop" + line[3:]
					log.Printf("Modified line %v to: %q", i, line)
					lastModified = i
					modified = true
				}
			}
			code[i] = line
		}

		_, acc, err = run(code, 0, 0)
		if err == nil {
			break
		}
		if lastModified+1 == len(input) {
			panic("Reached end of the line, but condition not met.")
		}
	}

	log.Printf("Answer: %v", acc)
	return acc
}
