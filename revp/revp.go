package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Fasta struct {
	name string
	seq  string
}

type RestrictionSite struct {
	pos int
	len int
}

var complement = map[string]string{
	"G": "C",
	"C": "G",
	"A": "T",
	"T": "A",
}

func parseFasta(path string) Fasta {
	// FASTA with single entry
	var name string
	var seq = strings.Builder{}
	contents := openFile(path)
	for _, line := range contents {
		if strings.HasPrefix(line, ">") {
			name = line[1:]
		} else {
			seq.WriteString(line)
		}
	}
	return Fasta{name, seq.String()}
}

func openFile(path string) []string {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	content := strings.Split(string(bytes), "\n")
	return content[:len(content)-1] // strip final newline
}

func revc(seq string) string {
	// reverse complement sequence
	builder := strings.Builder{}
	for i := len(seq) - 1; i >= 0; i-- {
		nuc := string(seq[i])
		nucc := complement[nuc]
		builder.WriteString(nucc)
	}
	return builder.String()
}

func isRestrictionSite(site string) bool {
	return site == revc(site)
}

func findRestrictionSites(seq string) []RestrictionSite {
	sites := make([]RestrictionSite, 0)
	for i := 4; i <= 12; i++ {
		for j := 0; j <= len(seq)-i; j++ {
			site := seq[j : j+i]
			if isRestrictionSite(site) {
				sites = append(sites, RestrictionSite{j + 1, i})
			}
		}
	}
	return sites
}

func printSites(sites []RestrictionSite) {
	for _, site := range sites {
		fmt.Printf("%d %d\n", site.pos, site.len)
	}
}

func main() {
	path := os.Args[1]
	fasta := parseFasta(path)
	sites := findRestrictionSites(fasta.seq)
	printSites(sites)
}
