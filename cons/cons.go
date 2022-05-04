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

type SeqMat [][]string

type ConMat []BaseCount

type Consensus [][]string

type BaseCount map[string]int

func strip(s string) string {
	if strings.HasPrefix(s, ">") {
		s = s[1:]
	}
	return s
}

func openFile(path string) []string {
	bytes, _ := ioutil.ReadFile(path)
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

func makeSeqMat(fastas []Fasta) SeqMat {
	charMat := make([][]string, len(fastas))
	for idx, f := range fastas {
		for i := 0; i < len(f.Seq); i++ {
			base := string(f.Seq[i])
			charMat[idx] = append(charMat[idx], base)
		}
	}
	return charMat
}

func consensus(x SeqMat) ConMat {
	output := ConMat{}
	for i := 0; i < len(x[0]); i++ {
		counter := make(BaseCount)
		for j := 0; j < len(x); j++ {
			counter[x[j][i]]++
		}
		output = append(output, counter)
	}
	return output
}

func mostCommonBase(bc BaseCount) string {
	maxCount := 0
	var mostCommonBase string
	for base, count := range bc {
		if count >= maxCount {
			maxCount = count
			mostCommonBase = string(base)
		}
	}
	return mostCommonBase
}

func consensusSeq(x ConMat) string {
	// form consensus sequence from ConMat
	seq := strings.Builder{}
	for _, baseCount := range x {
		mostCommon := mostCommonBase(baseCount)
		seq.WriteString(mostCommon)
	}
	return seq.String()
}

func (x ConMat) print() {
	// pretty print consensus matrix for output
	for _, base := range "ACGT" {
		fmt.Printf("%s: ", string(base))
		for _, seq := range x {
			fmt.Printf("%d ", seq[string(base)])
		}
		fmt.Println()
	}
}

func main() {
	path := os.Args[1]
	fastas := parseFasta(path)
	m := makeSeqMat(fastas)
	con := consensus(m)
	seq := consensusSeq(con)
	fmt.Println(seq)
	con.print()
}
