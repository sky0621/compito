package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/sky0621/compito/util"
)

var gSet *util.Set

var argIdx = 1

var showPassage bool

func init() {
	flag.BoolVar(&showPassage, "sp", false, "show passage flg")
	flag.Parse()
	if showPassage {
		argIdx = 2
	}
}

// FIXME this is prototype program
// FIXME think package
func main() {
	start := time.Now()

	// FIXME logging

	args := os.Args

	// FIXME check args

	word := args[argIdx]
	fmt.Printf("入力文字列：%s\n", word)

	lowerWord := strings.ToLower(word)

	inputCharacters := strings.Split(lowerWord, "")
	sort.Strings(inputCharacters)
	fmt.Printf("入力文字列をスライス化してソート：%s\n\n", inputCharacters)

	gSet = util.NewSet()

	fmt.Println("<処理開始>")
	makeWord(inputCharacters, "")

	l := gSet.List()
	sort.Strings(l)

	fmt.Printf("\n<処理結果>\n")
	for _, s := range l {
		fmt.Println(s)
	}

	end := time.Now()

	fmt.Printf("\n処理時間：%f秒\n", (end.Sub(start)).Seconds())
}

func makeWord(addCharacters []string, resultString string) {
	if len(addCharacters) == 0 {
		gSet.Add(resultString)
		return
	}
	for loopIdx, addCharacter := range addCharacters {
		targetCharacter, remains := util.Separate(addCharacters, loopIdx)
		makeWord(remains, resultString+targetCharacter)

		if showPassage {
			fmt.Printf("[追加前：%12s][追加文字：%s][追加後：%12s]\n", resultString, addCharacter, resultString+targetCharacter)
		}
	}
}
