package main

import (
	"bufio"
	"fmt"
	"os"
)

func parseLine(line string) (float32, float32, float32) {
	// parse line into 3 ints: populations of (AA, Aa, aa)
	var k, m, n float32
	fmt.Sscanf(line, "%f %f %f", &k, &m, &n)
	return k, m, n
}

func calcProbDominant(k, m, n float32) float32 {
	// prob that two randomly selecting mating organisms will produce an
	// individual posessing a dominant allele (AA or Aa).
	i := m * m + 4.0 * n  * n + 4.0 * m * n - 4.0  *n - m
	j := 4.0 * (k + m + n) * (k + m + n - 1.0)
	rst := 1.0 - (i / j)
	return rst
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		k, m, n := parseLine(line)
		prob := calcProbDominant(k, m, n)
		fmt.Printf("%.5f\n", prob)
	}
}
