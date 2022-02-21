package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanRunes)

	for input.Scan() {
		counts[input.Text()]++
	}

	fmt.Println(counts["A"], counts["C"], counts["G"], counts["T"])
}
