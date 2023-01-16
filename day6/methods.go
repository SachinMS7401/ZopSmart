package main

import "fmt"

type bar struct {
	x int
	y *int
}

func (b bar) change() {
	b.x = 33
	z := 5
	b.y = &z
}
func main() {
	i := 99
	b := bar{1, &i}
	b.change()
	fmt.Println(b.x, *b.y)
}
