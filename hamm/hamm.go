package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func parseInput(path string) (string, string) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(content), "\n")
	return lines[0], lines[1]
}

func hammingDist(s, t string) int {
	// assumes strings are the same length
	dist := 0
	for i := 0; i < len(s); i++ {
		if s[i] != t[i] {
			dist++
		}
	}
	return dist
}

func main() {
	path := os.Args[1]
	s, t := parseInput(path)
	dist := hammingDist(s, t)
	fmt.Println(dist)
}
