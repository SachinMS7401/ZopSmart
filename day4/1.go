package main

import "fmt"

func main() {
	a := []int{}
	capp := 0
	for i := 0; i < 10000; i += 1 {
		if cap(a) != capp {
			capp = cap(a)
			fmt.Println(len(a), cap(a))
		}

		a = append(a, i)
	}
}
