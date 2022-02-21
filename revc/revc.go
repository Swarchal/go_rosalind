package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func revc(s string) string {
	// reverse complement a DNA sequence
	nuc_map := map[string]string{
		"A": "T",
		"T": "A",
		"C": "G",
		"G": "C",
	}
	var sb strings.Builder
	for i := len(s) - 1; i >= 0; i-- {
		char := string(s[i])
		sb.WriteString(nuc_map[char])
	}
	return sb.String()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		seq := scanner.Text()
		seq_revc := revc(seq)
		fmt.Println(seq_revc)
	}
}


