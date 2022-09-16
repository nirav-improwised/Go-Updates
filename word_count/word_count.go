package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	word_count := make(map[string]int)

	fmt.Println("Enter a strig to get word count")
	reader := bufio.NewReader(os.Stdin)

	st, _ := reader.ReadString('\n')
	st = strings.Trim(st, "\n")

	arr_st := strings.Split(st, " ")

	for index, word := range arr_st {
		word := strings.Trim(word, " ")
		// word = strings.Trim(word, "\n")
		arr_st[index] = word
	}

	for _, word := range arr_st {
		if word_count[word] >= 1 {
			word_count[word]++
		} else {
			word_count[word] = 1
		}
	}
	for worD, count := range word_count {
		fmt.Printf("%v\t%v\n", worD, count)
	}
}
