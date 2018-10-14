package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

// FIXME this is prototype program
// FIXME think package
func main() {
	// FIXME logging

	args := os.Args

	// FIXME check args

	word := args[1]
	fmt.Printf("word:%s\n", word)

	lowerWord := strings.ToLower(word)

	aryWords := strings.Split(lowerWord, "")

	fmt.Println()

	set := newSet()
	parentString := ""
	// includes := []int{}
	for i, w := range aryWords {
		set.add(fmt.Sprintf("%s%s", parentString, w))

		parentString2 := fmt.Sprintf("%s%s", parentString, w)
		for i2, w2 := range aryWords {
			if i2 == i {
				continue
			}
			set.add(fmt.Sprintf("%s%s", parentString2, w2))

			parentString3 := fmt.Sprintf("%s%s", parentString2, w2)
			for i3, w3 := range aryWords {
				if i3 == i2 {
					continue
				}
				if i3 == i {
					continue
				}
				set.add(fmt.Sprintf("%s%s", parentString3, w3))

				parentString4 := fmt.Sprintf("%s%s", parentString3, w3)
				for i4, w4 := range aryWords {
					if i4 == i3 {
						continue
					}
					if i4 == i2 {
						continue
					}
					if i4 == i {
						continue
					}
					set.add(fmt.Sprintf("%s%s", parentString4, w4))

					parentString5 := fmt.Sprintf("%s%s", parentString4, w4)
					for i5, w5 := range aryWords {
						if i5 == i4 {
							continue
						}
						if i5 == i3 {
							continue
						}
						if i5 == i2 {
							continue
						}
						if i5 == i {
							continue
						}
						set.add(fmt.Sprintf("%s%s", parentString5, w5))
					}
				}
			}
		}
	}

	l := set.list()
	sort.Strings(l)
	// fmt.Println(l)
	for _, c := range l {
		fmt.Println(c)
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
