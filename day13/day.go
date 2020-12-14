package day13

import (
	"log"
	"strconv"
	"strings"
)

func part1(input []string) int {
	log.SetPrefix("Day 13: Part 1: ")
	log.SetFlags(0)

	earliest, err := strconv.Atoi(input[0])
	if err != nil {
		panic(err)
	}
	busIDs := strings.Split(input[1], ",")
	// delays := make(map[string]int)

	shortestWait := -1
	bestBusID := -1

	for _, busID := range busIDs {
		if busID == "x" {
			continue
		}
		n, err := strconv.Atoi(busID)
		if err != nil {
			panic(err)
		}
		delay := n*(earliest/n+1) - earliest
		if shortestWait < 0 || delay < shortestWait {
			bestBusID = n
			shortestWait = delay
		}
	}
	sum := shortestWait * bestBusID

	log.Printf("Answer: %v", sum)
	return sum
}

func part2(input string, debug bool) int {
	log.SetPrefix("Day 13: Part 2: ")
	log.SetFlags(0)

	busIDs := strings.Split(input, ",")
	busOffsets := make([][2]int, len(busIDs))
	j := 0
	for i, busID := range busIDs {
		if busID == "x" {
			continue
		}
		n, err := strconv.Atoi(busID)
		if err != nil {
			panic(err)
		}
		busOffsets[j][0] = i
		busOffsets[j][1] = n
		j++
	}
	busOffsets = busOffsets[:j]
	if debug {
		log.Println(busOffsets)
	}

	// I'm not math-smart; followed suggestion posted here:
	// https://www.reddit.com/r/adventofcode/comments/kc60ri/2020_day_13_can_anyone_give_me_a_hint_for_part_2/gfnnfm3?utm_source=share&utm_medium=web2x&context=3

	solver := func(busOffsets [][2]int, ts0 int, step int) int {
		ts := ts0
		for n := 1; ; n++ {
			ts = ts0 + step*n

			synced := 0
			for j := 0; j < len(busOffsets); j++ {
				offset := busOffsets[j][0]
				busID := busOffsets[j][1]
				delta := (ts + offset) % busID
				if delta == 0 {
					synced++
				}
			}
			if synced == len(busOffsets) {
				break
			}
		}
		return ts
	}

	step := busOffsets[0][1] // the first busID with offset of 0.
	ts0 := step
	for syncIndex := 2; syncIndex <= len(busOffsets); syncIndex++ {
		if debug {
			log.Println("Solving for syncIndex", syncIndex, "ts", ts0, "step", step)
		}
		ts1 := solver(busOffsets[:syncIndex], ts0, step)

		if debug {
			log.Println(ts0, "->", ts1, busOffsets[:syncIndex])
		}
		ts0 = ts1
		step = 1
		for _, busOffset := range busOffsets[:syncIndex] {
			step *= busOffset[1]
		}
	}

	log.Printf("Answer: %v", ts0)
	return ts0
}
