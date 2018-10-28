package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"sync"
	"time"
)

var exclusionList = getExclusionList()

// FIXME this is prototype program
// FIXME think package
func main() {
	start := time.Now()

	// FIXME logging

	args := os.Args

	// FIXME check args

	word := args[1]
	fmt.Printf("word:%s\n", word)

	lowerWord := strings.ToLower(word)

	aryWords := strings.Split(lowerWord, "")
	sort.Strings(aryWords)
	fmt.Printf("aryWords:%s\n", aryWords)

	fmt.Println()

	set := newSet()
	addWord(set, aryWords, "")

	next := time.Now()

	l := set.list()
	sort.Strings(l)
	// fmt.Println(l)
	for _, c := range l {
		fmt.Println(c)
	}

	end := time.Now()
	fmt.Println("===========================")
	fmt.Printf("生成単語数 %d\n", len(l))
	fmt.Printf("最初から単語の組み合わせ作成まで %f秒\n", (end.Sub(next)).Seconds())
	fmt.Printf("最初から最後の標準出力まで       %f秒\n", (end.Sub(start)).Seconds())
	fmt.Println("===========================")
}

func addWord(set *set, characters []string, parentChar string) {
	chLen := len(characters)
	for i, c := range characters {
		thisWord := fmt.Sprintf("%s%s", parentChar, c)
		if containsInSlice(thisWord, exclusionList) {
			continue
		}
		set.add(thisWord)
		cpCharacters := make([]string, chLen, chLen)
		copy(cpCharacters, characters)
		addWord(set, getRemaining(cpCharacters, i), fmt.Sprintf("%s%s", parentChar, c))
	}
}

func getRemaining(characters []string, index int) []string {
	if characters == nil {
		return nil
	}
	if len(characters) <= index {
		return nil
	}

	before := characters[:index]
	after := characters[index+1:]
	return append(before, after...)
}

func getExclusionList() []string {
	return []string{
		"ae",
		"afk", "afn", "afs", "atf", "atk", "ats",
		"ean", "eak", "ekn",
		"enna", "ennf", "ennk", "enns", "ennt",
		"fn", "fs", "fk",
		"frk", "ftr", "fta", "fte", "ftn", "fts", "fnn",
		"ias", "ika", "ikn", "ikr", "iks", "ikt", "ikf", "ifa", "ife", "ifk", "ifr", "ifs", "ift", "iea", "iet", "iek", "ief",
		"kt", "ks", "kf",
		"knn", "krn", "krf", "krt", "knf", "knt", "kns",
		"kren",
		"nn", "nk", "nf",
		"ntk", "nts", "nrk", "nre", "nte", "nta", "nti", "ntn", "ntr",
		"rk", "rt", "rs", "rn",
		"rfk", "rfs", "rft", "rfa", "rfe", "rfn",
		"sf",
		"snn", "snk", "snf", "snt", "stn", "stf", "stk", "srn", "srt", "srk", "snr", "skn", "sek",
		"tf", "tk", "tn",
		"tnn", "tsr", "tsn", "tse", "tsf", "trs", "trf", "trn", "trk", "tef", "tsk",
		"tska", "tske", "tskn", "tskf", "tskr", "tsan", "tsak", "tsaf", "tsae",
		"tsarn", "tsark", "tsarf",
	}
}

func containsInSlice(str string, slice []string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}

// -------------------------------------------------------------
// Set
// -------------------------------------------------------------
func newSet() *set {
	m := map[string]struct{}{}
	s := &set{}
	s.m = m
	return s
}

type set struct {
	m   map[string]struct{}
	mux sync.Mutex
}

func (s *set) list() []string {
	s.mux.Lock()
	defer s.mux.Unlock()
	r := []string{}
	// ignore worning
	for k, _ := range s.m {
		// fmt.Println(k, v)
		r = append(r, k)
	}
	return r
}

func (s *set) add(v string) {
	s.mux.Lock()
	defer s.mux.Unlock()
	s.m[v] = struct{}{}
}
