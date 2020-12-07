package day06

import (
	"log"
	"strings"

	"github.com/wistler/aoc-2020/internal"
)

func getUniqueAnswers(group []string) []string {
	uniq := []string{}
	for _, str := range group {
		for i := 0; i < len(str); i++ {
			r := string(str[i])
			if ok, _ := internal.Contains(uniq, r); !ok {
				uniq = append(uniq, r)
			}
		}
	}
	return uniq
}

func getCommonAnswers(group []string) []string {
	common := []string{}
	for n, str := range group {
		if n == 0 {
			for i := 0; i < len(str); i++ {
				ch := string(str[i])
				common = append(common, ch)
			}
			continue
		}
		for j := 0; j < len(common); {
			found := false
			for i := 0; i < len(str); i++ {
				if string(str[i]) == common[j] {
					found = true
					break
				}
			}
			if !found {
				internal.RemoveIndex(&common, j)
			} else {
				j++
			}
		}
	}
	return common
}

func getGroups(input []string) [][]string {
	groups := [][]string{}
	current := []string{}
	for _, line := range input {
		if strings.Trim(line, " ") == "" {
			groups = append(groups, current)
			current = []string{}
		} else {
			current = append(current, line)
		}
	}
	if len(current) != 0 {
		groups = append(groups, current)
		current = nil
	}
	return groups
}

func part1(input []string) int {
	log.SetPrefix("Day 6: Part 1: ")
	log.SetFlags(0)

	groups := getGroups(input)
	sum := 0
	for _, group := range groups {
		sum += len(getUniqueAnswers(group))
	}

	log.Printf("Answer: %v", sum)
	return sum
}

func part2(input []string) int {
	log.SetPrefix("Day 6: Part 2: ")
	log.SetFlags(0)

	groups := getGroups(input)
	sum := 0
	for _, group := range groups {
		sum += len(getCommonAnswers(group))
	}

	log.Printf("Answer: %v", sum)
	return sum
}
