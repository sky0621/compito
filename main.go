package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/sky0621/compito/util"
)

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

	target, remains := util.Separate(aryWords, 2)
	fmt.Printf("target: %s\n", target)
	for _, r := range remains {
		fmt.Printf("remain: %s\n", r)
	}

	end := time.Now()
	fmt.Println("===========================")
	fmt.Printf("最初から最後の標準出力まで       %f秒\n", (end.Sub(start)).Seconds())
	fmt.Println("===========================")
}
func main2() {
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

	for idx, char := range aryWords {

		for idx2, char2 := range aryWords {
			if idx2 == idx {
				continue
			}

			for idx3, char3 := range aryWords {
				if idx3 == idx2 || idx3 == idx {
					continue
				}

				for idx4, char4 := range aryWords {
					if idx4 == idx3 || idx4 == idx2 || idx4 == idx {
						continue
					}

					for idx5, char5 := range aryWords {
						if idx5 == idx4 || idx5 == idx3 || idx5 == idx2 || idx5 == idx {
							continue
						}

						fmt.Printf("%s%s%s%s%s\n", char, char2, char3, char4, char5)
					}
				}
			}
		}
	}
	fmt.Println()

	end := time.Now()
	fmt.Println("===========================")
	fmt.Printf("最初から最後の標準出力まで       %f秒\n", (end.Sub(start)).Seconds())
	fmt.Println("===========================")
}

// const
// [c][o][n][s][t]
// [c] [onst]
//     [o] [nst]
//         [n] [st]
//             [s] [t]

// [c] [o] [n] [s] [t]
//     [n] [s] [t]
//     [s] [t]
//     [t]
