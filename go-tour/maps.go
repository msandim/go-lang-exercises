package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	counter := make(map[string]int)
	
	for _, word := range strings.Fields(s) {
		counter[word] = counter[word] + 1
	}
	
	return counter
}

func main() {
	wc.Test(WordCount)
}
