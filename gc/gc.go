package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Fasta struct {
	Name string
	Seq  string
}

func strip(s string) string {
	if strings.HasPrefix(s, ">") {
		s = s[1:]
	}
	return s
}

func openFile(path string) []string {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	content := strings.Split(string(bytes), "\n")
	return content
}

func parseFasta(path string) []Fasta {
	// awful fasta parser, but it works
	var result []Fasta
	var name string
	content := openFile(path)
	fastaStore := make(map[string][]string)
	for _, line := range content {
		if strings.HasPrefix(line, ">") {
			// is neq sequence
			name = strip(line)
		} else {
			fastaStore[name] = append(fastaStore[name], line)
		}
	}
	for name, seqArr := range fastaStore {
		var fullSeq = strings.Builder{}
		for _, i := range seqArr {
			fullSeq.WriteString(i)
		}
		result = append(result, Fasta{name, fullSeq.String()})
	}
	return result
}

func pGC(seq string) float64 {
	// proportion of GC content in seq
	var lenStr = len(seq)
	var count = 0
	for i := 0; i < lenStr; i++ {
		char := string(seq[i])
		if char == "G" || char == "C" {
			count++
		}
	}
	return float64(count) / float64(lenStr)
}

func main() {
	path := os.Args[1]
	seqs := parseFasta(path)
	var currMaxProp = 0.0
	var currMaxName = ""
	for _, seq := range seqs {
		propGC := pGC(seq.Seq)
		if propGC > currMaxProp {
			currMaxProp = propGC
			currMaxName = seq.Name
		}
	}
	fmt.Println(currMaxName)
	fmt.Println(currMaxProp * 100)
}
