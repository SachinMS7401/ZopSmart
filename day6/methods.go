package main

import "fmt"

// Method does not change the outer receiver .pass by value receiver
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
