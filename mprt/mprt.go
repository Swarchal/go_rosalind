package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func parseInput(path string) []string {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalln(err)
	}
	return strings.Split(string(content), "\n")
}

func getSeq(uniprotID string) string {
	// return peptide sequence for a given uniprot ID
	url := fmt.Sprintf("http://www.uniprot.org/uniprot/%s.fasta", uniprotID)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return stripHeader(string(body))
}

func stripHeader(fasta string) string {
	// remove header, leave just sequence on single line
	splitFasta := strings.Split(fasta, "\n")
	return strings.Join(splitFasta[1:], "")
}

func findMotif(seq string) []int {
	// find location(s) of N-glycosylation motif in peptide
	//
	// go doesn't have lookaheads in their regex library for whatever stupid
	// reason, so this is a mess, but it's their fault ...
	var locations []int
	re := regexp.MustCompile("N[^P][[ST][^P]")
	for i := 0; i < len(seq)-4; i++ {
		substr := seq[i : i+4]
		if re.MatchString(substr) {
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
	for _, uniprotID := range parseInput(os.Args[1]) {
		seq := getSeq(uniprotID)
		locs := findMotif(seq)
		if len(locs) > 0 {
			fmt.Println(uniprotID)
			printSlice(locs)
		}
	}
}
