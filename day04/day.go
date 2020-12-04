package day04

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func checkPassportField(key string, value string) bool {
	switch key {
	case "byr":
		v, err := strconv.Atoi(value)
		if err != nil {
			return false
		}
		if v < 1920 || v > 2002 {
			return false
		}
	case "iyr":
		v, err := strconv.Atoi(value)
		if err != nil {
			return false
		}
		if v < 2010 || v > 2020 {
			return false
		}
	case "eyr":
		v, err := strconv.Atoi(value)
		if err != nil {
			return false
		}
		if v < 2020 || v > 2030 {
			return false
		}
	case "hgt":
		if strings.HasSuffix(value, "cm") {
			v := 0
			n, err := fmt.Sscanf(value, "%dcm", &v)
			if err != nil {
				return false
			}
			if n != 1 {
				return false
			}
			if v < 150 || v > 193 {
				return false
			}
		} else if strings.HasSuffix(value, "in") {
			v := 0
			n, err := fmt.Sscanf(value, "%din", &v)
			if err != nil {
				return false
			}
			if n != 1 {
				return false
			}
			if v < 59 || v > 76 {
				return false
			}
		} else {
			return false
		}
	case "hcl":
		if !strings.HasPrefix(value, "#") || strings.ToLower(value) != value {
			return false
		}
		_, err := strconv.ParseInt(value[1:], 16, 0)
		if err != nil {
			return false
		}
	case "ecl":
		if value != "amb" && value != "blu" && value != "brn" && value != "gry" && value != "grn" && value != "hzl" && value != "oth" {
			return false
		}
	case "pid":
		if len(value) != 9 {
			return false
		}
		for _, ch := range value {
			if !strings.ContainsAny(string(ch), "1234567890") {
				return false
			}
		}
	case "cid":
		// do nothing
	}

	return true
}

func checkPassport(keyValuesPairs []string, validateFieldContent bool) bool {
	fields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid", "cid"}
	ignoreMissing := []string{"cid"}

	var key, value string
	for _, pair := range keyValuesPairs {
		n, err := fmt.Sscanf(pair, "%3s:%s", &key, &value)
		if err != nil {
			panic(err)
		}
		if n != 2 {
			panic("Format error")
		}
		found := false
		for i, f := range fields {
			if key == f {
				fields[i] = fields[len(fields)-1]
				fields[len(fields)-1] = ""
				fields = fields[:len(fields)-1]
				found = true
				break
			}
		}
		if !found {
			panic(fmt.Sprintf("Unknown field: %q", key))
		}
		if validateFieldContent && checkPassportField(key, value) == false {
			return false
		}
	}

	if len(fields) != 0 {
		for _, fi := range fields {
			found := false
			for _, ig := range ignoreMissing {
				if fi == ig {
					found = true
					break
				}
			}
			if !found {
				return false
			}
		}
	}
	return true
}

func groupIntoPassportsAndValidate(input []string, validateFields bool) int {
	validPassports := 0
	passport := []string{}
	for _, line := range input {
		if strings.Trim(line, " ") == "" {
			if checkPassport(passport, validateFields) {
				validPassports++
			}
			passport = []string{}
		} else {
			passport = append(passport, strings.Split(line, " ")...)
		}
	}
	if len(passport) != 0 {
		if checkPassport(passport, validateFields) {
			validPassports++
		}
		passport = []string{}
	}

	return validPassports
}

func part1(input []string) int {
	log.SetPrefix("Day 4: Part 1: ")
	log.SetFlags(0)

	ans := groupIntoPassportsAndValidate(input, false)
	log.Printf("Answer: %v", ans)
	return ans
}

func part2(input []string) int {
	log.SetPrefix("Day 4: Part 2: ")
	log.SetFlags(0)

	ans := groupIntoPassportsAndValidate(input, true)
	log.Printf("Answer: %v", ans)
	return ans
}
