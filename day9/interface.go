package main

import "fmt"

type vowelfind interface {
	findvowel() []rune
}

type mysting string

func (s mysting) findvowel() []rune {
	var vowels []rune
	for _, rune := range s {
		if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' {
			vowels = append(vowels, rune)
		}
	}
	return vowels
}

func main() {
	m := mysting("hello world")
	var v vowelfind
	v = m
	fmt.Printf("vowels in the strings are %c", v.findvowel())
	fmt.Println()

}
