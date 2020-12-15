package day14

import (
	"fmt"
	"log"
	"strings"

	"github.com/wistler/aoc-2020/internal"
)

const dataWidth = 36

type dataType [dataWidth]byte

func bitmask(value dataType, mask dataType) dataType {
	var result dataType
	for i := 0; i < dataWidth; i++ {
		switch mask[i] {
		case 'X':
			result[i] = value[i]
		case '0', '1':
			result[i] = mask[i]
		}
	}
	return result
}

func bitmask2(value dataType, mask dataType) dataType {
	var result dataType
	for i := 0; i < dataWidth; i++ {
		switch mask[i] {
		case '0':
			result[i] = value[i]
		case 'X', '1':
			result[i] = mask[i]
		}
	}
	return result
}

func floating(addr dataType) []dataType {
	result := []dataType{}
	if ok, i := internal.ContainsByte(addr[:], 'X'); ok {
		addr[i] = '0'
		result = append(result, floating(addr)...)
		addr[i] = '1'
		result = append(result, floating(addr)...)
	} else {
		result = append(result, addr)
	}
	return result
}

func intVal(value dataType) int {
	result := 0
	for i := 0; i < dataWidth; i++ {
		if value[i] == '1' {
			result += 1 << (dataWidth - 1 - i)
		}
	}
	return result
}

func byteVal(value int) dataType {
	var result dataType
	for i := 0; i < dataWidth; i++ {
		bit := value & (1 << (dataWidth - 1 - i))
		if bit == 0 {
			result[i] = '0'
		} else {
			result[i] = '1'
		}
	}
	return result
}

func strToBytes(value string) dataType {
	if len(value) != dataWidth {
		panic("illegal word size")
	}
	var result dataType
	for i := 0; i < dataWidth; i++ {
		result[i] = value[i]
	}
	return result
}

func part1(input []string) int {
	log.SetPrefix("Day 14: Part 1: ")
	log.SetFlags(0)

	var mask dataType
	mem := make(map[int]dataType)

	for _, inst := range input {
		if strings.HasPrefix(inst, "mask") {
			mask = strToBytes(inst[len("mask = "):])
		} else {
			var addr, value int
			n, err := fmt.Sscanf(inst, "mem[%d] = %d", &addr, &value)
			if err != nil {
				log.Println(inst)
				panic(err)
			}
			if n != 2 {
				panic("input format error")
			}
			mem[addr] = bitmask(byteVal(value), mask)
		}
	}

	sum := 0
	for _, value := range mem {
		sum += intVal(value)
	}

	log.Printf("Answer: %v", sum)
	return sum
}

func part2(input []string) int {
	log.SetPrefix("Day 14: Part 2: ")
	log.SetFlags(0)

	var mask dataType
	mem := make(map[int]int)

	for _, inst := range input {
		if strings.HasPrefix(inst, "mask") {
			mask = strToBytes(inst[len("mask = "):])
		} else {
			var addr, value int
			n, err := fmt.Sscanf(inst, "mem[%d] = %d", &addr, &value)
			if err != nil {
				log.Println(inst)
				panic(err)
			}
			if n != 2 {
				panic("input format error")
			}
			for _, addr2 := range floating(bitmask2(byteVal(addr), mask)) {
				mem[intVal(addr2)] = value
			}
		}
	}

	sum := 0
	for _, value := range mem {
		sum += value
	}
	log.Printf("Answer: %v", sum)
	return sum
}
