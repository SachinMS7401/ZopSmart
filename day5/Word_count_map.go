package main

import (
	"fmt"
	"strings"
)

func wordcount(s string) map[string]int {
	words := strings.Fields(s)
	m := make(map[string]int)
	for _, word := range words {
		m[word] += 1
	}
	return m
}

func main() {
	s := "Hi i am Sachin"
	fmt.Println(wordcount(s))
}
