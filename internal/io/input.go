package io

import (
	"io/ioutil"
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
