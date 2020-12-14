package day14

import (
	"fmt"
	"log"
	"strings"

	"github.com/wistler/aoc-2020/internal"
)

func bitmask(value [36]byte, mask [36]byte) [36]byte {
	var result [36]byte
	for i := 0; i < 36; i++ {
		switch mask[i] {
		case 'X':
			result[i] = value[i]
		case '0', '1':
			result[i] = mask[i]
		}
	}
	return result
}

func bitmask2(value [36]byte, mask [36]byte) [36]byte {
	var result [36]byte
	for i := 0; i < 36; i++ {
		switch mask[i] {
		case '0':
			result[i] = value[i]
		case 'X', '1':
			result[i] = mask[i]
		}
	}
	return result
}

func floating(addr [36]byte) [][36]byte {
	result := [][36]byte{}
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

func intVal(value [36]byte) int {
	result := 0
	for i := 0; i < 36; i++ {
		if value[i] == '1' {
			result += 1 << (35 - i)
		}
	}
	return result
}

func byteVal(value int) [36]byte {
	var result [36]byte
	for i := 0; i < 36; i++ {
		bit := value & (1 << (35 - i))
		if bit == 0 {
			result[i] = '0'
		} else {
			result[i] = '1'
		}
	}
	return result
}

func strToBytes(value string) [36]byte {
	if len(value) != 36 {
		panic("illegal word size")
	}
	var result [36]byte
	for i := 0; i < 36; i++ {
		result[i] = value[i]
	}
	return result
}

func part1(input []string) int {
	log.SetPrefix("Day 14: Part 1: ")
	log.SetFlags(0)

	var mask [36]byte
	mem := make(map[int][36]byte)

	for _, inst := range input {
		if strings.HasPrefix(inst, "mask") {
			mask = strToBytes(inst[len("mask = "):])
		} else {
			var addr, data int
			n, err := fmt.Sscanf(inst, "mem[%d] = %d", &addr, &data)
			if err != nil {
				log.Println(inst)
				panic(err)
			}
			if n != 2 {
				panic("input format error")
			}
			mem[addr] = bitmask(byteVal(data), mask)
		}
	}

	sum := 0
	for _, data := range mem {
		sum += intVal(data)
	}

	log.Printf("Answer: %v", sum)
	return sum
}

func part2(input []string) int {
	log.SetPrefix("Day 14: Part 2: ")
	log.SetFlags(0)

	var mask [36]byte
	mem := make(map[int]int)

	for _, inst := range input {
		if strings.HasPrefix(inst, "mask") {
			mask = strToBytes(inst[len("mask = "):])
		} else {
			var addr, data int
			n, err := fmt.Sscanf(inst, "mem[%d] = %d", &addr, &data)
			if err != nil {
				log.Println(inst)
				panic(err)
			}
			if n != 2 {
				panic("input format error")
			}
			for _, addr2 := range floating(bitmask2(byteVal(addr), mask)) {
				mem[intVal(addr2)] = data
			}
		}
	}

	sum := 0
	for _, data := range mem {
		sum += data
	}
	log.Printf("Answer: %v", sum)
	return sum
}
