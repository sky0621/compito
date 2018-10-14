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
	fmt.Printf("lowerWord:%s\n", lowerWord)

	aryWords := strings.Split(lowerWord, "")
	fmt.Println(aryWords)

	sort.Strings(aryWords)
	fmt.Println(aryWords)

	for i, w := range aryWords {
		fmt.Printf("[%3d] %s\n", i, w)

	}
}

// func combi(ary []string, len int) {
// 	if len == 0 {
// 		return ary
// 	}

// }
