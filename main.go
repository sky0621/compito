package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/sky0621/compito/util"
)

var gSet *util.Set

// FIXME this is prototype program
// FIXME think package
func main() {
	start := time.Now()

	// FIXME logging

	args := os.Args

	// FIXME check args

	word := args[1]
	fmt.Printf("input word:%s\n", word)

	lowerWord := strings.ToLower(word)

	aryWords := strings.Split(lowerWord, "")
	sort.Strings(aryWords)
	fmt.Printf("aryWords:%s\n", aryWords)

	gSet = util.NewSet()

	makeWord(aryWords, "")

	l := gSet.List()
	sort.Strings(l)

	fmt.Println()
	for _, s := range l {
		fmt.Println(s)
	}

	end := time.Now()
	fmt.Println("===========================")
	fmt.Printf("最初から最後の標準出力まで       %f秒\n", (end.Sub(start)).Seconds())
	fmt.Println("===========================")
}

func makeWord(addCharacters []string, resultString string) {
	if len(addCharacters) == 0 {
		gSet.Add(resultString)
		return
	}
	for loopIdx, addCharacter := range addCharacters {
		targetCharacter, remains := util.Separate(addCharacters, loopIdx)
		makeWord(remains, resultString+targetCharacter)

		fmt.Printf("[base:%s][addCharacter:%s][resultString:%s]\n", resultString, addCharacter, resultString+targetCharacter)
	}
}
