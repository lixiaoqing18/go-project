package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	ss := strings.Split(s, " ")
	result := make(map[string]int)
	for _, word := range ss {
		count, ok := result[word]
		if !ok {
			result[word] = 1
		} else {
			result[word] = count + 1
		}
	}
	return result
}

func main() {
	wc.Test(WordCount)
}
