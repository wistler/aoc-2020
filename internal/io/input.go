package io

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// SplitOnNewLines splits a string on new line termination characters, in an OS-independant way
func SplitOnNewLines(data string) []string {
	if strings.Contains(data, "\r\n") {
		return strings.Split(data, "\r\n")
	}
	return strings.Split(data, "\n")
}

// ReadInputFile returns all line in the file as a string slice
func ReadInputFile(path string) []string {
	data, err := ioutil.ReadFile(path)
	check(err)

	return SplitOnNewLines(string(data))
}

// ReadInputFileAsInts returns all line in the file as a int slice
func ReadInputFileAsInts(path string) []int {
	tmp := ReadInputFile("./input.txt")
	input := make([]int, len(tmp))
	for i, t := range tmp {
		ti, err := strconv.Atoi(t)
		if err != nil {
			panic(err)
		}
		input[i] = ti
	}
	return input
}

// SplitIntoIntArr hi
func SplitIntoIntArr(line string, sep string) []int {
	parts := strings.Split(line, sep)
	result := []int{}
	for _, p := range parts {
		i, err := strconv.Atoi(p)
		if err != nil {
			panic(err)
		}
		result = append(result, i)
	}
	return result
}
