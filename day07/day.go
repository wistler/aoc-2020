package day07

import (
	"log"
	"strconv"
	"strings"

	"github.com/wistler/aoc-2020/internal"
)

func getOuterAndInnerBags(rule string) (string, string) {
	parts := strings.Split(rule, "s contain")
	return parts[0], parts[1]
}

func getOuterAndInnerBagsAsList(rule string) (string, []string) {
	parts := strings.Split(rule, "s contain")
	outer := parts[0]

	inner := []string{}
	if !strings.Contains(parts[1], "no other") {
		inner = strings.Split(parts[1], ",")
		for i, str := range inner {
			inner[i] = strings.TrimRightFunc(
				strings.TrimSpace(str),
				func(r rune) bool { return r == rune('.') || r == rune('s') })
		}
	}

	return outer, inner
}

func part1(rules []string, find string) int {
	log.SetPrefix("Day 7: Part 1: ")
	log.SetFlags(0)

	containers := []string{}
	ruleMap := make(map[string]string, len(rules))

	for _, rule := range rules {
		outer, inner := getOuterAndInnerBags(rule)
		ruleMap[outer] = inner
		if strings.Contains(inner, find) {
			containers = append(containers, outer)
		}
	}

	for {
		moreFound := false
		for outer, inner := range ruleMap {
			for _, container := range containers {
				if strings.Contains(inner, container) {
					if ok, _ := internal.Contains(containers, outer); !ok {
						containers = append(containers, outer)
						moreFound = true
					}
				}
			}
		}
		if !moreFound {
			// if there were any new entries added to the containers list, we'll need
			// to re-vist the ruleMap from the top again; else break.
			break
		}
	}

	sum := len(containers)
	log.Printf("Answer: %v", sum)
	return sum
}

func lookup(ruleMap map[string][]string, parent string) int {
	children := ruleMap[parent]
	if children == nil {
		panic("parent not found: " + parent)
	}
	sum := 0
	for _, childRule := range children {
		parts := strings.SplitN(childRule, " ", 2)
		childCount, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		childType := parts[1]
		sum += childCount + childCount*lookup(ruleMap, childType)
	}
	return sum
}

func part2(rules []string, find string) int {
	log.SetPrefix("Day 7: Part 2: ")
	log.SetFlags(0)

	ruleMap := make(map[string][]string, len(rules))

	for _, rule := range rules {
		outer, inner := getOuterAndInnerBagsAsList(rule)
		ruleMap[outer] = inner
	}

	sum := lookup(ruleMap, find)

	log.Printf("Answer: %v", sum)
	return sum
}
