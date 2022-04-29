package main

import (
	"bufio"
	"fmt"
	"os"
)

func parseLine(line string) (int, int) {
	var m, n int
	fmt.Sscanf(line, "%d %d", &m, &n)
	return m, n
}

func calcNRabbits(m, n int) int {
	// total number of rabbit pairs that will remain after the n-th month if all
	// rabbits live for m months, assuming each rabbit pair produces a new pair
	// every month, and reach maturity after 1 month.
	queue := make([]int, n)
	queue[0] = 1
	for i := 1; i < m; i++ {
		queue = nextMonth(queue)
	}
	return totalRabbit(queue)
}

func nextMonth(queue []int) []int {
	// progress rabbit queue one month
	// find number of new offspring
	var nOffspring int
	for i := 1; i < len(queue); i++ {
		nOffspring += queue[i]
	}
	// shift everything along by one to the right
	for j := len(queue) - 1; j > 0; j-- {
		queue[j] = queue[j-1]
	}
	// add nOffspring to start of queue
	queue[0] = nOffspring
	return queue
}

func totalRabbit(queue []int) int {
	sum := 0
	for i := 0; i < len(queue); i++ {
		sum += queue[i]
	}
	return sum
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		m, n := parseLine(line)
		nRabbits := calcNRabbits(m, n)
		fmt.Println(nRabbits)
	}
}
