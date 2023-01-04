package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

// readInput reads source string
// from command line arguments and returns them.
func readInput() (src string) {
	flag.Parse()
	src = strings.Join(flag.Args(), "")
	return src
}

func wordsCount(words []string) int {
	var i int
	for _, v := range words {
		if v == "" {
			continue
		}
		i++
	}
	return i
}

func main() {
	var cnt int
	var words []string
	src := readInput()
	words = strings.Split(src, " ")

	cnt = wordsCount(words)

	fmt.Println(strconv.Itoa(cnt))
}
