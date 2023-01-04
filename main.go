package main

import (
	"bufio"
	"fmt"
	"os"
)

// readInput возвращает срез слов после считывания строки
// и разделения строки на слова
func readInput() ([]string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	var words []string
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	return words, nil
}

func main() {
	var words []string
	var err error

	words, err = readInput()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(len(words))
	}
}
