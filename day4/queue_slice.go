package main

import "fmt"

func Dequeue(s []int) []int {
	if s == nil {
		return nil
	}
	f := s[0]
	s = s[1:]
	fmt.Println(f)
	return s
}
func Enqueue(s []int, d int) []int {

	s = append(s, d)
	return s
}

func main() {
	var s []int

	s = Enqueue(s, 3)

	s = Enqueue(s, 7)
	fmt.Println(s)
	fmt.Println(len(s))

	s = Dequeue(s)
	fmt.Println(s)

}
