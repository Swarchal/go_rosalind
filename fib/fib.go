package main

import (
	"bufio"
	"fmt"
	"os"
)

func parseLine(line string) (int, int) {
	var n, k int
	fmt.Sscanf(line, "%d %d", &n, &k)
	return n, k
}

func calcNRabbits(n, k int) int {
	// number of rabbit pairs after n months
	// given each pair creates k pairs as offspring
	// takes 1 month to reach offspring producing age
	nRabbits := 1
	nRabbitsPrevMonth := 0
	for i := 1; i < n; i++ {
		nOffspring := nRabbitsPrevMonth * k
		nRabbitsPrevMonth = nRabbits
		nRabbits += nOffspring
	}
	return nRabbits
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		n, k := parseLine(line)
		nRabbits := calcNRabbits(n, k)
		fmt.Println(nRabbits)
	}
}
