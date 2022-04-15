package main

import (
	"bufio"
	"fmt"
	"os"
)

func makeCountMap() map[string]int {
	// codon map from description
	var codonMap = map[string]string{
		"UUU": "F", "CUU": "L", "AUU": "I", "GUU": "V",
		"UUC": "F", "CUC": "L", "AUC": "I", "GUC": "V",
		"UUA": "L", "CUA": "L", "AUA": "I", "GUA": "V",
		"UUG": "L", "CUG": "L", "AUG": "M", "GUG": "V",
		"UCU": "S", "CCU": "P", "ACU": "T", "GCU": "A",
		"UCC": "S", "CCC": "P", "ACC": "T", "GCC": "A",
		"UCA": "S", "CCA": "P", "ACA": "T", "GCA": "A",
		"UCG": "S", "CCG": "P", "ACG": "T", "GCG": "A",
		"UAU": "Y", "CAU": "H", "AAU": "N", "GAU": "D",
		"UAC": "Y", "CAC": "H", "AAC": "N", "GAC": "D",
		"UAA": "Stop", "CAA": "Q", "AAA": "K", "GAA": "E",
		"UAG": "Stop", "CAG": "Q", "AAG": "K", "GAG": "E",
		"UGU": "C", "CGU": "R", "AGU": "S", "GGU": "G",
		"UGC": "C", "CGC": "R", "AGC": "S", "GGC": "G",
		"UGA": "Stop", "CGA": "R", "AGA": "R", "GGA": "G",
		"UGG": "W", "CGG": "R", "AGG": "R", "GGG": "G",
	}
	// change into counts per residue
	countMap := make(map[string]int)
	for _, aa := range codonMap {
		countMap[aa] += 1
	}
	return countMap
}

func rnaCombN(seq string, countMap map[string]int) int {
	// total number of RNA strings that could have formed peptide sequence
	// including stop codon. modulo 1,000,000
	count := countMap["Stop"]
	for i := 0; i < len(seq); i++ {
		aa := seq[i]
		count = (count * countMap[string(aa)]) % 1000000
		fmt.Println(count)
	}
	return count

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	countMap := makeCountMap()
	// calculate for each line of stdin
	for scanner.Scan() {
		seq := scanner.Text()
		n := rnaCombN(seq, countMap)
		fmt.Println(n)
	}
}
