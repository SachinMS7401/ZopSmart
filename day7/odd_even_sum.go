package main

import "fmt"

func sum(x []int) (int, int) {
	even := 0
	odd := 0
	for _, e := range x {
		if e%2 == 0 {
			even += e
		} else {
			odd += e
		}
	}
	return even, odd
}

func main() {
	x := []int{2, 4, 1, 5, 7, 8}
	even, odd := sum(x)
	fmt.Println(even, odd)
}
