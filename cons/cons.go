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

func getSize(fastas []Fasta) (int, int) {
	// want to know how many seqeuences, and their length
	// assuming their lengths are all equal
	nSeqs := len(fastas)
	var seqLens []int
	for _, fasta := range fastas {
		seqLens = append(seqLens, len(fasta.Seq))
	}
	return nSeqs, seqLens[0]
}

func makeSeqMat(fastas []Fasta) SeqMat {
	nSeqs, _ := getSize(fastas)
	charMat := make([][]string, nSeqs)
	for idx, f := range fastas {
		for i := 0; i < len(f.Seq); i++ {
			nuc := string(f.Seq[i])
			charMat[idx] = append(charMat[idx], nuc)
		}
	}
	return charMat
}

func consensus(x SeqMat) ConMat {
	output := ConMat{}
	for i := 0; i < len(x[0]); i++ {
		counter := initBaseCount()
		for j := 0; j < len(x); j++ {
			counter[x[j][i]]++
		}
		output = append(output, counter)
	}
	return output
}

func getMostCommonNuc(bc BaseCount) string {
	curr_max := 0
	var curr_max_nuc string
	for k, v := range bc {
		if v >= curr_max {
			curr_max = v
			curr_max_nuc = string(k)
		}
	}
	return curr_max_nuc
}

func consensusSeq(x ConMat) string {
	// form consensus sequence from ConMat
	seq := strings.Builder{}
	for _, baseCount := range x {
		mostCommon := getMostCommonNuc(baseCount)
		seq.WriteString(mostCommon)
	}
	return seq.String()
}

func initBaseCount() BaseCount {
	counter := make(BaseCount)
	for _, base := range "ACGT" {
		counter[string(base)] = 0
	}
	return counter
}

func (x ConMat) print() {
	for _, nuc := range "ACGT" {
		fmt.Printf("%s: ", string(nuc))
		for _, seq := range x {
			fmt.Printf("%d ", seq[string(nuc)])
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
