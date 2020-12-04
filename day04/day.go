package day04

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func checkPassportField(key string, value string) (bool, error) {
	switch key {
	case "byr":
		if v, err := strconv.Atoi(value); err != nil || v < 1920 || v > 2002 {
			return false, fmt.Errorf("byr %q", value)
		}
	case "iyr":
		if v, err := strconv.Atoi(value); err != nil || v < 2010 || v > 2020 {
			return false, fmt.Errorf("iyr %q", value)
		}
	case "eyr":
		if v, err := strconv.Atoi(value); err != nil || v < 2020 || v > 2030 {
			return false, fmt.Errorf("eyr %q", value)
		}
	case "hgt":
		v := 0
		if n, err := fmt.Sscanf(value, "%dcm", &v); err == nil && n == 1 {
			if v < 150 || v > 193 {
				return false, fmt.Errorf("hgt cm %q", value)
			}
		} else if n, err := fmt.Sscanf(value, "%din", &v); err == nil && n == 1 {
			if v < 59 || v > 76 {
				return false, fmt.Errorf("hgt in %q", value)
			}
		} else {
			return false, fmt.Errorf("hgt %q", value)
		}
	case "hcl":
		if matched, err := regexp.MatchString("^#[0-9a-f]{6}$", value); err != nil || !matched {
			return false, fmt.Errorf("hcl %q", value)
		}
	case "ecl":
		if matched, err := regexp.MatchString("^(amb|blu|brn|gry|grn|hzl|oth)$", value); err != nil || !matched {
			return false, fmt.Errorf("ecl %q", value)
		}
	case "pid":
		if matched, err := regexp.MatchString("^[0-9]{9}$", value); err != nil || !matched {
			return false, fmt.Errorf("pid %q", value)
		}
	case "cid":
		// do nothing
	default:
		return false, fmt.Errorf("Unknown field: %q", key)
	}

	return true, nil
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
		if validateFieldContent {
			if valid, _ := checkPassportField(key, value); !valid {
				// log.Printf("Field validation error: %s", err)
				return false
			}
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

// a passport is grouped into a set of "key:value" strings
func getPassports(input []string) [][]string {
	passports := [][]string{}
	current := []string{}
	for _, line := range input {
		if strings.Trim(line, " ") == "" {
			passports = append(passports, current)
			current = []string{}
		} else {
			current = append(current, strings.Split(line, " ")...)
		}
	}
	if len(current) != 0 {
		passports = append(passports, current)
		current = []string{}
	}
	return passports
}

func part1(input []string) int {
	log.SetPrefix("Day 4: Part 1: ")
	log.SetFlags(0)

	validPassports := 0
	passports := getPassports(input)
	for _, passport := range passports {
		if checkPassport(passport, false) {
			validPassports++
		}
	}

	log.Printf("Answer: %v", validPassports)
	return validPassports
}

func part2(input []string) int {
	log.SetPrefix("Day 4: Part 2: ")
	log.SetFlags(0)

	validPassports := 0
	passports := getPassports(input)
	for _, passport := range passports {
		if checkPassport(passport, true) {
			validPassports++
		}
	}

	log.Printf("Answer: %v", validPassports)
	return validPassports
}
