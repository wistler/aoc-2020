package day16

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/wistler/aoc-2020/internal/io"
)

type interval [2]int
type rule struct {
	name      string
	intervals [2]interval
}

func part1(input []string) int {
	log.SetPrefix("Day 16: Part 1: ")
	log.SetFlags(0)

	intervals := []interval{}
	var lineNumber int
	for ; ; lineNumber++ {
		line := input[lineNumber]
		if strings.TrimSpace(line) == "" {
			break
		}
		var i1, i2, i3, i4 int
		sep := strings.Index(line, ":")
		if sep == -1 {
			panic("format error: Cannot find separator")
		}
		n, err := fmt.Sscanf(line[sep+2:], "%d-%d or %d-%d", &i1, &i2, &i3, &i4)
		if err != nil {
			panic(err.Error() + " : " + line[sep+2:])
		}
		if n != 4 {
			panic("format error: " + line[sep+2:])
		}
		intervals = append(intervals, []interval{{i1, i2}, {i3, i4}}...)
	}

	for ; ; lineNumber++ {
		line := input[lineNumber]
		if strings.HasPrefix(line, "nearby tickets:") {
			lineNumber++
			break
		}
	}

	sum := 0
	for _, line := range input[lineNumber:] {
		for _, num := range strings.Split(line, ",") {
			n, err := strconv.Atoi(num)
			if err != nil {
				panic(num + " : " + err.Error())
			}
			matched := false
			for _, r := range intervals {
				if n >= r[0] && n <= r[1] {
					matched = true
				}
			}
			if !matched {
				sum += n
			}
		}

	}

	log.Printf("Answer: %v", sum)
	return sum
}

func contains(haystack []rule, needle rule) (bool, int) {
	for i, r := range haystack {
		if r == needle {
			return true, i
		}
	}
	return false, -1
}

func removeIndex(from *[]rule, index int) {
	f := *from
	f[index] = f[len(f)-1]
	f[len(f)-1] = rule{}
	f = f[:len(f)-1]
	*from = f
}

func part2(input []string) int {
	log.SetPrefix("Day 16: Part 2: ")
	log.SetFlags(0)

	rules := []rule{}
	var lineNumber int
	for ; ; lineNumber++ {
		line := input[lineNumber]
		if strings.TrimSpace(line) == "" {
			break
		}
		var i1, i2, i3, i4 int
		sep := strings.Index(line, ":")
		if sep == -1 {
			panic("format error: Cannot find separator")
		}
		n, err := fmt.Sscanf(line[sep+2:], "%d-%d or %d-%d", &i1, &i2, &i3, &i4)
		if err != nil {
			panic(err.Error() + " : " + line[sep+2:])
		}
		if n != 4 {
			panic("format error: " + line[sep+2:])
		}
		rules = append(rules, rule{line[:sep], [2]interval{{i1, i2}, {i3, i4}}})
	}

	myTicket := io.SplitIntoIntArr(input[lineNumber+2], ",")
	lineNumber += 5

	isValid := func(r rule, n int) bool {
		valid := false
		for _, i := range r.intervals {
			if n >= i[0] && n <= i[1] {
				valid = true
			}
		}
		return valid
	}

	validTickets := [][]int{myTicket}
	for _, line := range input[lineNumber:] {
		validTicket := true
		for _, num := range strings.Split(line, ",") {
			n, err := strconv.Atoi(num)
			if err != nil {
				panic(num + " : " + err.Error())
			}
			matched := false
			for _, r := range rules {
				if isValid(r, n) {
					matched = true
				}
			}
			if !matched {
				validTicket = false
				break
			}
		}
		if validTicket {
			validTickets = append(validTickets, io.SplitIntoIntArr(line, ","))
		}
	}

	mappedRules := make(map[int]rule)
	unmapped := make([]rule, len(rules))
	for i, r := range rules {
		unmapped[i] = r
	}

	for len(mappedRules) < len(rules) {
		for _, v := range mappedRules {
			if ok, i := contains(unmapped, v); ok {
				removeIndex(&unmapped, i)
			}
		}
		if len(unmapped) == 0 {
			panic("Unmapped rules empty")
		}
		fieldsMapped := 0
		for i := 0; i < len(validTickets[0]); i++ {
			if _, ok := mappedRules[i]; ok {
				continue // skipped mapped fields
			}
			satisfiedRules := []rule{}
			for _, r := range unmapped {
				valid := true
				for _, ticket := range validTickets {
					if !isValid(r, ticket[i]) {
						valid = false
						break
					}
				}
				if valid {
					satisfiedRules = append(satisfiedRules, r)
				}
			}
			if len(satisfiedRules) == 0 {
				log.Panicf("Field not satisfied by any rules: %d, %v", i, unmapped)
			}
			if len(satisfiedRules) == 1 {
				mappedRules[i] = satisfiedRules[0]
				fieldsMapped++
				break
			}
		}
		if fieldsMapped == 0 {
			log.Panicf("Unable to map any more fields with remaining rules: %v\nMapped rules:%v ", unmapped, mappedRules)
		}
	}

	// log.Printf("Rule Map: %v", mappedRules)
	prod := 1

	for i, r := range mappedRules {
		if strings.HasPrefix(r.name, "departure") {
			prod *= myTicket[i]
		}
	}

	log.Printf("Answer: %v", prod)
	return prod
}
