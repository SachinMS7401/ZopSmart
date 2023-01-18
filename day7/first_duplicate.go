package main

import "fmt"

func duplicate(s []int) int {
	m := make(map[int]int)
	for _, e := range s {
		if m[e] == 0 {
			m[e] += 1
		} else {
			return e
		}
		/*for j, e2 := range s {
			if i != j {
				if e == e2 {
					return e
				}
			}
		}
		*/
	}
	return -1
}

func main() {
	s := []int{1, 2, 3, 2, 1}
	f := duplicate(s)
	fmt.Println(f)
}
