package day19

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func createRuleMap(rules []string, debug bool) map[int]string {
	ruleMap := make(map[int]string)

	for {
		skipped := false
		parsed := false
	Loop1:
		for _, line := range rules {
			parts := strings.Split(line, ": ")
			ri, err := strconv.Atoi(parts[0])
			if err != nil {
				log.Panicln(err, line)
			}
			if _, ok := ruleMap[ri]; ok {
				continue
			}

			seq := strings.Split(parts[1], " ")
			rul := ""
			rch := ""
			for _, si := range seq {
				n, err := fmt.Sscanf(si, "\"%1s\"", &rch)
				if err == nil && n == 1 {
					rul += rch
					continue
				}
				if si == "|" {
					rul += si
					continue
				}
				i, err := strconv.Atoi(si)
				if err != nil {
					log.Panicln(err, line)
				}
				ru, uok := ruleMap[i]
				if uok {
					rul += ru
				} else {
					if debug {
						log.Println("Skippping rule:", line)
					}
					skipped = true
					continue Loop1
				}
			}
			if len(rul) == 1 {
				ruleMap[ri] = rul
			} else {
				ruleMap[ri] = "(" + rul + ")"
			}

			if debug {
				log.Println("Parsed rule:", line)
			}
			parsed = true
		}
		if !skipped || !parsed {
			// either all done, or recursive rules.
			break
		}
	}

	if debug {
		for k, v := range ruleMap {
			log.Println("[", k, "] =>", v)
		}
	}
	return ruleMap
}

func getRulesAndMessages(input []string) (rules []string, messages []string) {
	phase1 := true
	for _, line := range input {
		if strings.TrimSpace(line) == "-skip-" {
			continue
		}
		if strings.TrimSpace(line) == "" {
			phase1 = false
			continue
		}
		if phase1 {
			rules = append(rules, line)
		} else {
			messages = append(messages, line)
		}
	}
	return
}

func part1(input []string, debug bool) int {
	log.SetPrefix("Day 19: Part 1: ")
	log.SetFlags(0)

	rules, messages := getRulesAndMessages(input)
	ruleMap := createRuleMap(rules, debug)

	rx := "^" + ruleMap[0] + "$"

	sum := 0
	for _, mesg := range messages {
		matched, err := regexp.MatchString(rx, mesg)
		if err != nil {
			panic(err)
		}
		if matched {
			sum++
		}
	}

	log.Printf("Answer: %v", sum)
	return sum
}

func part2(input []string, debug bool) int {
	log.SetPrefix("Day 19: Part 2: ")
	log.SetFlags(0)

	for i := 0; i < len(input); i++ {
		if strings.HasPrefix(input[i], "0:") ||
			strings.HasPrefix(input[i], "8:") ||
			strings.HasPrefix(input[i], "11:") {

			input[i] = "-skip-"
		}
	}

	rules, messages := getRulesAndMessages(input)
	ruleMap := createRuleMap(rules, debug)

	// ruleMap[8] = fmt.Sprintf("(?P<8>%s+)", ruleMap[42])
	ruleMap[11] = fmt.Sprintf("(?P<11>(?P<11a>%s+)(?P<11b>%s+))", ruleMap[42], ruleMap[31])
	// ruleMap[0] = fmt.Sprintf("%s%s", ruleMap[8], ruleMap[11])

	// Rules 0, 8 & 11 summarize to: [42]{n}[31]{m}  n > m
	rx := "^" + ruleMap[11] + "$"

	sum := 0
	re := regexp.MustCompile(rx)
	// i11 := re.SubexpIndex("11")
	// i8 := re.SubexpIndex("8")
	i11a := re.SubexpIndex("11a")
	i11b := re.SubexpIndex("11b")

	// re11 := regexp.MustCompile(ruleMap[42] + ruleMap[31])
	re11a := regexp.MustCompile(ruleMap[42])
	re11b := regexp.MustCompile(ruleMap[31])

	for _, mesg := range messages {
		matched := re.MatchString(mesg)
		if matched {
			m := re.FindStringSubmatch(mesg)
			// s8 := m[i8]
			// s11 := m[i11]
			s11a := m[i11a]
			s11b := m[i11b]

			m11b := re11b.FindAllString(s11b, -1)
			n11b := len(m11b)

			m11a := re11a.FindAllString(s11a, -1)
			n11a := len(m11a)

			// Rules 0, 8 & 11 summarize to: [42]{n}[31]{m}  n > m
			valid := n11a > n11b

			if debug {
				log.Print("\n\n\n")
				log.Println("mesg :", mesg)
				// log.Printf("%#v\n", m)
				// log.Println("8|11 :", s8+"|"+s11)
				// log.Println("8|11 :", s8+"|"+s11a+"|"+s11b)
				log.Println("11a|b:", s11a+"|"+s11b)
				log.Println("n    :", n11a, "|", n11b)
				log.Println("m    :", m11a, "|", m11b)
				log.Println("valid:", valid)
			}

			if valid {
				sum++
			}
		}
	}

	log.Printf("Answer: %v", sum)
	return sum
}
