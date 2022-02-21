package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		dna_seq := scanner.Text()
		rna_seq := strings.Replace(dna_seq, "T", "U", -1)
		fmt.Println(rna_seq)
	}
}
