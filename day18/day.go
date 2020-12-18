package day18

import (
	"errors"
	"log"
	"strconv"
	"unicode"
)

type element struct {
	op   string
	term int
}

func op(ch rune) element {
	return element{
		op: string(ch),
	}
}

func term(num int) element {
	return element{
		op:   "",
		term: num,
	}
}

func (e *element) isTerm() bool {
	return e.op == ""
}

func (e *element) isOp() bool {
	return e.op != ""
}

func (e element) String() string {
	if e.isOp() {
		return e.op
	}
	return strconv.Itoa(e.term)
}

func pushOntoStack(stack []element, num int, delayMult bool, debug bool) ([]element, error) {
	if debug {
		log.Printf("push: %v <- %v", stack, num)
	}
	if len(stack) > 0 {
		if stack[len(stack)-1].isTerm() {
			return stack, errors.New("Consequtive Terms")
		}

		op := stack[len(stack)-1].op
		switch op {
		case "(":
			// do nothing
		case "+":
			if !stack[len(stack)-2].isTerm() {
				return stack, errors.New("Missing Term on stack[-2]")
			}
			term := stack[len(stack)-2].term
			num += term
			stack = stack[:len(stack)-2]
		case "*":
			if !stack[len(stack)-2].isTerm() {
				return nil, errors.New("Missing Term on stack[-2]")
			}
			if !delayMult {
				term := stack[len(stack)-2].term
				num *= term
				stack = stack[:len(stack)-2]
			}
		default:
			return nil, errors.New("Unexpected element on stack")
		}
	}
	stack = append(stack, term(num))
	if debug {
		log.Printf("push: result: %v", stack)
	}
	return stack, nil
}

func collapseStack(stack []element, debug bool) ([]element, error) {
	if debug {
		log.Printf("collaping: %v", stack)
	}
Loop:
	for len(stack) >= 3 && stack[len(stack)-1].isTerm() && stack[len(stack)-2].isOp() {
		op := stack[len(stack)-2].op
		switch op {
		case "*":
			term1 := stack[len(stack)-3].term
			term2 := stack[len(stack)-1].term
			newTerm := term1 * term2
			stack = stack[:len(stack)-3]
			if debug {
				log.Printf("collaping: %v >> %v * %v <<", stack, term1, term2)
			}

			stack = append(stack, term(newTerm))
			if debug {
				log.Printf("collaping: %v", stack)
			}

		case "(":
			break Loop
		}
	}
	if debug {
		log.Printf("collaping: result: %v", stack)
	}
	return stack, nil
}

func solve(line string, plusPriority bool, debug bool) int {
	var err error
	ans := 0
	stack := []element{}
	partialNumber := ""
	for i, r := range line {
		if unicode.IsDigit(r) {
			partialNumber = partialNumber + string(r)
			continue
		}
		if len(partialNumber) > 0 {
			num, err := strconv.Atoi(partialNumber)
			if err != nil {
				panic(err)
			}
			stack, err = pushOntoStack(stack, num, plusPriority, debug)
			if err != nil {
				log.Panicln(err, stack, i, r, line)
			}
			partialNumber = ""
		}
		switch r {
		case '+', '*', '(':
			if debug {
				log.Printf("push: %v <- %q", stack, r)
			}

			stack = append(stack, op(r))
		case ')':
			if debug {
				log.Printf("push: %v <- %q", stack, r)
			}
			stack, err = collapseStack(stack, debug)
			if err != nil {
				log.Panicln(err, stack, i, r, line)
			}
			if stack[len(stack)-1].isTerm() && stack[len(stack)-2].op == "(" {
				num := stack[len(stack)-1].term
				stack = stack[:len(stack)-2] // remove last 2 elements

				stack, err = pushOntoStack(stack, num, plusPriority, debug)
				if err != nil {
					log.Panicln(err, stack, i, r, line)
				}

			} else {
				log.Panicln("Unexpected ')'", stack, i, r, line)
			}
		}
	}
	if len(partialNumber) > 0 {
		num, err := strconv.Atoi(partialNumber)
		if err != nil {
			panic(err)
		}
		stack, err = pushOntoStack(stack, num, plusPriority, debug)
		if err != nil {
			log.Panicln(err, stack, line)
		}
		partialNumber = ""
	}
	stack, err = collapseStack(stack, debug)
	if err != nil {
		log.Panicln(err, stack, line)
	}

	if debug {
		log.Printf("Stack: %v", stack)
	}
	if len(stack) != 1 {
		log.Panicln("Unexpected terminal stack size:", stack, line)
	}
	if !stack[0].isTerm() {
		log.Panicln("Unexpected terminal stack state:", stack, line)
	}
	ans = stack[0].term
	return ans
}

func part1(input []string) int {
	log.SetPrefix("Day 18: Part 1: ")
	log.SetFlags(0)

	sum := 0
	for _, line := range input {
		sum += solve(line, false, false)
	}

	log.Printf("Answer: %v", sum)
	return sum
}

func part2(input []string) int {
	log.SetPrefix("Day 18: Part 2: ")
	log.SetFlags(0)

	sum := 0
	for _, line := range input {
		sum += solve(line, true, false)
	}

	log.Printf("Answer: %v", sum)
	return sum
}
