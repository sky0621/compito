package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

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
	fmt.Printf("aryWords:%s\n", aryWords)

	fmt.Println()

	set := newSet()
	combi("", aryWords, set, len(aryWords), []int{})

	next := time.Now()

	l := set.list()
	sort.Strings(l)
	// fmt.Println(l)
	for _, c := range l {
		fmt.Println(c)
	}

	end := time.Now()
	fmt.Println("===========================")
	fmt.Printf("最初から単語の組み合わせ作成まで %f秒\n", (end.Sub(next)).Seconds())
	fmt.Printf("最初から最後の標準出力まで       %f秒\n", (end.Sub(start)).Seconds())
	fmt.Println("===========================")
}

func combi(parentCharacter string, characters []string, resultSet *set, len int, skipIndex []int) {
	if len == 0 {
		return
	}

	thistimeSkipIndex := skipIndex
	for i, c := range characters {
		addTargetWord := fmt.Sprintf("%s%s", parentCharacter, c)
		thistimeSkipIndex = append(skipIndex, i)
		skipFlg := false
		for _, skip := range skipIndex {
			if i == skip {
				skipFlg = true
			}
		}
		if skipFlg {
			continue
		}
		resultSet.add(addTargetWord)

		combi(addTargetWord, characters, resultSet, len-1, thistimeSkipIndex)
	}
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
	m map[string]struct{}
}

func (s *set) list() []string {
	r := []string{}
	// ignore worning
	for k, _ := range s.m {
		// fmt.Println(k, v)
		r = append(r, k)
	}
	return r
}

func (s *set) add(v string) {
	s.m[v] = struct{}{}
}
