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

func findLocations(s, t string) []int {
	var locations []int
	t_len := len(t)
	for i := 0; i < len(s)-t_len; i++ {
		if s[i:i+t_len] == t {
			locations = append(locations, i+1)
		}
	}
	return locations
}

func printSlice(s []int) {
	for _, i := range s {
		fmt.Printf("%v ", i)
	}
	fmt.Println()
}

func main() {
	s, t := parseInput(os.Args[1])
	locations := findLocations(s, t)
	printSlice(locations)
}
