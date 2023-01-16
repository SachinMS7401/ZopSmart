package main

import "fmt"

type bar2 struct {
	x int
	y *int
}

func (b bar2) change() {
	b.x = 33
	*b.y = 44
}

func main() {
	i := 99
	b := bar2{1, &i}
	b.change()
	fmt.Println(b.x, *b.y)
}
