package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func rnaToPeptide(rna string) string {
	codons := NewCodonMapper()
	var sb strings.Builder
	for i := 0; i < len(rna)-3; i += 3 {
		codon := rna[i : i+3]
		peptide := codons[codon]
		sb.WriteString(peptide)
	}
	return sb.String()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		rna := scanner.Text()
		peptide := rnaToPeptide(rna)
		fmt.Println(peptide)
	}

}
